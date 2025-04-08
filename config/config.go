package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

var (
	instance *Config
	once     sync.Once
)

type Upstream struct {
	Url  string `yaml:"url"`
	Port string `yaml:"port"`
}

type Config struct {
	Listen    int        `yaml:"listen"`
	Protocol  string     `yaml:"protocol"`
	Algorithm string     `yaml:"algorithm"`
	Upstream  []Upstream `yaml:"upstream"`
}

func LoadConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("laeca")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("/etc/laeca/")

		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}

		instance = &Config{}
		if err := viper.Unmarshal(instance); err != nil {
			panic(fmt.Errorf("unable to decode into struct: %w", err))
		}
	})

	return instance
}
