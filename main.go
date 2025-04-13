package main

import (
	"log"
	"net/http"

	"github.com/progjman/ok/db"
	"github.com/progjman/ok/handler"
)

func main() {
	db.InitDB()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handler.ShowRegisterPage)
	http.HandleFunc("/check-username", handler.CheckUsername)

	http.HandleFunc("/register/password-step", handler.RegisterPassword)
	http.HandleFunc("/check-password", handler.CheckPassword)

	log.Println("Сервер на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
