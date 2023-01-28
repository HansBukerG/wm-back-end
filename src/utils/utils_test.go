package utils_test

import (
	"fmt"
	"log"
	"testing"

	model "github.com/HansBukerG/wm-back-end/src/models"
	"github.com/HansBukerG/wm-back-end/src/utils"
)

func TestUnifySlices(t *testing.T) {
	var slice_A model.Products
	var slice_B model.Products
	var slice_C model.Products

	product_A := model.Product{
		Id:          1,
		Brand:       "test",
		Description: "test",
		Image:       "test",
		Price:       10000,
	}
	product_B := model.Product{
		Id:          2,
		Brand:       "test",
		Description: "test",
		Image:       "test",
		Price:       10000,
	}
	product_C := model.Product{
		Id:          3,
		Brand:       "test",
		Description: "test",
		Image:       "test",
		Price:       10000,
	}
	product_D := model.Product{
		Id:          3,
		Brand:       "test",
		Description: "test",
		Image:       "test",
		Price:       10000,
	}

	slice_A = append(slice_A, &product_A, &product_B)
	slice_B = append(slice_B, &product_C, &product_D)
	slice_C = utils.UnifySlices(slice_A, slice_B)

	for _, product := range slice_C {
		fmt.Printf("t: id: %v,brand: %v,description: %v,price: %v\n", product.Id, product.Brand, product.Description, product.Price)
	}
	t.Log("Success!")
}
func TestIsPalindrome(t *testing.T) {

	filter := "abba"
	if utils.IsPalindrome(filter) {
		fmt.Printf("filter: is palindrome")

	} else {
		fmt.Printf("filter: is not palindrome")
	}
	t.Log("Success!")

}
func TestApplyDiscount(t *testing.T) {
	var slice_A model.Products
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
	t.Log("Success!")
}

func TestFind(t *testing.T) {
	var slice_A model.Products

	product_A := model.Product{
		Id:          1,
		Brand:       "test",
		Description: "test",
		Image:       "test",
		Price:       10000,
	}
	product_B := model.Product{
		Id:          2,
		Brand:       "test",
		Description: "test",
		Image:       "test",
		Price:       10000,
	}
	product_C := model.Product{
		Id:          3,
		Brand:       "test",
		Description: "test",
		Image:       "test",
		Price:       10000,
	}
	product_D := model.Product{
		Id:          4,
		Brand:       "test",
		Description: "test",
		Image:       "test",
		Price:       10000,
	}

	slice_A = append(slice_A, &product_A, &product_B, &product_C, &product_D)

	if utils.Find(&product_C, slice_A) {
		fmt.Printf("parameter found")
	} else {
		fmt.Printf("parameter not found")
	}
	t.Log("Success!")

}

func TestEmptyProduct(t *testing.T) {
	product := utils.EmptyProduct()
	fmt.Printf("empty product: %s ", product.Description)
	t.Log("Success!")
}

func TestCheckValue(t *testing.T) {
	filter := "1"

	log.Printf("value of CheckValue(): %d ", utils.CheckValue(filter))

	t.Log("Success!")
}

func TestAplyDiscount(t *testing.T) {
	var slice_A model.Products

	product_A := model.Product{
		Id:          1,
		Brand:       "test",
		Description: "test",
		Image:       "test",
		Price:       10000,
	}
	product_B := model.Product{
		Id:          2,
		Brand:       "test",
		Description: "test",
		Image:       "test",
		Price:       10000,
	}
	product_C := model.Product{
		Id:          3,
		Brand:       "test",
		Description: "test",
		Image:       "test",
		Price:       10000,
	}
	product_D := model.Product{
		Id:          4,
		Brand:       "test",
		Description: "test",
		Image:       "test",
		Price:       10000,
	}

	slice_A = append(slice_A, &product_A, &product_B, &product_C, &product_D)

	slice_B := utils.ApplyDiscount(slice_A)

	utils.PrintSlice(slice_B)
}
func TestRemoveItem(t *testing.T) {
	var slice_A model.Products

	product_A := model.Product{
		Id:          1,
		Brand:       "test",
		Description: "test",
		Image:       "test",
		Price:       10000,
	}
	product_B := model.Product{
		Id:          2,
		Brand:       "test",
		Description: "test",
		Image:       "test",
		Price:       10000,
	}
	product_C := model.Product{
		Id:          3,
		Brand:       "test",
		Description: "test",
		Image:       "test",
		Price:       10000,
	}
	product_D := model.Product{
		Id:          4,
		Brand:       "test",
		Description: "test",
		Image:       "test",
		Price:       10000,
	}

	slice_A = append(slice_A, &product_A, &product_B, &product_C, &product_D)

	slice_B := utils.RemoveItem(slice_A, 1)

	utils.PrintSlice(slice_B)
}
func TestHasDiscount(t *testing.T) {
	var slice_A model.Products

	product_A := model.Product{
		Id:          1,
		Brand:       "test",
		Description: "test",
		Image:       "test",
		Price:       10000,
	}
	product_B := model.Product{
		Id:          2,
		Brand:       "test",
		Description: "test",
		Image:       "test",
		Price:       10000,
	}
	product_C := model.Product{
		Id:          3,
		Brand:       "test",
		Description: "test",
		Image:       "test",
		Price:       10000,
	}
	product_D := model.Product{
		Id:          4,
		Brand:       "test",
		Description: "test",
		Image:       "test",
		Price:       10000,
	}

	slice_A = append(slice_A, &product_A, &product_B, &product_C, &product_D)

	for _, v := range slice_A {
		if utils.HasDiscount(v) {
			log.Printf("This item has discount")
		} else {
			log.Printf("This item does not have a discount")
		}
	}
}
