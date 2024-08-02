package stocks

import (
	"database/sql"
	"log"
)

type Stocks struct {
	db *sql.DB
}

type IStock interface {
	GetTopTwoStocks() []Stock
}

func CreateStocks(db *sql.DB) IStock {
	return &Stocks{
		db: db,
	}
}

type Stock struct {
	Id int
	Types string
	Price float64
}

func (stock *Stocks) GetTopTwoStocks() []Stock {
	rows, err := stock.db.Query("select id, type, price from stocks where type = 'A' order by id desc limit 2")
	if err != nil {
		log.Printf("Error in retreiving rows")
		return make([]Stock, 0)
	}
	defer rows.Close()
	result := make([]Stock, 0)
	for rows.Next() {
		var alb Stock
		if err := rows.Scan(&alb.Id, &alb.Types, &alb.Price); err != nil {
			log.Printf("Error in reading row")
		}
		result = append(result, alb)
	}
	return result
}