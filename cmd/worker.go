package cmd

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/hibiken/asynq"
	"github.com/ngtrdai197/cobra-cmd/config"
	"github.com/ngtrdai197/cobra-cmd/pkg/worker"
	"github.com/ngtrdai197/cobra-cmd/pkg/worker/email_delivery"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var workerCmd = &cobra.Command{
	Use:   "worker-cmd",
	Short: "Serve worker application",
	Long: `A longer description that spans multiple lines and likely contains
			examples and usage of using your application. For example:

			Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := config.GetConfig(validator.New())
		if err != nil {
			panic(fmt.Errorf("Config file invalidate with error: %w", err))
		}

		initWorkerServer(c)
	},
}

func initWorkerServer(c *config.Config) {
	s := asynq.NewServer(
		asynq.RedisClientOpt{Addr: fmt.Sprintf("%s:%s", c.RedisHost, c.RedisPort), DB: c.RedisDb},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				worker.QUEUE_PRIORITY_CRITICAL: 6,
				worker.QUEUE_PRIORITY_DEFAULT:  3,
				worker.QUEUE_PRIORITY_LOW:      1,
			},
			// See the godoc for other configuration options
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.HandleFunc(email_delivery.DELIVERY_EMAIL_QUEUE, email_delivery.HandleEmailDeliveryTask)

	if err := s.Run(mux); err != nil {
		log.Fatal().Msgf("Could not run delivery email worker: %v", err)
	}
	log.Info().Msg("Email delivery worker is running")
}

func init() {
	rootCmd.AddCommand(workerCmd)
}
