package main

import (
	"github.com/767829413/tmp-exc/config"
	"github.com/767829413/tmp-exc/persion"
	"time"
)

type P struct {
	Name string
}

func main() {
	config.InitConfig()
	go persion.PrintPersionName()
	time.Sleep(time.Second * 5)
}
