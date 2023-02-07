package main

import (
	"fmt"
	"hexhoc/go-examples/config"
	"hexhoc/go-examples/internal/app"
	"log"
)

func main() {
	fmt.Println("STARTING APP")
	c, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	app.Run(c)
	fmt.Println(c)
}
