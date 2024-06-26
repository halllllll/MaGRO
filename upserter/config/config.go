package config

import "github.com/caarlos0/env/v11"

// TODO: 不要
type Env string

const (
	EnvDevelopment Env = "dev"
	EnvProductin   Env = "prod"
)

type Config struct {
	Env        Env    `env:"MAGRO_ENV" envDefault:"dev"`
	DBPort     int    `env:"DBPORT,required"`
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
	DBName     string `env:"DB_NAME,required"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
