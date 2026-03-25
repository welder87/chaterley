package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var dbCon *sql.DB

// Возвращает синглтон подключения к базе данных
func GetDBCon() *sql.DB {
	if dbCon != nil {
		return dbCon
	}

	// Настравиваем подключение к БД
	connect, err := sql.Open("sqlite", "./data/app.db")

	if err != nil {
		panic("Couldn`t connect to DB.")
	}
	dbCon = connect

	// Настраиваем пул соединений
	dbCon.SetMaxOpenConns(10)   // Максимум открытых соединений
	dbCon.SetMaxIdleConns(5)    // Максимум простаивающих соединений
	dbCon.SetConnMaxLifetime(0) // Время жизни соединения (0 = без ограничений)

	// Проверяем соединение к БД
	if err = dbCon.Ping(); err != nil {
		panic("Failed to ping DB.")
	}
	return dbCon
}
