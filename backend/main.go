package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kitayu/go-diaries/api"
	"github.com/kitayu/go-diaries/config"
	db "github.com/kitayu/go-diaries/infrastructure/db"
	server "github.com/kitayu/go-diaries/infrastructure/server"
)

type DiaryRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {
	conn, err := db.NewDB()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer conn.Close()

	srv := server.NewServer(api.BuildRouter(conn))

	log.Printf("Serving on localhost:%v\n", config.Config.ServerPort)
	log.Fatal(srv.ListenAndServe())
}
