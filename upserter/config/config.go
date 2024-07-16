package config

import "github.com/caarlos0/env/v11"

type ServiceName string

const (
	ServiceLgate    ServiceName = "lgate"
	ServiceLoilo    ServiceName = "loilo"
	SeriveMiraiseed ServiceName = "miraiseed"
)

type Config struct {
	DBPort     int    `env:"DBPORT,required"`
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
	// DBHost string `env:"DB_HOSTNAME,required"`
	DBName  string      `env:"DB_NAME,required"`
	Service ServiceName `env:"SERVICE,required"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
