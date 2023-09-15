package main

import (
	"fmt"
	"log"

	".github.com/Luzik-D/BasicCRUD/internal/config"
)

func main() {
	fmt.Println("Hello world")

	// init logger
	logger := log.Default()
	logger.Println("Init logger")

	// read config
	cfgFile, err := config.Load()
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(cfgFile)

	// init storage (db connection)

	// init http server
}
