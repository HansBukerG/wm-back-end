package product_repository_test

import (
	"testing"

	model "github.com/HansBukerG/wm-back-end/src/models"
	product_repository "github.com/HansBukerG/wm-back-end/src/repositories/product.repository"
	"github.com/HansBukerG/wm-back-end/src/utils"
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

func TestChannelReadByString(t *testing.T){
	// channelProducts := make(chan model.Products)
	brand, search := "brand", "asdf"
	
	productsByBrandChannel, errChannel := make(chan model.Products), make(chan error)
	productsByDescriptionChannel, errChannel2 := make(chan model.Products), make(chan error)
	go product_repository.ChannelReadByString(brand,search,9000000,productsByBrandChannel,errChannel)
	go product_repository.ChannelReadByString(brand,search,9000000,productsByDescriptionChannel,errChannel2)
	
	// go product_repository.ChannelReadByString(brand,search,productsByBrandChannel,errChannel)
	// go product_repository.ChannelReadByString(brand,search,productsByDescriptionChannel,errChannel2)
	
	productsByBrand,productsByDescription := <-productsByBrandChannel, <-productsByDescriptionChannel

	t.Log("productsByBrand:")
	utils.PrintSlice(productsByBrand)
	t.Log("productsByDescription:")
	utils.PrintSlice(productsByDescription)
	
	t.Log("Success!")
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

