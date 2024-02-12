package kafka

import (
	"context"
	"strings"

	"github.com/IBM/sarama"
	"github.com/thangpham4/self-project/pkg/envx"
	"github.com/thangpham4/self-project/pkg/kafkax"
	"github.com/thangpham4/self-project/pkg/logger"
)

var (
	kafkaHost = envx.String("KAFKA_HOST", "localhost:9092")
	topics    = []string{"test-topic"}
	group     = "0"
	version   = "2.1.1"
	assignor  = "range"
	oldest    = true
)

type OrderKafka struct {
	*kafkax.KafkaConfig
}

func NewOrderKafka() (*OrderKafka, error) {
	version, err := sarama.ParseKafkaVersion(version)
	if err != nil {
		logger.Error(err, "Error parsing Kafka version")
		return nil, err
	}

	// Create a new consumer
	config := sarama.NewConfig()
	config.Version = version

	switch assignor {
	case "sticky":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategySticky()}
	case "roundrobin":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	case "range":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
	default:
		logger.Error(err, "Unrecognized consumer group partition assignor", "assignor", assignor)
		return nil, err
	}

	if oldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	client, err := sarama.NewConsumerGroup(strings.Split(kafkaHost, ","), group, config)
	if err != nil {
		logger.Error(err, "Error creating consumer group client")
		return nil, err
	}
	return &OrderKafka{
		KafkaConfig: &kafkax.KafkaConfig{
			Client:             client,
			ImplementFunctions: []kafkax.ImplementFuncs{LogKafkaMessage},
			Topics:             topics,
			Group:              group,
		},
	}, nil
}

func LogKafkaMessage(ctx context.Context, msg []byte) {
	loggerInstance := logger.Factory("Kafka Message Order")
	loggerInstance.V(logger.LogInfoLevel).Info(string(msg))
}
