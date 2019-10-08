package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/greatontime/goblogapi/router"
)

func main(){
	r := router.Router()
	fmt.Println("Starting server on the prot 8080...")
	log.Fatal(http.ListenAndServe(":8080",r))
}