package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"sync"
)

type Config struct {
	App
	Mysql
}

type (
	App struct {
		Address  string `ini:"address"`
		LogLevel string `ini:"log_level"`
		PageSize int    `ini:"page_size"`
	}

	Mysql struct {
		Connect string `ini:"connect"`
		MaxIdle int    `ini:"max_idle"`
		MaxOpen int    `ini:"max_open"`
	}
)

var (
	Conf *Config
	once sync.Once
)

func NewConfig(env string) {
	once.Do(func() {
		cfg, err := ini.ShadowLoad(fmt.Sprintf("config/%s.ini", env))
		if err != nil {
			log.Fatal(err)
		}

		c := new(Config)
		err = cfg.MapTo(c)
		Conf = c
	})
}
