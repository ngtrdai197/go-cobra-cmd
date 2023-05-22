package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/ngtrdai197/cobra-cmd/config"
	"github.com/spf13/viper"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	viper.AddConfigPath("..")
	viper.AddConfigPath("..")
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w", err))
	}
	c, err := config.GetConfig(validator.New())
	if err != nil {
		panic(fmt.Errorf("config file invalidate with error: %w", err))
	}
	testDB, err = sql.Open(c.DbDriver, c.DbSource)

	if err != nil {
		log.Fatal(err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}
