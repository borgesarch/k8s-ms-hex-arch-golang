package repositories

import (
	u "clean-code-golang/core/domain/users"
	k "clean-code-golang/core/infrastructure/repositories/user"

	"encoding/json"

	"go.uber.org/dig"
)

type Config struct {
	Prefix         string
	UserRepository k.UserRepository
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

	container.Provide(func(cfg *Config) k.UserRepository {
		return k.UserRepository(u.User{})
	})

	return container
}
