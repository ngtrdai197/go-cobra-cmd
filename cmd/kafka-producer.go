package cmd

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/ngtrdai197/cobra-cmd/config"
	"github.com/ngtrdai197/cobra-cmd/pkg/kafka"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var kafkaProducerCmd = &cobra.Command{
	Use:   "kafka-producer-cmd",
	Short: "Serve worker application",
	Long: `A longer description that spans multiple lines and likely contains
			examples and usage of using your application. For example:

			Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, _ []string) {
		c, err := config.GetConfig(validator.New())
		if err != nil {
			panic(fmt.Errorf("config file invalidate with error: %w", err))
		}
		log.Info().Msg("kafkaProducerCmd")
		kp := kafka.NewProducer(strings.Split(c.KafkaBrokers, ","), c)
		if err := kp.SendCreateUserEmailMessage(cmd.Context(), kafka.TestData{
			Name:  "Dai Nguyen",
			Phone: "023127832",
		}); err != nil {
			log.Error().Str("topic", c.KafkaCreateUserSendEmail).AnErr("error", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(kafkaProducerCmd)
}
