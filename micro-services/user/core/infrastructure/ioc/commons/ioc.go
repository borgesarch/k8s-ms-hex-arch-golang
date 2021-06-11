package commons

import (
	"encoding/json"
	d "ms-hex-arch-golang-k8s/core/infrastructure/data"
	e "ms-hex-arch-golang-k8s/core/infrastructure/encrypt"

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
