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

//ActiveBlog handler to call function activeblog
func ActiveBlog(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	dao.BlogActive(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

//UnActiveBlog handler to call function activeblog
func UnActiveBlog(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	dao.BlogUnActive(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}