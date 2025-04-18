package handlers

import (
	"html/template"
	"log"
	"net/http"
	"ok/db"
	"regexp"
	"strings"
)

type FormData struct {
	Username string
	Password string
	Email    string
	Value    string
	Message  string
	Status   string
}

func GetRegForm(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")

	data := FormData{
		Username: username, // Никнейм, переданный из предыдущего шага
		Message:  "Напишите никнейм",
	}

	tmpl, err := template.ParseFiles("templates/username.html")
	if err != nil {
		http.Error(w, "Ошибка шаблона", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	tmpl.Execute(w, data)
}

func isValidUsername(username string) bool {
	ok, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, username)
	return ok
}

func CheckUsername(w http.ResponseWriter, r *http.Request) {
	username := strings.TrimSpace(r.FormValue("username"))

	data := FormData{
		Value: username,
	}

	if username == "" {
		data.Message = "Введите никнейм"
		data.Status = "invalid"
	} else if len(username) < 3 || len(username) > 20 {
		data.Message = "From 3 to 20 characters"
		data.Status = "invalid"
	} else if !isValidUsername(username) {
		data.Message = "Only letters, numbers and _"
		data.Status = "invalid"
	} else if db.CheckUsernameDB(username) {
		data.Message = "This nickname is taken"
		data.Status = "invalid"
	} else {
		data.Message = "Nice! Available"
		data.Status = "valid"
	}

	tmpl, err := template.ParseFiles("templates/check-username.html")
	if err != nil {
		http.Error(w, "Ошибка шаблона", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	tmpl.Execute(w, data)
}
