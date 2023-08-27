package env

import (
	"fmt"
	env "github.com/caarlos0/env/v9"
)

var cfg = &Config{}

func GetCfg() *Config {
	return cfg
}

type Config struct {
	Env           Environment `env:"APP_ENV" envDefault:"local"`
	Port          int         `env:"PORT" envDefault:"8080"`
	MySQLUser     string      `env:"MYSQL_USER" envDefault:"app"`
	MySQLPassword string      `env:"MYSQL_PASSWORD" envDefault:"password"`
	MySQLDatabase string      `env:"MYSQL_DATABASE" envDefault:"db"`
	MySQLHost     string      `env:"MYSQL_HOST" envDefault:"localhost"`
	MySQLPort     string      `env:"MYSQL_PORT" envDefault:"3306"`
}

func Build() error {
	if err := env.Parse(cfg); err != nil {
		return err
	}
	if err := cfg.Env.IsValid(); err != nil {
		return err
	}

	return nil
}

type Environment string

const (
	Local Environment = "local"
)

func (e Environment) IsValid() error {
	switch e {
	case Local:
		return nil
	default:
		return fmt.Errorf("invalid environment: %s", e)
	}
}

func (c *Config) IsLocal() bool {
	return c.Env == Local
}
