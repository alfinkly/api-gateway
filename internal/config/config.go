package config

import (
	"errors"
	"fmt"
	"github.com/goforj/godump"
	"github.com/spf13/viper"
)

type Config struct {
	HTTPAddr     string   `mapstructure:"http_port"`
	DBDSN        string   `mapstructure:"db_dsn"`
	KafkaBrokers []string `mapstructure:"kafka_brokers"`
	GRPCAddr     string   `mapstructure:"grpc_addr"`
}

func Load() (cfg *Config, err error) {
	v := viper.New()

	v.SetDefault("http_port", ":8080")
	v.SetDefault("db_dsn", "postgres://postgres:pass@localhost:5432/app?sslmode=disable")
	v.SetDefault("kafka_brokers", []string{"localhost:9092"})
	v.SetDefault("grpc_addr", "localhost:50051")

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	v.AutomaticEnv()

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into config struct: %w", err)
	}

	if v.ConfigFileUsed() != "" {
		fmt.Printf("Loaded configuration from %s\n", v.ConfigFileUsed())
	}

	godump.Dump(cfg)

	return
}
