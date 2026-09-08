package core

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	HTTP  HTTPConfig  `mapstructure:"http"`
	DB    DBConfig    `mapstructure:"db"`
	Kafka KafkaConfig `mapstructure:"kafka"`
}

type HTTPConfig struct {
	Port int `mapstructure:"port"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSLMode  string `mapstructure:"sslmode"`
}

func (c DBConfig) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.Name, c.SSLMode)
}

type KafkaConfig struct {
	Brokers  []string      `mapstructure:"brokers"`
	Topic    string        `mapstructure:"topic"`
	GroupID  string        `mapstructure:"group_id"`
	Consumer KafkaConsumer `mapstructure:"consumer"`
}

type KafkaConsumer struct {
	EnableAutoCommit   bool   `mapstructure:"enable_auto_commit"`
	AutoCommitInterval string `mapstructure:"auto_commit_interval"`
	SessionTimeout     string `mapstructure:"session_timeout"`
}

func Load(path string) (*Config, error) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(path)
	v.AddConfigPath(".")

	setDefaults(v)

	v.SetEnvPrefix("APP")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("read config: %w", err)
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	return &cfg, nil
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("http.port", 8080)
	v.SetDefault("db.host", "localhost")
	v.SetDefault("db.port", 5432)
	v.SetDefault("db.user", "postgres")
	v.SetDefault("db.password", "postgres")
	v.SetDefault("db.name", "cards_operations")
	v.SetDefault("db.sslmode", "disable")
	v.SetDefault("kafka.brokers", []string{"localhost:9092"})
	v.SetDefault("kafka.topic", "cards-operations")
	v.SetDefault("kafka.group_id", "cards-operations")
	v.SetDefault("kafka.consumer.enable_auto_commit", true)
	v.SetDefault("kafka.consumer.auto_commit_interval", "5s")
	v.SetDefault("kafka.consumer.session_timeout", "45s")
}
