package main

import (
	"fmt"
	"log"

	".github.com/Luzik-D/BasicCRUD/internal/config"
	".github.com/Luzik-D/BasicCRUD/internal/storage"
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
	st, err := mysql.New()
	if err != nil {
		logger.Fatal(err)
	}

	b := storage.Book{Title: "b", Author: "a"}
	qerr := st.AddBook(b)
	if qerr != nil {
		fmt.Printf("failed to add the book: %s\n", qerr)
	}
	res, _ := st.GetAllBooks()
	fmt.Println(res)

	// init http server
}
