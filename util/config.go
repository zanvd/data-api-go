package util

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Environment             string `mapstructure:"ENVIRONMENT"`
	HTTPServerAddress       string `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress       string `mapstructure:"GRPC_SERVER_ADDRESS"`
	GCPBigQueryProjectId    string `mapstructure:"GCP_BIG_QUERY_PROJECT_ID"`
	GCPBigQueryLedgersTable string `mapstructure:"GCP_BIG_QUERY_LEDGERS_TABLE"`
	GCPBigQueryDataSet      string `mapstructure:"GCP_BIG_QUERY_DATASET"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
