package services

import (
	b "clean-code-golang/core/application/auth"

	"encoding/json"

	"go.uber.org/dig"
)

type Config struct {
	Prefix      string
	AuthService b.AuthService
}

func Resolver() *dig.Container {

	container := dig.New()

	err := container.Provide(func() (*Config, error) {
		var config Config
		err := json.Unmarshal([]byte(`{"prefix": "[foo] "}`), &config)
		return &config, err
	})

	if err != nil {
		panic(err)
	}

	container.Provide(func(cfg *Config) b.AuthService {
		return b.AuthService(1)
	})

	return container
}
