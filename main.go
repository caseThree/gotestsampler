package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/caseThree/sampleserver/calculations"
	"github.com/caseThree/sampleserver/stocks"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_URL = "127.0.0.1"
	DB_PORT = "3306"
	DB_DATABASE_NAME = "stocks"
	DB_USER = "root"
	DB_PASSWORD = "abcd"
)

func main() {
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USER, DB_PASSWORD, DB_URL, DB_PORT, DB_DATABASE_NAME)
	db, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		log.Fatalf("Cannot connect to database: %s\n", err.Error())
	}
	stocks := stocks.CreateStocks(db)
	calculator := calculations.CreateCalculate(stocks)
	value, err := calculator.Calculate()
	if err != nil {
		log.Fatalln("Calculate error")
	}
	fmt.Printf("Calculated increase: %.2f", value)
}