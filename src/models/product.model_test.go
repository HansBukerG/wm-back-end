package model_test

import (
	"testing"

	model "github.com/HansBukerG/wm-back-end/src/models"
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

type testDiscount struct {
	filter model.Product
	result bool
}

func TestHasDiscount(t *testing.T) {
	sliceA = append(sliceA, &productA, &productB, &productC, &productD)
	filter := []testDiscount{
		{productA, false},
		{productA2, true},
	}
	for _, datum := range filter {
		result := datum.filter.HasDiscount()
		if result != datum.result {
			t.Errorf("HasDiscount(%d) failed, expected %v, got %v", datum.filter.Id, datum.result, result)
		}
	}
}

func TestRemoveItem(t *testing.T) {
	filter := append(sliceA, &productA, &productB, &productC, &productD)
	filter = filter.RemoveItem(1)
	result := append(sliceA, &productB, &productC, &productD)

	if len(filter) != len(result) {
		t.Errorf("RemoveItem() Error, expected %d, got %d", len(result), len(filter))
	}
}

func TestAplyDiscountToProduct(t *testing.T) {
	productA.ApplyDiscountToProduct()

	if !productA.HasDiscount() {
		t.Errorf("HasDiscount() has nos applied a discount")
	}
}

func TestSetEmptyProduct(t *testing.T) {
	var product model.Product

	product.SetEmptyProduct()

	if product.Id != 0 {
		t.Errorf("SetEmptyProduct() has not setted an empty value")
	}
}

type testSlice struct {
	filter model.Products
}

func TestUnifySlices(t *testing.T) {
	sliceA = append(sliceA, &productA, &productB)
	sliceB = append(sliceB, &productC, &productD, &productA2)
	sliceB = append(sliceB, &productC, &productD, &productA2)
	var resultSlice model.Products

	resultSlice = append(resultSlice,&productA, &productB, &productC, &productD )

	test := []testSlice{
		{sliceA},
		{sliceB},
		{sliceC},
		{sliceA},
		{sliceA},
	}
	
	var testSlice model.Products

	for _, datum := range test {
		testSlice = testSlice.AddProducts(datum.filter)
	}

	if len(testSlice) != len(resultSlice){
		t.Errorf("Error in AddProducts(), expected size: %d, got: %d", len(resultSlice), len(testSlice))
	}
}

func TestFind(t *testing.T) {
	sliceA = append(sliceA, &productA, &productB, &productC, &productD)

	if !sliceA.Find(&productC) {
		t.Errorf("Error Find(), expected true, but got false")
	}
}

func TestSortSlice(t *testing.T){
	sliceB = append(sliceB, &productC, &productD, &productA2)
	sliceB.SortSlice()
	for _, datum := range sliceB{
		t.Logf("Result Order: %d", datum.Id)
	}
	sliceB.PrintSlice()
}