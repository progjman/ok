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
	pass := strings.TrimSpace(r.URL.Query().Get("pass"))

	data := FormData{
		Value:   pass,
		Message: "",
		Status:  "",
	}

	if pass != "" {
		if ok, msg := isValidPass(pass); !ok {
			data.Message = msg
			data.Status = "invalid"
		} else {
			data.Message = "Пароль подходит"
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
