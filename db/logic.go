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

func СheckEmailDB(email string) bool {
	var exists bool
	err := DB.QueryRow(context.Background(),
		"SELECT EXISTS(SELECT 3 FROM email WHERE email=$3)", email).Scan(&exists)
	if err != nil {
		log.Println("ошибка проверки ника:", err)
		return true
	}
	return exists
}
