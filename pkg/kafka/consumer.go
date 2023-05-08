package kafka

import (
	"context"
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

// Worker kafka consumer worker fetch and process messages from reader
type Worker func(ctx context.Context, r *kafka.Reader, w *kafka.Writer, wg *sync.WaitGroup, workerID int)

type consumerGroup struct {
	Brokers []string
	GroupID string
}

// NewConsumerGroup kafka consumer group constructor
func NewConsumerGroup(brokers []string, groupID string) *consumerGroup {
	return &consumerGroup{Brokers: brokers, GroupID: groupID}
}

func (c *consumerGroup) GetNewKafkaReader(kafkaURL []string, groupTopics []string, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:                kafkaURL,
		GroupID:                groupID,
		GroupTopics:            groupTopics,
		StartOffset:            kafka.FirstOffset,
		MinBytes:               MIN_BYTES,
		MaxBytes:               MAX_BYTES,
		QueueCapacity:          QUEUE_CAPACITY,
		HeartbeatInterval:      HEARTBEAT_INTERVAL,
		CommitInterval:         COMMIT_INTERVAL,
		PartitionWatchInterval: PARTITION_WATCH_INTERVAL,
		MaxAttempts:            MAX_ATTEMPTS,
		MaxWait:                MAX_WAIT,
		Dialer: &kafka.Dialer{
			Timeout: DIAL_TIMEOUT,
		},
	})
}

// ConsumeTopic start consumer group with given worker and pool size
func (c *consumerGroup) ConsumeTopic(ctx context.Context, groupTopics []string, w *kafka.Writer, poolSize int, worker Worker) {
	r := c.GetNewKafkaReader(c.Brokers, groupTopics, c.GroupID)

	defer func() {
		if err := r.Close(); err != nil {
			log.Warn().Msgf("consumerGroup.r.Close: %v", err)
		}
	}()

	log.Info().Msgf("Starting consumer groupID: %s, topic: %+v, pool size: %v", c.GroupID, groupTopics, poolSize)

	wg := &sync.WaitGroup{}
	for i := 0; i <= poolSize; i++ {
		wg.Add(1)
		go worker(ctx, r, w, wg, i)
	}
	wg.Wait()
}
