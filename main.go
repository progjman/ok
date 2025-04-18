package main

import (
	"log"
	"net/http"
	"ok/db"
	"ok/handlers"
)

func main() {

	db.InitDB()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.IndexHandler)

	http.HandleFunc("/username", handlers.GetRegForm)

	http.HandleFunc("/check-username", handlers.CheckUsername)

	http.HandleFunc("/password", handlers.GetRegFormPass)
	http.HandleFunc("/check-password", handlers.CheckPass)

	http.HandleFunc("/email", handlers.GetRegEmail)
	http.HandleFunc("/check-email", handlers.CheckEmail)

	http.HandleFunc("/save-user", handlers.SaveUser)

	log.Println("✅ Сервер шуршит!")
	http.ListenAndServe(":8080", nil)
}
