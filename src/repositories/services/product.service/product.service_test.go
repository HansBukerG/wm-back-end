package product_service_test

import (
	"testing"

	"github.com/HansBukerG/wm-back-end/src/repositories/services/product.service"
	"github.com/HansBukerG/wm-back-end/src/utils"
)

func TestRead(t *testing.T)()  {
	products,err := product_service.Read()

	if err != nil {
		t.Error("There is an error in call Read(): " + err.Error())
		t.Fail()
	}
	utils.PrintSlice(products)
	t.Log("Success!")
}

func TestReadById(t *testing.T){
	/*
	Example filters
	*/
	filter := 123
	products,err := product_service.ReadById(filter)
	if err != nil {
		t.Error("Error in query for products")
		t.Fail()
	}
	if len(products) == 0 {
		t.Error("Query has 0 values inside")
		t.Fail()
	}else{
		t.Log("Success!")
	}
}

func TestReadByString(t *testing.T){
	/*
	Example filters
	*/
	// filter := "ooy"
	// filter := "asdfdsa" //palindrome
	// filter := "vqhev"
	filter := "asdfdsa"
	// filter := "121"

	products,err := product_service.ReadByString(filter)

	if err != nil {
		t.Error("Error in query for products")
		t.Fail()
	}
	if len(products) == 0 {
		t.Error("Query has 0 values inside")
		t.Fail()
	}else{
		t.Log("Success!")
	}
}

