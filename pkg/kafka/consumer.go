package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/ngtrdai197/cobra-cmd/config"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

// Worker kafka consumer worker fetch and process messages from reader
type Worker func(ctx context.Context, r *kafka.Reader, w *kafka.Writer, wg *sync.WaitGroup, workerID int)

type ConsumerGroup struct {
	Brokers []string
	GroupID string
}

// NewConsumerGroup kafka consumer group constructor
func NewConsumerGroup(brokers []string, groupID string) *ConsumerGroup {
	return &ConsumerGroup{Brokers: brokers, GroupID: groupID}
}

type ReaderMessageProcessor struct {
	config *config.Config
}

type TestData struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (c *ConsumerGroup) GetNewKafkaReader(kafkaURL []string, groupTopics []string, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:                kafkaURL,
		GroupID:                groupID,
		GroupTopics:            groupTopics,
		StartOffset:            kafka.FirstOffset,
		MinBytes:               MinBytes,
		MaxBytes:               MaxBytes,
		QueueCapacity:          QueueCapacity,
		HeartbeatInterval:      HeartbeatInterval,
		CommitInterval:         CommitInterval,
		PartitionWatchInterval: PartitionWatchInterval,
		MaxAttempts:            MaxAttempts,
		MaxWait:                MaxWait,
		Dialer: &kafka.Dialer{
			Timeout: DialTimeout,
		},
	})
}

// ConsumeTopic start consumer group with given worker and pool size
func (c *ConsumerGroup) ConsumeTopic(ctx context.Context, groupTopics []string, w *kafka.Writer, poolSize int, worker Worker) {
	r := c.GetNewKafkaReader(c.Brokers, groupTopics, c.GroupID)

	defer func() {
		if err := r.Close(); err != nil {
			log.Warn().Msgf("consumerGroup.r.Close: %v", err)
		}
	}()

	log.Info().Msgf("starting consumer groupID: %s, topic: %+v, pool size: %v", c.GroupID, groupTopics, poolSize)

	wg := &sync.WaitGroup{}
	for i := 0; i <= poolSize; i++ {
		wg.Add(1)
		go worker(ctx, r, w, wg, i)
	}
	wg.Wait()
}

func NewReaderMessageProcessor(config *config.Config) *ReaderMessageProcessor {
	return &ReaderMessageProcessor{config: config}
}

func (rmp *ReaderMessageProcessor) ProcessMessages(ctx context.Context, r *kafka.Reader, w *kafka.Writer, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		m, err := r.FetchMessage(ctx)
		if err != nil {
			log.Warn().Msgf("workerID: %v, err: %v", workerID, err)
			continue
		}

		log.Debug().Str("topic", m.Topic).Int("partition", m.Partition).Int("offset", int(m.Offset)).Time("time", m.Time).Int("worker_id", workerID)

		switch m.Topic {
		case rmp.config.KafkaCreateUserSendEmail:
			Test(ctx, rmp, r, m)
		}
	}
}

func Test(ctx context.Context, readerMessage *ReaderMessageProcessor, r *kafka.Reader, m kafka.Message) {
	defer readerMessage.commitMessage(ctx, r, m)
	var data TestData
	err := json.Unmarshal(m.Value, &data)
	if err != nil {
		fmt.Println(err)
	}
	log.Info().Msgf("%s %+v", readerMessage.config.KafkaCreateUserSendEmail, data)
}

func (rmp *ReaderMessageProcessor) commitMessage(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	if err := r.CommitMessages(ctx, m); err != nil {
		log.Warn().Msgf("commitMessage error details=%v", err)
		return
	}
	log.Info().Str("topic", m.Topic).Int("partition", m.Partition).Int("offset", int(m.Offset)).Msg("committed kafka message")
}
