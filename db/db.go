package db

import (
	"context"
	"fmt" // ← ты забыл подключить fmt, а оно нужно для строки подключения
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv" // ← не забудь импортировать godotenv, а то Load() не сработает
)

var DB *pgxpool.Pool

func InitDB() {
	// Загружаем переменные окружения из .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Не удалось загрузить .env файл")
	}

	// Чтение данных подключения
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	sslmode := os.Getenv("DB_SSLMODE")

	// Подключение через DSN — pgx это любит
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		user, password, host, dbname, sslmode,
	)

	DB, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("❌ Ошибка подключения к БД: %v", err)
	}

	log.Println("✅ Подключение к базе установлено")

	// Контекст с таймаутом на случай зависания
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Создаём таблицу (временно — для разработки)
	_, err = DB.Exec(ctx, `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,                  -- Идентификатор пользователя
		username TEXT UNIQUE NOT NULL,           -- Никнейм
		password TEXT NOT NULL,                 -- Пароль
		email TEXT UNIQUE NOT NULL              -- Почта
	);

	CREATE TABLE IF NOT EXISTS contacts (
		user_id INT REFERENCES users(id),      -- Внешний ключ на пользователя
		contact_id INT REFERENCES users(id),   -- Внешний ключ на контакт
		PRIMARY KEY (user_id, contact_id)       -- Составной первичный ключ
	);

	CREATE TABLE IF NOT EXISTS ignores (
		user_id INT REFERENCES users(id),      -- Внешний ключ на пользователя
		ignored_id INT REFERENCES users(id),   -- Внешний ключ на игнорируемого пользователя
		PRIMARY KEY (user_id, ignored_id)       -- Составной первичный ключ
	);
	`)

	if err != nil {
		log.Fatalf("❌ Ошибка создания таблицы: %v", err)
	} else {
		log.Println("✅ Таблица готова")
	}
}
