package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kitayu/go-diaries/api"
	"github.com/kitayu/go-diaries/config"
	db "github.com/kitayu/go-diaries/infrastructure/db"
	server "github.com/kitayu/go-diaries/infrastructure/server"
	"github.com/rs/cors"
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
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	srv := server.NewServer(c.Handler(api.BuildRouter(conn)))

	log.Printf("Serving on localhost:%v\n", config.Config.ServerPort)
	log.Fatal(srv.ListenAndServe())
}
