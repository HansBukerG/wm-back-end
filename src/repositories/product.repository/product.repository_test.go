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

var success_message = "Success!"

func TestReadById(t *testing.T) {
	test_id := 123

	_, err := product_repository.ReadById(test_id)

	if err != nil {
		t.Error("Error in query for products")
		t.Fail()
	} else {
		t.Log(success_message)
	}
}

func TestReadByString(t *testing.T) {
	test_filter := "iñmfdpd fqfwt ikpxov"

	substrings := strings.Fields(test_filter)

	var products model.Products
	var err error

	for _, filter := range substrings {
		channelProducts := make(chan model.Products)
		errChan := make(chan error)
		go product_repository.ChannelReadByString(filter, channelProducts, errChan)
		products = utils.UnifySlices(products, <-channelProducts)
		err = <-errChan
	}

	if err != nil {
		t.Error("Error in query for products")
		t.Fail()
	}
	if len(products) == 0 {
		t.Error("Query has found 0 documents")
		t.Fail()
	}
	utils.PrintSlice(products)
	t.Log(success_message)
}

func TestReadProducts(t *testing.T) {
	products, err := product_repository.ReadProducts()

	if err != nil {
		t.Error("There is an error in call ReadProducts():" + err.Error())
		t.Fail()
	}
	utils.PrintSlice(products)
	t.Log(success_message)
}
