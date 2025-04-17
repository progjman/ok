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

// Проверка на допустимые символы для email
func isValidEmail(email string) bool {
	// Регулярное выражение для проверки email с символом "@"
	ok, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email)
	return ok
}

// Обработчик проверки Email
func CheckEmail(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimSpace(r.URL.Query().Get("email"))

	data := FormData{
		Value:   email,
		Message: "",
		Status:  "",
	}

	if email != "" {
		if len(email) > 100 {
			data.Message = "Email слишком длинный"
			data.Status = "invalid"
		} else if !isValidEmail(email) {
			data.Message = "Неверный формат email"
			data.Status = "invalid"
		} else if db.CheckEmailDB(email) {
			data.Message = "Email уже зарегистрирован"
			data.Status = "invalid"
		} else {
			data.Message = "Email подходит"
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
