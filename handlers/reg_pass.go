package handlers

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"
	// тут должен быть твой модуль для работы с базой
)

// // Структура для данных, передаваемых в шаблон
// type FormData struct {
// 	Value   string // введённый ник
// 	Message string // сообщение пользователю
// 	Status  string // valid | invalid
// }

func GetRegFormPass(w http.ResponseWriter, r *http.Request) { // Загружает поле ввода
	tmpl, err := template.ParseFiles("templates/reg_form_pass.html")
	if err != nil {
		http.Error(w, "Ошибка шаблона", http.StatusInternalServerError)
		log.Println("Ошибка шаблона:", err)
		return
	}
	tmpl.Execute(w, nil) // Можно передать пустой объект, если нет данных
}

// Проверка на допустимые символы
func isValidPass(pass string) bool {
	ok, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, pass)
	return ok
}

// Обработчик проверки ника
func CheckPass(w http.ResponseWriter, r *http.Request) {
	pass := strings.TrimSpace(r.URL.Query().Get("pass"))

	data := FormData{
		Value:   pass,
		Message: "",
		Status:  "",
	}

	if len(pass) > 0 {
		if len(pass) < 3 || len(pass) > 20 {
			data.Message = "From 3 to 20 characters: a-z, A-Z, 0-9, _"
			data.Status = "invalid"
		} else if !isValidPass(pass) {
			data.Message = "Only Latin letters and digits are allowed"
			data.Status = "invalid"
		} else {
			data.Message = "This password is available"
			data.Status = "valid"
		}
	}

	tmpl, err := template.ParseFiles("templates/reg_pass.html")
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
