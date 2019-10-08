package router

import (
	"github.com/greatontime/goblogapi/handlers"

	"github.com/gorilla/mux"
)

//Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/task",handlers.GetAllBlog).Methods("GET","OPTIONS")

	return router
}