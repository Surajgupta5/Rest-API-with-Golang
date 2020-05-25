package main

import (
	"fmt"
	"log"
	"net/http"
	"restapiwithgo/router"
)

func main() {
	rou := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	// var host string
	fmt.Println("localhost")
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", rou))
}
