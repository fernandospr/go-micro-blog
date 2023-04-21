package factories

import (
	"github.com/ilyakaznacheev/cleanenv"

	"go-micro-blog/config"
)

func NewConfig(path string) config.Config {
	var cfg config.Config
	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		panic("Missing config.yml")
	}
	return cfg
}
