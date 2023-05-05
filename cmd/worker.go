package cmd

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/ngtrdai197/cobra-cmd/config"
	"github.com/ngtrdai197/cobra-cmd/pkg/redis"
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
		fmt.Println("Worker CMD called")

		c, err := config.LoadConfig(validator.New())
		if err != nil {
			panic(fmt.Errorf("Config file invalidate with error: %w", err))
		}
		redis.NewRedisConnection(c.RedisUrl)
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)
}
