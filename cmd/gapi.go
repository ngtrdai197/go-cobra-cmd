package cmd

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/ngtrdai197/cobra-cmd/config"
	gapi "github.com/ngtrdai197/cobra-cmd/pkg/grpc"
	"github.com/spf13/cobra"
)

var gapiCmd = &cobra.Command{
	Use:   "gapi-cmd",
	Short: "Serve gRPC api application",
	Long: `A longer description that spans multiple lines and likely contains
			examples and usage of using your application. For example:

			Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
	Run: func(_ *cobra.Command, _ []string) {
		c, err := config.GetConfig(validator.New())
		if err != nil {
			panic(fmt.Errorf("config file invalidate with error: %w", err))
		}

		initGAPI(c)
	},
}

func initGAPI(c *config.Config) {
	s := gapi.NewServer(c)
	s.Start()
}

func init() {
	rootCmd.AddCommand(gapiCmd)
}
