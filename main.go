package main

import (
	_ "evan-gf/boot"
	_ "evan-gf/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
