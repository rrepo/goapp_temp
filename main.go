package main

import (
	"fmt"
	"modules/app/controllers"
	"modules/config"
	"modules/example"
)
func main() {
	fmt.Println(config.Config)
	fmt.Println(example.Example())
	controllers.StartMainServer()
}

