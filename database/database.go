package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func SetupDB() *sql.DB {
	db := initDB("storage.db")
	migrate(db)

	return db
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	// Here we check for any db errors then exit
	if err != nil {
		panic(err)
	}

	// If we don't get any errors but somehow still don't get a db connection
	// we exit as well
	if db == nil {
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS tasks(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		user_id INTEGER NOT NULL
	);
	
	CREATE TABLE IF NOT EXISTS users(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		email VARCHAR NOT NULL,
		password VARCHAR NOT NULL
	);
	
	INSERT INTO users (email, password)
	VALUES ('test01@test.com', 'test01');

	INSERT INTO users (email, password)
	VALUES ('test02@test.com', 'test02');
    `

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}
