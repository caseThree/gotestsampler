package calculations

import (
	"testing"

	"github.com/caseThree/sampleserver/mocks"
	"github.com/caseThree/sampleserver/stocks"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/suite"
)

const (
	DB_URL = "127.0.0.1"
	DB_PORT = "3306"
	DB_DATABASE_NAME = "stocks_test"
	DB_USER = "root"
	DB_PASSWORD = "abcd"
)

type CalculationSuite struct {
	suite.Suite
}

func TestCalculatorSuite(t *testing.T) {
	suite.Run(t, &CalculationSuite{})
}

func (es *CalculationSuite) TestCalculate() {
	mstock := mocks.IStock{}
	calci := CreateCalculate(&mstock)
	mstock.On("GetTopTwoStocks").Return([]stocks.Stock{
		{
			Id: 1,
			Price: 200,
			Types: "A",
		},
		{
			Id: 2,
			Price: 100,
			Types: "A",
		},
	})
	val, err := calci.Calculate()
	es.Nil(err)
	es.Equal(float64(100), val)
}

func (es *CalculationSuite) TestCalculateNoStocks() {
	mstock := mocks.IStock{}
	calci := CreateCalculate(&mstock)
	mstock.On("GetTopTwoStocks").Return([]stocks.Stock{})
	_, err := calci.Calculate()
	es.NotNil(err)
}