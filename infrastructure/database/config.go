package database

import "github.com/caarlos0/env/v11"

type ConfigMySQL struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     string `env:"DB_PORT" envDefault:"3306"`
	User     string `env:"DB_USER" envDefault:"app"`
	Password string `env:"DB_PASSWORD" envDefault:"password"`
	Database string `env:"DB_NAME" envDefault:"api_database"`
	Driver   string `env:"DB_DRIVER" envDefault:"mysql"`
}

type ConfigPostgres struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Driver   string
}

type ConfigSQLite struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Driver   string
}

func NewConfigMySQL() *ConfigMySQL {
	var config ConfigMySQL
	err := env.Parse(&config)
	if err != nil {
		panic(err)
	}
	return &config
}

func NewConfigPostgres() *ConfigPostgres {
	var config ConfigPostgres
	err := env.Parse(&config)
	if err != nil {
		panic(err)
	}
	return &config
}

func NewConfigSQLite() *ConfigSQLite {
	var config ConfigSQLite
	err := env.Parse(&config)
	if err != nil {
		panic(err)
	}
	return &config
}
