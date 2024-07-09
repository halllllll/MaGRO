package config

import "github.com/caarlos0/env/v11"

// env list
type Env string

const (
	EnvDevelopment Env = "dev"
	EnvProductin   Env = "prod"
)

type Config struct {
	Env        Env    `env:"MAGRO_ENV" envDefault:"dev"`
	Port       int    `env:"GO_APP_PORT,required"`
	DBPort     int    `env:"DBPORT,required"`
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
	DBName     string `env:"DB_NAME,required"`
	ClientId   string `env:"ENTRA_CLIENT_ID,required"`
	// below 2 envs only for temporary implement untill asign entra app to user manager privilege
	// ClientSecret string `env:"CLIENT_SECRET"`
	// TenantId string `env:"TENANT_ID"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
