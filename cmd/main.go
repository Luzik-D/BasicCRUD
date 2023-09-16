package main

import (
	"fmt"
	"log"

	".github.com/Luzik-D/BasicCRUD/internal/config"
	".github.com/Luzik-D/BasicCRUD/internal/storage/mysql"
)

func main() {
	// init logger
	logger := log.Default()

	// read config
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal(err)
	}

	logger.Printf("Start app with env = %s\n", cfg.Env)
	fmt.Println(cfg)

	// init storage (db connection)
	storage, err := mysql.New()
	if err != nil {
		logger.Fatal(err)
	}

	res, _ := storage.GetAllBooks()
	fmt.Println(res)

	// init http server
}
