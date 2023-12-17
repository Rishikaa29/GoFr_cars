// db/database.go
package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// InitDB initializes the database connection
func InitDB() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/teoalida_databases")
	if err != nil {
		panic(err)
	}

	// Check if the database connection is successful
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to the database")
}
