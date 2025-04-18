package handlers

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"
	// тут должен быть твой модуль для работы с базой
)

func GetRegFormPass(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")

	data := FormData{
		Username: username,
		Message:  "Придумайте пароль",
	}

	tmpl, err := template.ParseFiles("templates/password.html")
	if err != nil {
		http.Error(w, "Ошибка шаблона", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	tmpl.Execute(w, data)
}

// Предкомпилируем регулярки один раз
var (
	uppercasePattern = regexp.MustCompile(`[A-Z]`)                               // хотя бы одна заглавная
	allowedChars     = regexp.MustCompile(`^[a-zA-Z0-9!@#\$%\^&\*\(\)_\+\-=]+$`) // допустимые символы
)

// Проверка пароля на все условия
func isValidPass(pass string) (bool, string) {
	if len(pass) < 7 {
		return false, "Пароль должен быть не короче 7 символов"
	}
	if !allowedChars.MatchString(pass) {
		return false, "Допустимы только латиница, цифры и спецсимволы (!@#$%^&*()_+-=)"
	}
	if !uppercasePattern.MatchString(pass) {
		return false, "Пароль должен содержать хотя бы одну заглавную букву"
	}
	return true, ""
}

// Обработчик проверки пароля
func CheckPass(w http.ResponseWriter, r *http.Request) {
	// Получаем пароль и ник из POST-формы
	username := strings.TrimSpace(r.FormValue("username"))
	password := strings.TrimSpace(r.FormValue("password"))

	data := FormData{
		Username: username, // чтобы не потерять ник
		Value:    password,
		Message:  "",
		Status:   "",
	}

	if password != "" {
		if ok, msg := isValidPass(password); !ok {
			data.Message = msg
			data.Status = "invalid"
		} else {
			data.Message = "Пароль подходит"
			data.Status = "valid"
		}
	}

	tmpl, err := template.ParseFiles("templates/check-password.html")
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
