package db

import (
	"database/sql"
	"time"

	_ "modernc.org/sqlite"
)

var (
	writeDbCon *sql.DB
	readDbCon  *sql.DB
)

// Возвращает синглтон подключения к базе данных для записи
func GetWriteDbCon() *sql.DB {
	if writeDbCon != nil {
		return writeDbCon
	}
	writeDbCon = configureDbCon(1, 1, 0)
	return writeDbCon
}

// Возвращает синглтон подключения к базе данных для чтения
func GetReadDBCon() *sql.DB {
	if readDbCon != nil {
		return readDbCon
	}
	readDbCon = configureDbCon(10, 5, 0, "PRAGMA query_only = ON")
	return readDbCon
}

func configureDbCon(
	maxOpenConns int,
	maxIdleConns int,
	connMaxLifetime time.Duration,
	additionalPragmas ...string,
) *sql.DB {
	conn, err := sql.Open("sqlite", "./data/app.db")
	if err != nil {
		panic("Couldn`t connect to DB.")
	}
	pragmas := []string{
		"PRAGMA foreign_keys = ON",
		"PRAGMA journal_mode = WAL",
		"PRAGMA synchronous = NORMAL",
		"PRAGMA busy_timeout = 5000",
		"PRAGMA cache_size = -20000", // 20 MB
		"PRAGMA temp_store = MEMORY",
		"PRAGMA mmap_size = 2147483648", // 2 GB
		"PRAGMA auto_vacuum = INCREMENTAL",
		"PRAGMA page_size = 8192",
	}
	for _, pragma := range pragmas {
		if _, err := conn.Exec(pragma); err != nil {
			panic("Failed to set pragma.")
		}
	}
	for _, pragma := range additionalPragmas {
		if _, err := conn.Exec(pragma); err != nil {
			panic("Failed to set additional pragmas.")
		}
	}
	conn.SetMaxOpenConns(maxOpenConns)       // Максимум открытых соединений
	conn.SetMaxIdleConns(maxIdleConns)       // Максимум простаивающих соединений
	conn.SetConnMaxLifetime(connMaxLifetime) // Время жизни соединения (0 = без ограничений)
	if err := conn.Ping(); err != nil {
		panic("Failed to ping DB.")
	}
	return conn
}
