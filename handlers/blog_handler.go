// handlers/blog_handler.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"webapp/models"
)

var database *sql.DB

func Init(db *sql.DB) {
	database = db
}

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := database.Query("SELECT id, title, content FROM posts ORDER BY id DESC")
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	posts := []models.Post{}
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content); err != nil {
			log.Fatal(err)
			http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}

	json.NewEncoder(w).Encode(posts)
}
