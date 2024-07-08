// main.go
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"webapp/db"
	"webapp/handlers"
)

func main() {
	// Connect to PostgreSQL
	database := db.ConnectDB()
	defer database.Close()

	// Initialize handlers
	handlers.Init(database)

	// Router
	r := mux.NewRouter()

	// Regular user routes
	r.HandleFunc("/posts", handlers.GetPostsHandler).Methods("GET")

	// Administrator routes
	adminRouter := r.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("/post", handlers.CreatePostHandler).Methods("POST")
	adminRouter.HandleFunc("/post/{id}", handlers.EditPostHandler).Methods("PUT")
	adminRouter.HandleFunc("/post/{id}", handlers.DeletePostHandler).Methods("DELETE")

	// Start server
	log.Println("Starting server on localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
