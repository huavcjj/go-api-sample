package web

import "github.com/caarlos0/env/v11"

type ConfigGin struct {
	Host             string   `env:"WEB_HOST" envDefault:"0.0.0.0"`
	Port             string   `env:"WEB_PORT" envDefault:"8080"`
	CorsAllowOrigins []string `env:"WEB_CORS_ALLOW_ORIGINS" envDefault:"http://0.0.0.0:8001"`
}

type ConfigEcho struct {
	Host             string   `env:"WEB_HOST" envDefault:"0.0.0.0"`
	Port             string   `env:"WEB_PORT" envDefault:"8080"`
	CorsAllowOrigins []string `env:"WEB_CORS_ALLOW_ORIGINS" envDefault:"http://0.0.0.0:8001"`
}

func NewConfigGin() *ConfigGin {
	var config ConfigGin
	err := env.Parse(&config)
	if err != nil {
		panic(err)
	}
	return &config
}

func NewConfigEcho() *ConfigEcho {
	var config ConfigEcho
	err := env.Parse(&config)
	if err != nil {
		panic(err)
	}
	return &config
}
