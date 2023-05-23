package cmd

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/hibiken/asynq"
	"github.com/ngtrdai197/cobra-cmd/config"
	"github.com/ngtrdai197/cobra-cmd/pkg/worker"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client-cmd",
	Short: "Serve worker application",
	Long: `A longer description that spans multiple lines and likely contains
			examples and usage of using your application. For example:

			Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, _ []string) {
		c, err := config.GetConfig(validator.New())
		if err != nil {
			panic(fmt.Errorf("Config file invalidate with error: %w", err))
		}

		initClient(cmd.Context(), c)
	},
}

func initClient(ctx context.Context, c *config.Config) {
	d := worker.NewRedisTaskDistributor(asynq.RedisClientOpt{
		Addr: fmt.Sprintf("%s:%s", c.RedisHost, c.RedisPort), DB: c.RedisDb,
	})
	err := d.DistributeTaskEmailDelivery(ctx, &worker.EmailDeliveryPayload{
		Name:  "Dai Nguyen",
		Phone: "0375629888",
	}, asynq.Queue(worker.QueuePriorityCritical))
	if err != nil {
		log.Error().Msgf("Has error %s = %v", worker.DeliveryEmailQueue, err)
	}
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
