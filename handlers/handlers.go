package handlers

import (
	"encoding/json"
	"net/http"


	"github.com/greatontime/goblogapi/dao"
	"github.com/greatontime/goblogapi/models"
)

//GetAllBlog handler
func GetAllBlog(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Context-Type","application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin","*")
	payload := dao.GetAllBlog()
	json.NewEncoder(w).Encode(payload)
}

//CreateBlog create blog
func CreateBlog(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var blog models.Blog
	_ = json.NewDecoder(r.Body).Decode(&blog)
	// fmt.Println(task, r.Body)
	dao.InsertOneBlog(blog)
	json.NewEncoder(w).Encode(blog)
}