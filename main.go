package main

import (
	"log"
	"net/http"

	"github.com/progjman/ok/db"
	"github.com/progjman/ok/handler"
)

func main() {
	db.InitDB()
	http.HandleFunc("/", handler.ShowRegisterPage)
	http.HandleFunc("/check-username", handler.CheckUsername)
	log.Println("Сервер на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
