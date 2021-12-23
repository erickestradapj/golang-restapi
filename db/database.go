package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// user:password@tcp(localhost:3306)/dbname
const url = "root:123456@tcp(localhost:3306)/goweb_db"

/* ===== CONNECTION VARIABLE ===== */
var db *sql.DB

/* ===== OPEN CONNECTION ===== */
func Connect() {
	connection, err := sql.Open("mysql", url)

	if err != nil {
		panic(err)
	}

	fmt.Println("=> Connection successfully")
	db = connection
}

/* ===== CLOSE CONNECTION ===== */
func Close() {
	db.Close()
}

/* ===== TEST CONNECTION ===== */
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

/* ===== EXISTS TABLE ===== */
func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := db.Query(sql)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	return rows.Next()
}

/* ===== CREATE TABLE ===== */
func CreateTable(schema, name string) {
	if !ExistsTable(name) {
		_, err := db.Exec(schema)
		if err != nil {
			fmt.Println(err)
		}
	}
}

/* ===== TRUNCATE TABLE ===== */
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	db.Exec(sql)
}

/* ===== POLYMORPHISM - EXEC ===== */
func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connect()
	result, err := db.Exec(query, args...)
	Close()

	if err != nil {
		fmt.Println(err)
	}

	return result, err
}

/* ===== POLYMORPHISM - QUERY ===== */
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Connect()
	rows, err := db.Query(query, args...)
	Close()

	if err != nil {
		fmt.Println(err)
	}

	return rows, nil
}
