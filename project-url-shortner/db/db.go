package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

//DB represents the database connection Pool
var DB *sql.DB
//InitDB initializes the database connection
func InitDB(dataSourceName string) error {
	var err error
	DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")
	DB.SetMaxOpenConns(10);
	DB.SetMaxIdleConns(5)
	//return DB, nil
	createTable()
	return nil

}
func createTable(){
	//create URL table
	createURLTable := `CREATE TABLE IF NOT EXISTS urls( id SERIAL PRIMARY KEY, short_url TEXT NOT NULL UNIQUE, long_url TEXT NOT NULL, created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP);`
	_, err := DB.Exec(createURLTable)
	if err != nil {
		panic("failed to create events table")
	}
}
