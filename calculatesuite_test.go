package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExampleSuite struct {
	suite.Suite
}

func TestExampleSuite(t *testing.T) {
	suite.Run(t, &ExampleSuite{})
}

func (es *ExampleSuite) SetupSuite() {
	fmt.Println("Setup suite")
	// on suite creation
}

func (es *ExampleSuite) TearDownSuite() {
	fmt.Println("Tear Down suite")
	// on suite destroy
}

func (es *ExampleSuite) SetupTest() {
	fmt.Println("Setup tests")
	// before every test
}

func (es *ExampleSuite) TearDownTest() {
	fmt.Println("Tear Down tests")
	// after every test
}

func (es *ExampleSuite) BeforeTest(suiteName, testName string) {
	fmt.Println("Setup tests before")
	// before every test but we can add optional logic
}

func (es *ExampleSuite) AfterTest(suiteName, testName string) {
	fmt.Println("Tear Down tests after")
	// before every test but we can add optional logic
}

func (es *ExampleSuite) TestTrue() {
	es.True(true)
}