package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func GetRegFormPass(w http.ResponseWriter, r *http.Request) { // Загружает поле ввода
	tmpl, err := template.ParseFiles("templates/reg_form_pass.html")
	if err != nil {
		http.Error(w, "Ошибка шаблона", http.StatusInternalServerError)
		log.Println("Ошибка шаблона:", err)
		return
	}
	tmpl.Execute(w, nil) // Можно передать пустой объект, если нет данных
}
