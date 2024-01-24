package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"github.com/thangpham4/self-project/pkg/blobx"
	"github.com/thangpham4/self-project/pkg/envx"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/services"
)

var (
	kafkaHost = envx.String("KAFKA_HOST", "localhost:9092")
	topicName = "test-topic"
)

type MockHandler struct {
	mockService    *services.MockService
	blobConnection *azblob.Client
	logger         logger.Logger
}

type Message struct {
	Before string
	After  string
}

func NewMockHandler(
	mockService *services.MockService,
	blobConnection *azblob.Client,
) *MockHandler {
	return &MockHandler{
		mockService:    mockService,
		blobConnection: blobConnection,
		logger:         logger.Factory("MockHandler"),
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

	producer, err := sarama.NewSyncProducer([]string{kafkaHost}, config)
	if err != nil {
		m.logger.Error(err, "error in creating producer")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a message
	msg := &sarama.ProducerMessage{
		Topic: topicName,
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
		"topic":     topicName,
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

func (m *MockHandler) ListBlob(ctx *gin.Context) {
	blobInstance := blobx.NewBlobService(m.blobConnection)
	out, err := blobInstance.ListBlob()
	if err != nil {
		m.logger.Error(err, "err in list blob")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"list_blob": out,
	})
}

func (m *MockHandler) ReadBlob(ctx *gin.Context) {
	path := ctx.Query("path")
	blobInstance := blobx.NewBlobService(m.blobConnection)
	out, err := blobInstance.GetCSV(path)
	if err != nil {
		m.logger.Error(err, "error in reading csv", "path", path)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"data": out,
	})
}
