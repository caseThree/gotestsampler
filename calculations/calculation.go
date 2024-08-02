package calculations

import (
	"errors"
	"fmt"

	"github.com/caseThree/sampleserver/stocks"
)

type Calculate struct {
	stocks stocks.IStock
}

func CreateCalculate(stock stocks.IStock) *Calculate {
	return &Calculate{stocks: stock}
}

func (calci *Calculate) Calculate() (float64, error) {
	top2stock := calci.stocks.GetTopTwoStocks()
	if len(top2stock) < 2 {
		fmt.Printf("Cannot calcualte result")
		return 0, errors.New("cannot calculate result as we don't have")
	}
	increase := ((top2stock[0].Price - top2stock[1].Price)/top2stock[1].Price)*100
	return increase, nil
}