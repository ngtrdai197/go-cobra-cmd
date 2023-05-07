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

var clientCmd = &cobra.Command{
	Use:   "client-cmd",
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

		initClient(c)
	},
}

func initClient(c *config.Config) {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: fmt.Sprintf("%s:%s", c.RedisHost, c.RedisPort), DB: c.RedisDb})

	task, err := email_delivery.NewEmailDeliveryTask(123, "Dai Nguyen 2")
	if err != nil {
		log.Fatal().Msgf("Error create task email delivery detail = %v", err)
	}
	// Process the task immediately.
	info, err := client.Enqueue(task, asynq.Queue(worker.QUEUE_PRIORITY_CRITICAL))
	if err != nil {
		log.Fatal().Msgf("Error add task email delivery into queue detail = %v", err)
	}
	log.Printf(" [*] Successfully enqueued task: %+v", info)
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
