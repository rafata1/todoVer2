package core

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)
import _ "github.com/go-sql-driver/mysql"

func connectToDB() *sqlx.DB {
	db, err := sqlx.Connect("mysql", "root:1@(localhost:3306)/todo_app_database")
	if err != nil {
		log.Fatalf("Err occured while connecting to db")
		return &sqlx.DB{}
	}
	fmt.Println("Connect to db successfully")
	return db
}
