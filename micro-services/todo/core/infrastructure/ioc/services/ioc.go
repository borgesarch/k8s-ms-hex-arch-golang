package services

import (
	a "ms-hex-arch-golang-k8s/core/application/todo"
	t "ms-hex-arch-golang-k8s/core/domain/todo"

	"encoding/json"

	"go.uber.org/dig"
)

type Config struct {
	Prefix      string
	TodoService a.TodoService
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

	container.Provide(func(cfg *Config) a.TodoService {
		return a.TodoService(t.Todo{})
	})

	return container
}
