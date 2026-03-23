package main

import (
	database "comments-api/config"
	"comments-api/internal/comments"
	"comments-api/internal/user"
	"log"
	"net/http"

	"github.com/rs/cors"
)



func main(){

	db, err := database.InitDb()

	if err != nil {
		log.Fatalf("Error DB: %v", err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	commentRepo := comments.NewRepository(db)
	commentService := comments.NewService(commentRepo)
	commentHandler := comments.NewHandler(commentService)

	userRepo := user.NewRepository(db)
	userHandler := user.NewHandler(userRepo)

	mux.HandleFunc("GET /comments", commentHandler.GetAll)
	mux.HandleFunc("GET /comments/{id}", commentHandler.GetById)
	mux.HandleFunc("POST /comments", commentHandler.Create)
	mux.HandleFunc("DELETE /comments/{id}", commentHandler.Delete)


	mux.HandleFunc("GET /users", userHandler.GetAll)
	mux.HandleFunc("POST /users", userHandler.Register)
	mux.HandleFunc("POST /login", userHandler.Login)

	c := cors.New(cors.Options{
        AllowedOrigins:   []string{"*", "http://127.0.0.1:5500"},
        AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS", "PATCH", "PUT"},
        AllowedHeaders:   []string{"Content-Type", "Authorization"},
        AllowCredentials: true,
    })

	handler := c.Handler(mux)


	log.Println("Inicializando servidor, escuchando en puerto 3000")
	http.ListenAndServe(":3000", handler)
}