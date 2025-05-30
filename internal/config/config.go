package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	HTTPAddr     string
	DBDSN        string
	KafkaBrokers []string
	GRPCAddr     string
}

func Load() (*Config, error) {
	v := viper.New()

	v.SetDefault("http_addr", ":8080")
	v.SetDefault("db_dsn", "postgres://postgres:pass@localhost:5432/app?sslmode=disable")
	v.SetDefault("kafka_brokers", []string{"localhost:9092"})
	v.SetDefault("grpc_addr", "localhost:50051")

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	v.SetEnvPrefix("APP")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into config struct: %w", err)
	}

	if v.ConfigFileUsed() != "" {
		fmt.Printf("Loaded configuration from %s\n", v.ConfigFileUsed())
	}

	return &cfg, nil
}
