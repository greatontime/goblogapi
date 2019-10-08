package handlers

import (
	"encoding/json"
	"net/http"


	"github.com/greatontime/goblogapi/dao"
)

//GetAllBlog handler
func GetAllBlog(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Context-Type","application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin","*")
	payload := dao.GetAllBlog()
	json.NewEncoder(w).Encode(payload)
}