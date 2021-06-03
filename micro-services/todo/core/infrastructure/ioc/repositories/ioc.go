package repositories

import (
	t "clean-code-golang/core/domain/todo"
	l "clean-code-golang/core/infrastructure/repositories/todo"

	"encoding/json"

	"go.uber.org/dig"
)

type Config struct {
	Prefix         string
	TodoRepository l.TodoRepository
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

	container.Provide(func(cfg *Config) l.TodoRepository {
		return l.TodoRepository(t.Todo{})
	})

	return container
}
