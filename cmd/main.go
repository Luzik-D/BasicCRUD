package main

import (
	"fmt"
	"log"
	"net/http"

	".github.com/Luzik-D/BasicCRUD/internal/config"
	".github.com/Luzik-D/BasicCRUD/internal/http-server/handlers"
	".github.com/Luzik-D/BasicCRUD/internal/storage/map_storage"
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
	//st, err := mysql.New()
	st, err := map_storage.New()
	if err != nil {
		logger.Fatal(err)
	}

	res, _ := st.GetBooks()
	fmt.Println(res)

	// init http server
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Greeting)
	mux.HandleFunc("/books", handlers.HandleBooks(st))
	mux.HandleFunc("/books/", handlers.HandleBook(st))

	logger.Fatal(http.ListenAndServe(cfg.Address, mux))
}
