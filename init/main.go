package main

import (
	"flag"
	"go-ecommerce/config"
	"go-ecommerce/init/app"
)

var envFlag = flag.String("config", "init/env.toml", "env not found")

func main() {
	flag.Parse()
	c := config.NewConfig(*envFlag)
	app.NewApp(c)
}
