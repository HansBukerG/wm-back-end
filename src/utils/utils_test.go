package utils_test

import (
	"fmt"
	"testing"

	model "github.com/HansBukerG/wm-back-end/src/models"
	"github.com/HansBukerG/wm-back-end/src/utils"
)

func TestUnifySlices(t *testing.T){
	var slice_A model.Products
	var slice_B model.Products
	var slice_C model.Products

	product_A := model.Product{
		Id				: 1,
		Brand			: "test",
		Description		: "test",
		Image			: "test",
		Price			: 10000,
		Discount_price	: 0,
	}
	product_B := model.Product{
		Id				: 2,
		Brand			: "test",
		Description		: "test",
		Image			: "test",
		Price			: 10000,
		Discount_price	: 0,
	}
	product_C := model.Product{
		Id				: 3,
		Brand			: "test",
		Description		: "test",
		Image			: "test",
		Price			: 10000,
		Discount_price	: 0,
	}
	product_D := model.Product{
		Id				: 3,
		Brand			: "test",
		Description		: "test",
		Image			: "test",
		Price			: 10000,
		Discount_price	: 0,
	}

	slice_A = append(slice_A, &product_A,&product_B)
	slice_B = append(slice_B, &product_C,&product_D)
	slice_C = utils.UnifySlices(slice_A,slice_B)

	for _,product := range slice_C{
		fmt.Printf("t: id: %v,brand: %v,description: %v,price: %v, Discount_price: %v\n", product.Id,product.Brand,product.Description,product.Price,product.Discount_price)
	}
	t.Log("Success!")
}
func TestIsPalindrome(t *testing.T){

	filter:="abba"
	if (utils.IsPalindrome(filter)){
		fmt.Printf("filter: is palindrome")
		
	}else{
		fmt.Printf("filter: is not palindrome")
	}
	t.Log("Success!")

}
func TestApplyDiscount(t *testing.T){
	var slice_A model.Products
	product := model.Product{
		Id				: 3,
		Brand			: "test",
		Description		: "test",
		Image			: "test",
		Price			: 10000,
		Discount_price	: 0,
	}
	slice_A = append(slice_A, &product)

	slice_A = utils.ApplyDiscount(slice_A)

	for _,product := range slice_A{
		fmt.Printf("t: id: %v,brand: %v,description: %v,price: %v, Discount_price: %v\n", product.Id,product.Brand,product.Description,product.Price,product.Discount_price)
	}
	t.Log("Success!")
}

func TestFind(t *testing.T){
	var slice_A model.Products

	product_A := model.Product{
		Id				: 1,
		Brand			: "test",
		Description		: "test",
		Image			: "test",
		Price			: 10000,
		Discount_price	: 0,
	}
	product_B := model.Product{
		Id				: 2,
		Brand			: "test",
		Description		: "test",
		Image			: "test",
		Price			: 10000,
		Discount_price	: 0,
	}
	product_C := model.Product{
		Id				: 3,
		Brand			: "test",
		Description		: "test",
		Image			: "test",
		Price			: 10000,
		Discount_price	: 0,
	}
	product_D := model.Product{
		Id				: 4,
		Brand			: "test",
		Description		: "test",
		Image			: "test",
		Price			: 10000,
		Discount_price	: 0,
	}

	slice_A = append(slice_A, &product_A,&product_B,&product_C,&product_D)

	if utils.Find(&product_C, slice_A){
		fmt.Printf("parameter found")
	}else{
		fmt.Printf("parameter not found")
	}

}