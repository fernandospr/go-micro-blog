package factories

import (
	"go-micro-blog/config"

	"github.com/ilyakaznacheev/cleanenv"
)

func NewConfig(path string) config.Config {
	var cfg config.Config
	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		panic("Missing config.yml")
	}
	return cfg
}
