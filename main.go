package main

import (
	"chaterley/internal/app/db"
)

func main() {
	database := db.GetDBCon()
	defer database.Close()
}
