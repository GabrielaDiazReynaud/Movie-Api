package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
func DBInit(){
 var err error
 DB, err= sql.Open("mysql","root:@tcp(127.0.0.1:3306)/netlfixMovies")
 if err != nil {
	panic("Could not open DB")
 }
}