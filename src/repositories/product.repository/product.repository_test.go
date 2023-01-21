package product_repository_test

import (
	"testing"

	product_repository "github.com/HansBukerG/wm-back-end/src/repositories/product.repository"
)

func TestReadById(t *testing.T){
	test_id:= 123

	_,err:=product_repository.ReadById(test_id)

	if err != nil {
		t.Error("Error in query for products")
		t.Fail()
	}else{
		t.Log("Success!")
	}
}

func TestReadByString(t *testing.T){
	test_filter:= "asdfdsa"
	test_field:="brand" 

	products,err := product_repository.ReadByString(test_field,test_filter)
	if err != nil {
		t.Error("Error in query for products")
		t.Fail()
	}
	if len(products) == 0 {
		t.Error("Query has found 0 documents")
		t.Fail()
	}
	t.Log("Success!")
}

