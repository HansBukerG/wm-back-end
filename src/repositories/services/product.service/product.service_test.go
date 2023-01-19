package product_service_test

import (
	"fmt"
	"testing"
	"github.com/HansBukerG/wm-back-end/src/repositories/services/product.service"

)

func TestSearchByString(t *testing.T){
	/*
	Example filters
	*/
	// filter := "ooy"
	// filter := "asdfdsa" //palindrome
	// filter := "vqhev"
	filter := "123"
	// filter := "121"

	products,err := product_service.SearchByString(filter)

	if err != nil {
		t.Error("Error in query for products")
		t.Fail()
	}
	if len(products) == 0 {
		t.Error("Query has 0 values inside")
		t.Fail()
	}else{
		for _, product := range products{
			fmt.Printf("t: id: %v,brand: %v,description: %v,price: %v, Discount_price: %v\n", product.Id,product.Brand,product.Description,product.Price,product.Discount_price)
		}
		t.Log("Success!")
	}
}