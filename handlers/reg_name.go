package handlers

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"

	"oh/db" // тут должен быть твой модуль для работы с базой
)

// Структура для данных, передаваемых в шаблон
type FormData struct {
	Value   string // введённый ник
	Message string // сообщение пользователю
	Status  string // valid | invalid
}

func GetRegForm(w http.ResponseWriter, r *http.Request) { // Загружает поле ввода
	tmpl, err := template.ParseFiles("templates/reg_form.html")
	if err != nil {
		http.Error(w, "Ошибка шаблона", http.StatusInternalServerError)
		log.Println("Ошибка шаблона:", err)
		return
	}
	tmpl.Execute(w, nil) // Можно передать пустой объект, если нет данных
}

// Проверка на допустимые символы
func isValidUsername(nick string) bool {
	ok, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, nick)
	return ok
}

// Обработчик проверки ника
func CheckUsername(w http.ResponseWriter, r *http.Request) {
	username := strings.TrimSpace(r.URL.Query().Get("username"))

	data := FormData{
		Value:   username,
		Message: "",
		Status:  "",
	}

	if len(username) > 0 {
		if len(username) < 3 || len(username) > 20 {
			data.Message = "From 3 to 20 characters: a-z, A-Z, 0-9, _"
			data.Status = "invalid"
		} else if !isValidUsername(username) {
			data.Message = "Only Latin letters and digits are allowed"
			data.Status = "invalid"
		} else if db.СheckUsernameDB(username) {
			data.Message = "This nickname is already taken"
			data.Status = "invalid"
		} else {
			data.Message = "This nickname is available"
			data.Status = "valid"
		}
	}

	tmpl, err := template.ParseFiles("templates/reg_name.html")
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
