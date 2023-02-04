package utils_test

import (
	"fmt"
	"log"
	"testing"

	model "github.com/HansBukerG/wm-back-end/src/models"
	"github.com/HansBukerG/wm-back-end/src/utils"
)

var sliceA model.Products
var sliceB model.Products
var sliceC model.Products

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

type testSlice struct {
	filter model.Products
}

func TestUnifySlices(t *testing.T) {
	sliceA = append(sliceA, &productA, &productB)
	sliceB = append(sliceB, &productC, &productD, &productA2)
	sliceB = append(sliceB, &productC, &productD, &productA2)
	test := []testSlice{
		{sliceA},
		{sliceB},
		{sliceC},
		{sliceA},
		{sliceA},
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

	utils.ApplyDiscountToProduct(&productB)
	if productB.Original_price != productBResult.Original_price {
		t.Errorf("ApplyDiscountToProduct(%d) FAILED, Expected %d, got %d", productB.Id, productBResult.Original_price, productB.Original_price)
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
	sliceA = append(sliceA, &product)

	sliceA = utils.ApplyDiscount(sliceA)

	for _, product := range sliceA {
		fmt.Printf("t: id: %v,brand: %v,description: %v,price: %v\n", product.Id, product.Brand, product.Description, product.Price)
	}
	t.Log(successMessage)
}

func TestFind(t *testing.T) {
	sliceA = append(sliceA, &productA, &productB, &productC, &productD)

	if utils.Find(&productC, sliceA) {
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

	// t.Log(successMessage)
}

func TestAplyDiscount(t *testing.T) {
	sliceA = append(sliceA, &productA, &productB, &productC, &productD)

	sliceB := utils.ApplyDiscount(sliceA)

	utils.PrintSlice(sliceB)
}
func TestRemoveItem(t *testing.T) {

	sliceA = append(sliceA, &productA, &productB, &productC, &productD)

	sliceB := utils.RemoveItem(sliceA, 1)

	utils.PrintSlice(sliceB)
}
func TestHasDiscount(t *testing.T) {
	sliceA = append(sliceA, &productA, &productB, &productC, &productD)

	for _, v := range sliceA {
		if utils.HasDiscount(v) {
			log.Printf("This item has discount")
		} else {
			log.Printf("This item does not have a discount")
		}
	}
}
