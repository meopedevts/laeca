package config

import (
	"errors"
	"fmt"
	"slices"
	"sync"

	"github.com/spf13/viper"
)

var (
	config *Config
	once   sync.Once

	supportedProtocols  = []string{"tcp", "http"}
	supportedAlgorithms = []string{"round-robin"}
)

var (
	ErrConfigFile            = errors.New("failed to load config file")
	ErrListenNotFound        = errors.New("listen port not found")
	ErrProtocol              = errors.New("protocol not found")
	ErrProtocolNotSupported  = errors.New("protocol not supported")
	ErrAlgorithm             = errors.New("algorithm not found")
	ErrAlgorithmNotSupported = errors.New("algorithm not supported")
	ErrUpstream              = errors.New("upstream not found")
)

type Upstream struct {
	Url  string `yaml:"url"`
	Port string `yaml:"port"`
}

type Config struct {
	Listen              int        `yaml:"listen"`
	Protocol            string     `yaml:"protocol"`
	SupportedProtocols  []string   `yaml:"supportedProtocols"`
	Algorithm           string     `yaml:"algorithm"`
	SupportedAlgorithms []string   `yaml:"supportedAlgorithms"`
	Upstream            []Upstream `yaml:"upstream"`
}

func LoadConfig() *Config {
	once.Do(func() {
		viper.SetConfigName(".laeca")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("/etc/laeca/")

		viper.SetDefault("supportedProtocols", supportedProtocols)
		viper.SetDefault("supportedAlgorithms", supportedAlgorithms)

		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}

		config = &Config{}
		if err := viper.Unmarshal(config); err != nil {
			panic(fmt.Errorf("unable to decode into struct: %w", err))
		}

		if err := config.validate(); err != nil {
			panic(fmt.Errorf("invalid config: %w", err))
		}
	})

	return config
}

func (c *Config) validate() error {
	if c.Listen == 0 {
		return ErrListenNotFound
	}

	if len(c.Upstream) == 0 {
		return ErrUpstream
	}

	if c.Protocol == "" {
		return ErrProtocol
	}
	if !slices.Contains(c.SupportedProtocols, c.Protocol) {
		return ErrProtocolNotSupported
	}

	if c.Algorithm == "" {
		return ErrAlgorithm
	}
	if !slices.Contains(c.SupportedAlgorithms, c.Algorithm) {
		return ErrAlgorithmNotSupported
	}

	return nil
}
