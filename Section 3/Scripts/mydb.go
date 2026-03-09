package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Customer struct {
	Id          int
	name        string
	total_spent float64
}

func main() {

	// Connect to DB...
	cfg := mysql.NewConfig()

	cfg.User = "root"
	cfg.Passwd = "11h2838*56K17"
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "practice"

	// Get a database handle...
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	fmt.Println("connected!")

	customers, err := Get_Customers()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Customers found: %v\n", customers)

	customer, err := Customer_By_ID(4)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Customer found: %v\n", customer)

	// i, err := Add_Customer(Customer{
	// 	name:        "Hussin",
	// 	total_spent: 12500.054,
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("ID of added customer is: %v\n", i)
	i, err := Delete_Customer(6)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID of deleted customer is: %v\n", i)
}

func Get_Customers() ([]Customer, error) {

	var customers []Customer

	rows, err := db.Query("SELECT * FROM Customers")
	if err != nil {
		return nil, fmt.Errorf("Error in Getting data")
	}

	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var custom Customer
		if err := rows.Scan(&custom.Id, &custom.name, &custom.total_spent); err != nil {
			return nil, fmt.Errorf("Error in Getting data")
		}
		customers = append(customers, custom)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error in Getting data")
	}
	return customers, nil
}

func Customer_By_ID(id int) (Customer, error) {
	var customer Customer

	row := db.QueryRow("SELECT * FROM Customers WHERE customer_id = ?", id)
	if err := row.Scan(&customer.Id, &customer.name, &customer.total_spent); err != nil {
		if err == sql.ErrNoRows {
			return customer, fmt.Errorf("Customer_By_ID %d: no such customer", id)
		}
		return customer, fmt.Errorf("Customer_By_ID %d: %v\n", id, err)
	}
	return customer, nil
}

func Add_Customer(customer Customer) (int64, error) {

	result, err := db.Exec("INSERT INTO Customers(name, total_spent) VALUES(?,?)", customer.name, customer.total_spent)
	if err != nil {
		return 0, fmt.Errorf("Add_Customer: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Add_Customer: %v", err)
	}
	return id, nil

}

func Delete_Customer(id int64) (int64, error) {

	result, err := db.Exec("DELETE FROM Customers WHERE customer_id = ?", id)
	if err != nil {
		return 0, fmt.Errorf("Add_Customer: %v", err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("Add_Customer: %v", err)
	}

	return id, nil
}
