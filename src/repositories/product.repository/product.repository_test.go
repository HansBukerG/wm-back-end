package product_repository_test

import (
	// "log"
	// "sort"
	"strings"
	"testing"

	model "github.com/HansBukerG/wm-back-end/src/models"
	product_repository "github.com/HansBukerG/wm-back-end/src/repositories/product.repository"
	"github.com/HansBukerG/wm-back-end/src/utils"
)

func TestReadById(t *testing.T) {
	test_id := 123

	_, err := product_repository.ReadById(test_id)

	if err != nil {
		t.Error("Error in query for products")
		t.Fail()
	} else {
		t.Log("Success!")
	}
}

func TestChannelReadByString(t *testing.T) {
	var field string
	var field2 string
	var search string

	field = "brand"
	field2 = "description"
	search = "zdczs"

	channelProductsByBrand, errChan := make(chan model.Products), make(chan error)
	channelProductsByDescription, errChan2 := make(chan model.Products), make(chan error)
	var products model.Products

	go product_repository.ChannelReadByString(field, search, channelProductsByBrand, errChan)
	go product_repository.ChannelReadByString(field2, search, channelProductsByDescription, errChan2)

	productsByBrand, errBrand := <-channelProductsByBrand, <-errChan
	productsByDescription, errDescription := <-channelProductsByDescription, <-errChan2

	if errBrand != nil {
		t.Error("Error in call ChannelReadByString() for Brand: " + errBrand.Error())
		t.Fail()
	}
	if errDescription != nil {
		t.Error("Error in call ChannelReadByString() for Description: " + errDescription.Error())
		t.Fail()
	}

	t.Log("productsByBrand:")
	utils.PrintSlice(productsByBrand)
	t.Log("productsByDescription:")
	utils.PrintSlice(productsByDescription)

	products = utils.UnifySlices(productsByBrand, productsByDescription)

	t.Log("products Unified:")
	utils.PrintSlice(products)

	t.Log("Success!")
}
func TestChannelReadByString1Collection(t *testing.T) {
	// channelProducts := make(chan model.Products)
	var field string
	// var field2 string
	var search string

	field, search = "brand", "asdfsadaf"
	// field2, search = "description", "asdf"

	channelProductsByBrand, errChan := make(chan model.Products), make(chan error)
	// channelProductsByDescription, errChan2 := make(chan model.Products), make(chan error)

	go product_repository.ChannelReadByString(field, search, channelProductsByBrand, errChan)
	// go product_repository.ChannelReadByString(field2,search,channelProductsByDescription, errChan2)

	productsByBrand, errBrand := <-channelProductsByBrand, <-errChan
	// productsByDescription, errDescription := <-channelProductsByDescription, <-errChan2

	if errBrand != nil {
		t.Error("Error in call ChannelReadByString(): errBrand.Error()")
		t.Fail()
	}

	t.Log("productsByBrand:")
	utils.PrintSlice(productsByBrand)
	// t.Log("productsByDescription:")
	// utils.PrintSlice(productsByDescription)

	t.Log("Success!")
}

func TestReadByString(t *testing.T) {
	// test_filter := "iñmfdpd fqfwt ikpxov"
	// test_field:="brand"

	// substrings := strings.Fields(test_filter)

	// var products model.Products
	// var err error

	// for _, filter := range substrings {
	// 	// products = utils.UnifySlices(products, product_repository.ReadByStringTwo(filter))
	// }

	// sort.Slice(products, func(i, j int) bool {
	// 	return products[i].Id < products[j].Id
	// })

	// // products,err := product_repository.ReadByString(test_field,test_filter)
	// if err != nil {
	// 	t.Error("Error in query for products")
	// 	t.Fail()
	// }
	// if len(products) == 0 {
	// 	t.Error("Query has found 0 documents")
	// 	t.Fail()
	// }
	// utils.PrintSlice(products)
	// t.Log("Success!")
}
func TestReadByStringTwo(t *testing.T) {
	test_filter := "iñmfdpd fqfwt ikpxov"
	// test_field:="brand"

	substrings := strings.Fields(test_filter)

	var products model.Products
	var err error
	// channelProducts := make(chan model.Products)

	for _, filter := range substrings {
		channelProducts := make(chan model.Products)
		errChan := make(chan error)
		go product_repository.ChannelReadByStringTwo(filter, channelProducts, errChan)
		products = utils.UnifySlices(products,<-channelProducts)
		err = <- errChan
	}

	// sort.Slice(products, func(i, j int) bool {
	// 	return products[i].Id < products[j].Id
	// })

	// products,err := product_repository.ReadByString(test_field,test_filter)
	if err != nil {
		t.Error("Error in query for products")
		t.Fail()
	}
	if len(products) == 0 {
		t.Error("Query has found 0 documents")
		t.Fail()
	}
	utils.PrintSlice(products)
	t.Log("Success!")
}

func TestReadProducts(t *testing.T) {

	products, err := product_repository.ReadProducts()

	if err != nil {
		t.Error("There is an error in call ReadProducts():" + err.Error())
		t.Fail()
	}
	utils.PrintSlice(products)
	t.Log("Success!")

}
