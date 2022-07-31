package main

import (
	"fmt"
	"golang-react-todo/router"
	"net/http"
)

func main(){
	r := router.Router()
	fmt.Println("")
	http.ListenAndServe(":9000",r)
}