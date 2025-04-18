package db

import (
	"context"
	"html/template"
	"log"
	"net/http"
)

func CheckUsernameDB(username string) bool {
	var exists bool
	err := DB.QueryRow(context.Background(),
		"SELECT EXISTS(SELECT 1 FROM users WHERE username=$1)", username).Scan(&exists)
	if err != nil {
		log.Println("ошибка проверки ника:", err)
		return false
	}
	return exists
}

func CheckEmailDB(email string) bool {
	var exists bool
	err := DB.QueryRow(context.Background(),
		"SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", email).Scan(&exists)
	if err != nil {
		log.Println("ошибка проверки email:", err)
		return true // считаем, что "занят", если ошибка
	}
	return exists
}

// SaveUser сохраняет нового пользователя в базу данных
func SaveUser(username, password, email string) error {

	_, err := DB.Exec(
		context.Background(),
		"INSERT INTO users (username, password, email) VALUES ($1, $2, $3)",
		username, password, email,
	)
	return err

}

func FinalStep(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	// Сохраняем в БД
	err := SaveUser(username, password, email)
	if err != nil {
		http.Error(w, "Ошибка сохранения пользователя", http.StatusInternalServerError)
		log.Println("Ошибка БД:", err)
		return
	}

	// Показываем страницу успеха
	data := struct {
		Username string
	}{
		Username: username,
	}

	tmpl, err := template.ParseFiles("templates/success.html")
	if err != nil {
		http.Error(w, "Ошибка шаблона", http.StatusInternalServerError)
		log.Println("Ошибка шаблона:", err)
		return
	}

	tmpl.Execute(w, data)
}
