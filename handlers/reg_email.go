package handlers

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"

	"oh/db" // тут должен быть твой модуль для работы с базой
)

func GetRegElail(w http.ResponseWriter, r *http.Request) { // Загружает поле ввода
	tmpl, err := template.ParseFiles("templates/reg_form_email.html")
	if err != nil {
		http.Error(w, "Ошибка шаблона", http.StatusInternalServerError)
		log.Println("Ошибка шаблона:", err)
		return
	}
	tmpl.Execute(w, nil) // Можно передать пустой объект, если нет данных
}

// Проверка на допустимые символы
func isValidEmail(nick string) bool {
	ok, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, nick)
	return ok
}

// Обработчик проверки ника
func CheckEmail(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimSpace(r.URL.Query().Get("email"))

	data := FormData{
		Value:   email,
		Message: "",
		Status:  "",
	}

	if len(email) > 0 {
		if len(email) < 3 || len(email) > 20 {
			data.Message = "From 3 to 20 characters: a-z, A-Z, 0-9, _"
			data.Status = "invalid"
		} else if !isValidEmail(email) {
			data.Message = "Only Latin letters and digits are allowed"
			data.Status = "invalid"
		} else if db.СheckEmailDB(email) {
			data.Message = "This nickname is already taken"
			data.Status = "invalid"
		} else {
			data.Message = "This nickname is available"
			data.Status = "valid"
		}
	}

	tmpl, err := template.ParseFiles("templates/reg_email.html")
	if err != nil {
		http.Error(w, "Ошибка шаблона", http.StatusInternalServerError)
		log.Println("Ошибка шаблона:", err)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Ошибка рендера", http.StatusInternalServerError)
		log.Println("Ошибка рендера:", err)
	}
}
