package kafka

import (
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go/compress"

	"github.com/ngtrdai197/cobra-cmd/config"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"github.com/sethvargo/go-retry"
)

const (
	TimeStepRetrySendKafka = 500
	MaxRetrySendKafka      = 5
	PoolSize               = 3
)

type Producer struct {
	W      *kafka.Writer
	config *config.Config
}

func NewProducer(brokers []string, config *config.Config) *Producer {
	return &Producer{
		W: &kafka.Writer{
			Addr:         kafka.TCP(brokers...),
			Balancer:     &kafka.LeastBytes{},
			RequiredAcks: WriterRequiredAcks,
			MaxAttempts:  WriterMaxAttempts,
			ErrorLogger:  kafka.LoggerFunc(log.Error().Msgf),
			Compression:  compress.Snappy,
			ReadTimeout:  WriterReadTimeout,
			WriteTimeout: WriterWriteTimeout,
			Async:        false,
		},
		config: config,
	}
}

func (p *Producer) SendCreateUserEmailMessage(ctx context.Context, data TestData) error {
	msgBytes, err := json.Marshal(data)
	if err != nil {
		return errors.Wrap(err, "json.Marshal")
	}
	timeStepRetry := retry.NewConstant(TimeStepRetrySendKafka * time.Millisecond)
	if err := retry.Do(ctx, retry.WithMaxRetries(MaxRetrySendKafka, timeStepRetry), func(ctx context.Context) error {
		if err := p.W.WriteMessages(ctx, kafka.Message{
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
	log.Info().Msg("send kafka create user email successfully")
	return nil
}

func (p *Producer) Close() error {
	return p.W.Close()
}

func (p *Producer) GetConfig() *config.Config {
	return p.config
}
