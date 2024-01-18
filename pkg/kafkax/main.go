package kafkax

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
	"github.com/thangpham4/self-project/pkg/errgroup"
	"github.com/thangpham4/self-project/pkg/logger"
)

type KafkaConfig struct {
	Client             sarama.ConsumerGroup
	ImplementFunctions []ImplementFuncs
	Topics             []string
	Group              string
	Logger             logger.Logger
}

func (k *KafkaConfig) Start(ctx context.Context) error {
	k.Logger.Info("Starting to consume", "group", k.Group, "topics", k.Topics)
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error { return k.consume(ctx) })
	return eg.Wait()
}

func (k *KafkaConfig) Stop() {
	k.Client.Close()
}

func (k *KafkaConfig) consume(ctx context.Context) error {
	consumer := Consumer{
		ready: make(chan bool),
	}
	// consumptionIsPaused := false
	wg := &sync.WaitGroup{}
	wg.Add(1)

	errs := make(chan error)
	go func(errs chan<- error) {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			err := k.Client.Consume(ctx, k.Topics, &consumer)
			errs <- err
			// check if context was canceled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}(errs)

	for err := range errs {
		if err != nil {
			return err
		}
	}

	<-consumer.ready // Await till the consumer has been set up
	log.Println("Sarama consumer up and running!...")

	// sigusr1 := make(chan os.Signal, 1)
	// signal.Notify(sigusr1, syscall.SIGUSR1)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	keepRunning := true

	for keepRunning {
		select {
		case <-ctx.Done():
			log.Println("terminating: context canceled")
			keepRunning = false
		case <-sigterm:
			log.Println("terminating: via signal")
			keepRunning = false
		// case <-sigusr1:
		// 	toggleConsumptionFlow(k.Client, &consumptionIsPaused)
		}
	}
	wg.Wait()
	return nil
}

//nolint:unused
func toggleConsumptionFlow(client sarama.ConsumerGroup, isPaused *bool) {
	if *isPaused {
		client.ResumeAll()
		log.Println("Resuming consumption")
	} else {
		client.PauseAll()
		log.Println("Pausing consumption")
	}

	*isPaused = !*isPaused
}

type ImplementFuncs func(ctx context.Context, msg []byte)

type Consumer struct {
	ready chan bool
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/IBM/sarama/blob/main/consumer_group.go#L27-L29
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Printf("message channel was closed")
				return nil
			}
			log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
			session.MarkMessage(message, "")
		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/IBM/sarama/issues/1192
		case <-session.Context().Done():
			return nil
		}
	}
}
