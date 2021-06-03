package services

import (
	b "clean-code-golang/core/application/auth"
	a "clean-code-golang/core/application/todo"
	t "clean-code-golang/core/domain/todo"

	"encoding/json"

	"go.uber.org/dig"
)

type Config struct {
	Prefix      string
	TodoService a.TodoService
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

	container.Provide(func(cfg *Config) a.TodoService {
		return a.TodoService(t.Todo{})
	})

	return container
}
