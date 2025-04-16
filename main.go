package main

import (
	"log"
	"net/http"
	"oh/db"
	"oh/handlers"
)

func main() {

	db.InitDB()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/check-username", handlers.CheckUsername)

	http.HandleFunc("/reg-name", handlers.GetRegForm)
	http.HandleFunc("/reg-form-pass", handlers.GetRegFormPass)

	log.Println("✅ Сервер шуршит!")
	http.ListenAndServe(":8080", nil)
}
