package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/services"
)

type MockHandler struct {
	mockService *services.MockService
	logger      logger.Logger
}

type Message struct {
	Before string
	After  string
}

func NewMockHandler(
	mockService *services.MockService,
) *MockHandler {
	return &MockHandler{
		mockService: mockService,
		logger:      logger.Factory("MockHandler"),
	}
}

func (m *MockHandler) GetCache(ctx *gin.Context) {
	m.logger.V(logger.LogDebugLevel).Info("Running MockHandlerfor cache")
	mockEn, err := m.mockService.GetMockCache(ctx)
	mockEn.Path = ctx.Request.URL
	if err != nil {
		m.logger.Error(err, "unknown")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "err unknown"})
	}
	ctx.IndentedJSON(http.StatusOK, mockEn)
	if err != nil {
		m.logger.Error(err, "error parse json")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("error in parsing json, err: %w", err)})
	}
}

func (m *MockHandler) SendMessage(ctx *gin.Context) {
	var message Message

	err := json.NewDecoder(ctx.Request.Body).Decode(&message)
	if err != nil {
		m.logger.Error(err, "error in parse json", "struct", ctx.Request.Body)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msgStr, err := json.Marshal(message)
	if err != nil {
		m.logger.Error(err, "error in parse json", "struct", ctx.Request.Body)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	// config.Producer.Partitioner = sarama.NewRandomPartitioner("mock-topic")

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		m.logger.Error(err, "error in creating producer")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a message
	msg := &sarama.ProducerMessage{
		Topic: "test-topic",
		Value: sarama.StringEncoder(string(msgStr)),
	}

	// Send the message
	part, offset, err := producer.SendMessage(msg)
	if err != nil {
		m.logger.Error(err, "error in sending message")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	producer.Close()

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"messgae":   string(msgStr),
		"status":    "OK",
		"partition": part,
		"offset":    offset,
	})
}

func (m *MockHandler) ReceiveMessage(ctx *gin.Context) {
	config := sarama.NewConfig()

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		m.logger.Error(err, "error in creating consumer")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Subscribe to the topic
	topic := "test-topic"
	part, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		m.logger.Error(err, "error in consuming partition")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := <-part.Messages()
	msgStr := message.Value

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"messgae": string(msgStr),
	})
}
