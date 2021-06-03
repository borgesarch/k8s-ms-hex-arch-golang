package commons

import (
	d "clean-code-golang/core/infrastructure/data"
	e "clean-code-golang/core/infrastructure/encrypt"
	"encoding/json"

	"go.uber.org/dig"
)

type Config struct {
	Prefix      string
	EncrityHash e.EncrityHash
	Connection  d.Connection
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

	container.Provide(func(cfg *Config) d.Connection {
		return d.Connection(1)
	})

	container.Provide(func(cfg *Config) e.EncrityHash {
		return e.EncrityHash(1)
	})

	return container
}
