package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/go-playground/validator/v10"
	"github.com/ngtrdai197/cobra-cmd/config"
	"github.com/ngtrdai197/cobra-cmd/pkg/kafka"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var kafkaConsumerCmd = &cobra.Command{
	Use:   "kafka-consumer-cmd",
	Short: "Serve worker application",
	Long: `A longer description that spans multiple lines and likely contains
			examples and usage of using your application. For example:

			Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
		defer cancel()

		c, err := config.GetConfig(validator.New())
		if err != nil {
			panic(fmt.Errorf("Config file invalidate with error: %w", err))
		}
		if err := initConsumer(ctx, c); err != nil {
			log.Fatal().Err(err)
		}
	},
}

func initConsumer(ctx context.Context, c *config.Config) error {
	// Kafka Producer
	kp := kafka.NewProducer(strings.Split(c.KafkaBrokers, ","), c)
	defer kp.Close()

	cg := kafka.NewConsumerGroup(strings.Split(c.KafkaBrokers, ","), "create_user_send_email")

	// TODO: Replace nil args to kafka writer and service worker process
	go cg.ConsumeTopic(ctx, getConsumerGroupTopics(c), nil, kafka.PoolSize, nil)
	<-ctx.Done()
	return nil
}

// getConsumerGroupTopics Get kafka consumer group topics
func getConsumerGroupTopics(c *config.Config) []string {
	return []string{
		c.KafkaCreateUserSendEmail,
	}
}

func init() {
	rootCmd.AddCommand(kafkaConsumerCmd)
}
