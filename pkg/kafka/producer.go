package kafka

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ngtrdai197/cobra-cmd/config"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"github.com/sethvargo/go-retry"
)

const (
	TimeStepRetrySendKafka = 500
	MaxRetrySendKafka      = 5
)

type Producer struct {
	writer *kafka.Writer
	config *config.Config
}

func NewProducer(brokers []string, config *config.Config) *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:         kafka.TCP(brokers...),
			Balancer:     &kafka.RoundRobin{},
			Async:        false,
			ReadTimeout:  WRITER_READ_TIMEOUT,
			WriteTimeout: WRITER_WRITE_TIMEOUT,
			MaxAttempts:  WRITER_MAX_ATTEMPTS,
			RequiredAcks: WRITER_REQUIRED_ACKS,
		},
		config: config,
	}
}

func (p *Producer) SendCreateUserEmailMessage(ctx context.Context, data interface{}) error {
	msgBytes, err := json.Marshal(data)
	if err != nil {
		return errors.Wrap(err, "json.Marshal")
	}
	timeStepRetry := retry.NewConstant(TimeStepRetrySendKafka * time.Millisecond)
	if err := retry.Do(ctx, retry.WithMaxRetries(MaxRetrySendKafka, timeStepRetry), func(ctx context.Context) error {
		if err := p.writer.WriteMessages(ctx, kafka.Message{
			Topic: p.config.KafkaCreateUserSendEmail,
			Value: msgBytes,
			Time:  time.Time{},
		}); err != nil {
			return errors.Wrap(err, "p.writer.WriteMessages")
		}

		return nil
	}); err != nil {
		return err
	}
	log.Info().Msg("Send kafka create user email sucessfully")
	return nil
}

func (p *Producer) Close() error {
	return p.writer.Close()
}

func (p *Producer) GetConfig() *config.Config {
	return p.config
}
