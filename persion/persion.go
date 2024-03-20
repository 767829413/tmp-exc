package persion

import (
	"fmt"

	"github.com/767829413/tmp-exc/config"
)

type Persion struct {
	Name string
}

var pre = config.C.Pre

func PrintPersionName() {
	p := Persion{Name: pre + " my name is persion"}
	fmt.Println(p.Name)
}
