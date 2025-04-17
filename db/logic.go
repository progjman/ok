package db

import (
	"context"
	"log"
)

func СheckUsernameDB(username string) bool {
	var exists bool
	err := DB.QueryRow(context.Background(),
		"SELECT EXISTS(SELECT 1 FROM users WHERE username=$1)", username).Scan(&exists)
	if err != nil {
		log.Println("ошибка проверки ника:", err)
		return true
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
