package services

import (
	b "ms-hex-arch-golang-k8s/core/application/users"

	"encoding/json"

	"go.uber.org/dig"
)

type Config struct {
	Prefix      string
	UserService b.UserService
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

	container.Provide(func(cfg *Config) b.UserService {
		return b.UserService(1)
	})

	return container
}
