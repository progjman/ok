package handlers

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"

	"ok/db" // тут должен быть твой модуль для работы с базой
)

func GetRegEmail(w http.ResponseWriter, r *http.Request) {
	// Забираем имя пользователя и пароль из скрытой формы
	username := r.FormValue("username")
	password := r.FormValue("password")

	data := FormData{
		Username: username,
		Password: password,
		Message:  "Введите ваш email",
	}

	tmpl, err := template.ParseFiles("templates/email.html")
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

// Проверка на допустимые символы для email
func isValidEmail(email string) bool {
	// Регулярное выражение для проверки email с символом "@"
	ok, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email)
	return ok
}

// Обработчик проверки Email
func CheckEmail(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimSpace(r.FormValue("email"))
	username := r.FormValue("username")
	password := r.FormValue("password")

	data := FormData{
		Username: username,
		Password: password,
		Value:    email,
		Message:  "",
		Status:   "",
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

	tmpl, err := template.ParseFiles("templates/check-email.html")
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

// SaveUser сохраняет нового пользователя в БД
func SaveUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	err := db.SaveUser(username, password, email)
	if err != nil {
		http.Error(w, "Ошибка сохранения", http.StatusInternalServerError)
		log.Println("Ошибка сохранения пользователя:", err)
		return
	}

	tmpl, err := template.ParseFiles("templates/success.html")
	if err != nil {
		http.Error(w, "Ошибка шаблона", http.StatusInternalServerError)
		log.Println("Ошибка шаблона:", err)
		return
	}

	tmpl.Execute(w, nil)
}
