package main

import (
	"fmt"
	"log"

	".github.com/Luzik-D/BasicCRUD/internal/config"
	".github.com/Luzik-D/BasicCRUD/internal/http-server/handlers"
	".github.com/Luzik-D/BasicCRUD/internal/storage/mysql"
	"github.com/gin-gonic/gin"
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
	//st, err := map_storage.New()
	if err != nil {
		logger.Fatal(err)
	}

	res, _ := st.GetBooks()
	fmt.Println(res)

	// init http server
	router := gin.New()
	router.GET("/books", func(c *gin.Context) {
		handlers.Greeting(c.Writer, c.Request)
	})

	// TODO: IMPLEMENT ALL ROUTES AND REPLACE MAIN BRANCH
	/*router.GET("/books", handlers.ShowBooks(st))
	router.GET("/books/:id", handlers.ShowBook(st))
	router.POST("/books", handlers.AddBook(st))
	router.PUT("/books/:id", handlers.ChangeBook(st))
	router.PATCH("/books/:id", handlers.PatchBook(st))
	router.DELETE("/books/:id", handlers.DeleteBook(st))*/

	logger.Fatal(router.Run(cfg.Address))
}
