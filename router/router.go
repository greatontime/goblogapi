package router

import (
	"github.com/greatontime/goblogapi/handlers"

	"github.com/gorilla/mux"
)

//Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/blog",handlers.GetAllBlog).Methods("GET","OPTIONS")
	router.HandleFunc("/api/blog",handlers.CreateBlog).Methods("POST","OPTIONS")
	router.HandleFunc("/api/blog/{id}", handlers.ActiveBlog).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/unactiveblog/{id}", handlers.UnActiveBlog).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteblog/{id}", handlers.DeleteBlog).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteallblog", handlers.DeleteAllBlog).Methods("DELETE", "OPTIONS")

	return router
}