// handlers/admin_handler.go
package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"webapp/models"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	_, err = database.Exec("INSERT INTO posts (title, content) VALUES ($1, $2)", post.Title, post.Content)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]

	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	_, err = database.Exec("UPDATE posts SET title=$1, content=$2 WHERE id=$3", post.Title, post.Content, postID)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Failed to update post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]

	_, err := database.Exec("DELETE FROM posts WHERE id=$1", postID)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Failed to delete post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
