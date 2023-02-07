package main

import (
	"fmt"
	"hexhoc/go-examples/config"
	"log"
)

func main() {
	fmt.Println("STARTING APP")
	c, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	fmt.Println(c)
}
