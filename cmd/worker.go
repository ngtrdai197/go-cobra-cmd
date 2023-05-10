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
			panic(fmt.Errorf("config file invalidate with error: %w", err))
		}

		initWorkerServer(cmd.Context(), c)
	},
}

func initWorkerServer(ctx context.Context, c *config.Config) {
	processor := worker.NewRedisTaskProcessor(asynq.RedisClientOpt{Addr: fmt.Sprintf("%s:%s", c.RedisHost, c.RedisPort), DB: c.RedisDb})

	processor.Start()
	if err := processor.Start(); err != nil {
		log.Fatal().Msgf("could not start task processor: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(workerCmd)
}
