package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() {

	cfg := mysql.NewConfig()

	cfg.User = "root"
	cfg.Passwd = os.Getenv("MYSQL_PASS")
	cfg.Addr = "127.0.0.1:3306"
	cfg.Net = "tcp"
	cfg.DBName = "practice"

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!")
}

type Customer struct {
	Id    int64
	Name  string
	Total float64
}

func main() {

	Connect()
	c, _ := Get_Customers()
	fmt.Println(c)
}

func Get_Customers() ([]Customer, error) {

	var customers []Customer

	rows, err := db.Query("SELECT * FROM Customers")

	if err != nil {
		return nil, fmt.Errorf("Error in getting Data")
	}

	defer rows.Close()

	for rows.Next() {
		var c Customer
		if err := rows.Scan(&c.Id, &c.Name, &c.Total); err != nil {
			return nil, fmt.Errorf("Error in parsing data")
		}

		customers = append(customers, c)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error in Getting data")
	}

	return customers, nil
}
