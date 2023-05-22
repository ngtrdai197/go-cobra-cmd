package cmd

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/ngtrdai197/cobra-cmd/config"
	api "github.com/ngtrdai197/cobra-cmd/pkg/public-api"
	"github.com/spf13/cobra"
)

var publicApiCmd = &cobra.Command{
	Use:   "public-api-cmd",
	Short: "Serve public api application",
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

		initPublicAPI(c)
	},
}

func initPublicAPI(c *config.Config) {
	s := api.NewServer(c)
	s.Start(c.PublicApiAddress)
}

func init() {
	rootCmd.AddCommand(publicApiCmd)
}
