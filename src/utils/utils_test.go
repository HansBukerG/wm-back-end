package utils_test

import (
	"fmt"
	"log"
	"testing"

	model "github.com/HansBukerG/wm-back-end/src/models"
	"github.com/HansBukerG/wm-back-end/src/utils"
)

var slice_A model.Products
var slice_B model.Products
var slice_C model.Products
var slice_D model.Products

var product_A = model.Product{
	Id:          1,
	Brand:       "test",
	Description: "test",
	Image:       "test",
	Price:       10000,
}
var product_B = model.Product{
	Id:          2,
	Brand:       "test",
	Description: "test",
	Image:       "test",
	Price:       10000,
}
var product_B_result = model.Product{
	Id:                  2,
	Brand:               "test",
	Description:         "test",
	Image:               "test",
	Price:               5000,
	Discount_percentaje: 50,
	Original_price:      10000,
}
var product_C = model.Product{
	Id:          34,
	Brand:       "asdfdsa",
	Description: "test",
	Image:       "test",
	Price:       10000,
}
var product_D = model.Product{
	Id:          45,
	Brand:       "test",
	Description: "test",
	Image:       "test",
	Price:       10000,
}
var product_E = model.Product{
	Id:          5,
	Brand:       "asddsa dsaasd",
	Description: "asddsa dsaasd",
	Image:       "test",
	Price:       10000,
}
var product_F = model.Product{
	Id:          6,
	Brand:       "test",
	Description: "asdfdsa test",
	Image:       "test",
	Price:       10000,
}
var product_A_2 = model.Product{
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
		{product_A, true},
		{product_A_2, true},
		{product_B, true},
		{product_C, true},
		{product_D, false},
		{product_E, true},
		{product_F, true},
	}

	for _, datum := range testData {
		result := utils.LookForPalindromes(&datum.filter)
		if result != datum.result {
			t.Errorf("LookForPalindromes(%d) FAILED, Expected %v, got %v", datum.filter.Id, datum.result, result)
		}
	}
}

type testSlice struct {
	filter model.Products
}

func TestUnifySlices(t *testing.T) {
	slice_A = append(slice_A, &product_A, &product_B)
	slice_B = append(slice_B, &product_C, &product_D, &product_A_2)
	slice_B = append(slice_B, &product_C, &product_D, &product_A_2)
	test := []testSlice{
		{slice_A},
		{slice_B},
		{slice_C},
		{slice_A},
		{slice_A},
	}

	var result model.Products

	for _, datum := range test {
		result = utils.UnifySlices(result, datum.filter)
	}

	// voidSlice := utils.UnifySlices()

	t.Log(successMessage)
}

type testDiscount struct {
	filter model.Product
	result model.Product
}

func TestApplyDiscountToProduct(t *testing.T) {

	utils.ApplyDiscountToProduct(&product_B)
	if product_B.Original_price != product_B_result.Original_price {
		t.Errorf("ApplyDiscountToProduct(%d) FAILED, Expected %d, got %d", product_B.Id, product_B_result.Original_price, product_B.Original_price)
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
func TestApplyDiscount(t *testing.T) {
	product := model.Product{
		Id:          3,
		Brand:       "test",
		Description: "test",
		Image:       "test",
		Price:       10000,
	}
	slice_A = append(slice_A, &product)

	slice_A = utils.ApplyDiscount(slice_A)

	for _, product := range slice_A {
		fmt.Printf("t: id: %v,brand: %v,description: %v,price: %v\n", product.Id, product.Brand, product.Description, product.Price)
	}
	t.Log(successMessage)
}

func TestFind(t *testing.T) {
	slice_A = append(slice_A, &product_A, &product_B, &product_C, &product_D)

	if utils.Find(&product_C, slice_A) {
		fmt.Printf("parameter found")
	} else {
		fmt.Printf("parameter not found")
	}
	t.Log(successMessage)

}

func TestEmptyProduct(t *testing.T) {
	product := utils.EmptyProduct()
	fmt.Printf("empty product: %s ", product.Description)
	t.Log(successMessage)
}


type testCheckValue struct{
	filter string
	result int
}

func TestCheckValue(t *testing.T) {
	filter := []testCheckValue{
		{"123", 1},
		{"as", 0},
		{"asd", 2},
	}

	for _, datum:= range filter{
		result := utils.CheckValue(datum.filter)
		if result != datum.result{
			t.Errorf("CheckValue(%s) FAILED, Expected %d, got %d", datum.filter, datum.result, result)
		}
	}

	// t.Log(successMessage)
}

func TestAplyDiscount(t *testing.T) {
	slice_A = append(slice_A, &product_A, &product_B, &product_C, &product_D)

	slice_B := utils.ApplyDiscount(slice_A)

	utils.PrintSlice(slice_B)
}
func TestRemoveItem(t *testing.T) {

	slice_A = append(slice_A, &product_A, &product_B, &product_C, &product_D)

	slice_B := utils.RemoveItem(slice_A, 1)

	utils.PrintSlice(slice_B)
}
func TestHasDiscount(t *testing.T) {
	slice_A = append(slice_A, &product_A, &product_B, &product_C, &product_D)

	for _, v := range slice_A {
		if utils.HasDiscount(v) {
			log.Printf("This item has discount")
		} else {
			log.Printf("This item does not have a discount")
		}
	}
}