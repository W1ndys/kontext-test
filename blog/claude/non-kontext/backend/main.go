package main

import (
	"blog-backend/handler"
	"blog-backend/middleware"
	"blog-backend/model"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize database
	model.InitDB()

	r := mux.NewRouter()

	// --- Public routes ---
	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	}).Methods("GET")

	r.HandleFunc("/api/login", handler.Login).Methods("POST")
	r.HandleFunc("/api/articles", handler.GetArticles).Methods("GET")
	r.HandleFunc("/api/articles/{id}", handler.GetArticle).Methods("GET")
	r.HandleFunc("/api/categories", handler.GetCategories).Methods("GET")
	r.HandleFunc("/api/tags", handler.GetTags).Methods("GET")

	// --- Admin routes (JWT protected) ---
	admin := r.PathPrefix("/api/admin").Subrouter()
	admin.Use(middleware.AuthMiddleware)
	admin.HandleFunc("/articles", handler.CreateArticle).Methods("POST")
	admin.HandleFunc("/articles/{id}", handler.UpdateArticle).Methods("PUT")
	admin.HandleFunc("/articles/{id}", handler.DeleteArticle).Methods("DELETE")
	admin.HandleFunc("/upload", handler.Upload).Methods("POST")

	// --- Static file serving ---
	// Uploaded files
	r.PathPrefix("/uploads/").Handler(
		http.StripPrefix("/uploads/", http.FileServer(http.Dir("data/uploads"))),
	)

	// SPA: serve static assets and fallback to index.html
	r.PathPrefix("/assets/").Handler(http.FileServer(http.Dir("static")))
	r.PathPrefix("/").HandlerFunc(spaHandler)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

// spaHandler serves index.html for any route not matched above (SPA fallback)
func spaHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}
