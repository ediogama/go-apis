package main

import (
	"github.com/ediogama/go-apis/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
