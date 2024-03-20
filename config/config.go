package config

import "time"

type Config struct {
	Pre string
}

var C *Config

func InitConfig() {
	time.Sleep(3 * time.Second)
	C = &Config{Pre: "123"}
}
