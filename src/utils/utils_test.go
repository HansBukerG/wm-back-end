package utils_test

import (
	"fmt"
	"testing"

	model "github.com/HansBukerG/wm-back-end/src/models"
	"github.com/HansBukerG/wm-back-end/src/utils"
)

var productA = model.Product{
	Id:          1,
	Brand:       "test",
	Description: "test",
	Image:       "test",
	Price:       10000,
}
var productB = model.Product{
	Id:          2,
	Brand:       "test",
	Description: "test",
	Image:       "test",
	Price:       10000,
}
var productBResult = model.Product{
	Id:                  2,
	Brand:               "test",
	Description:         "test",
	Image:               "test",
	Price:               5000,
	Discount_percentaje: 50,
	Original_price:      10000,
}
var productC = model.Product{
	Id:          34,
	Brand:       "asdfdsa",
	Description: "test",
	Image:       "test",
	Price:       10000,
}
var productD = model.Product{
	Id:          45,
	Brand:       "test",
	Description: "test",
	Image:       "test",
	Price:       10000,
}
var productE = model.Product{
	Id:          5,
	Brand:       "asddsa dsaasd",
	Description: "asddsa dsaasd",
	Image:       "test",
	Price:       10000,
}
var productF = model.Product{
	Id:          6,
	Brand:       "test",
	Description: "asdfdsa test",
	Image:       "test",
	Price:       10000,
}
var productA2 = model.Product{
	Id:                  1,
	Brand:               "test",
	Description:         "test",
	Image:               "test",
	Price:               20000,
	Discount_percentaje: 50,
	Original_price:      20000,
}

var successMessage = "Success!"

type testPalindrome struct {
	filter model.Product
	result bool
}

func TestLookForPalindromes(t *testing.T) {
	testData := []testPalindrome{
		{productA, true},
		{productA2, true},
		{productB, true},
		{productC, true},
		{productD, false},
		{productE, true},
		{productF, true},
	}

	for _, datum := range testData {
		result := utils.LookForPalindromes(&datum.filter)
		if result != datum.result {
			t.Errorf("LookForPalindromes(%d) FAILED, Expected %v, got %v", datum.filter.Id, datum.result, result)
		}
	}
}

func TestIsPalindrome(t *testing.T) {

	filter := "abba"
	if utils.IsPalindrome(filter) {
		fmt.Printf("filter: is palindrome")

	} else {
		fmt.Printf("filter: is not palindrome")
	}
	t.Log(successMessage)

}

type testCheckValue struct {
	filter string
	result int
}

func TestCheckValue(t *testing.T) {
	filter := []testCheckValue{
		{"123", 1},
		{"as", 0},
		{"asd", 2},
	}

	for _, datum := range filter {
		result := utils.CheckValue(datum.filter)
		if result != datum.result {
			t.Errorf("CheckValue(%s) FAILED, Expected %d, got %d", datum.filter, datum.result, result)
		}
	}
}



