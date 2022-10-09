package database

import "database/sql"

func GetConnection() *sql.DB {
	connectionStr := "user=postgres dbname=postgres password=postgres host=localhost port=35432 sslmode=disable"
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err.Error())
	}

	return db
}
