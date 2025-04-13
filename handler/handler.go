package handler

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func ShowRegisterPage(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "register.html", nil)
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data any) {
	tmplPath := filepath.Join("templates", tmpl)
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Ошибка шаблона: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Ошибка вывода шаблона: "+err.Error(), http.StatusInternalServerError)
	}
}
