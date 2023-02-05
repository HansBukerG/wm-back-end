package product_repository_test

import (
	// "log"
	// "sort"

	"net/http"
	// "strings"
	"testing"

	// model "github.com/HansBukerG/wm-back-end/src/models"
	model "github.com/HansBukerG/wm-back-end/src/models"
	product_repository "github.com/HansBukerG/wm-back-end/src/repositories/product.repository"
)

var successMessage = "Success!"

type ProductData struct {
	filter    int
	errResult int
}

func TestReadById(t *testing.T) {
	testData := []ProductData{
		{0, http.StatusNotFound},
		{-1, http.StatusNotFound},
		{-5, http.StatusNotFound},
		{2, http.StatusAccepted},
		{3, http.StatusAccepted},
	}

	for _, datum := range testData {
		_, err := product_repository.ReadById(datum.filter)
		if err != datum.errResult {
			t.Errorf("ReadById(%d) FAILED, Expected %d, got %d", datum.filter, datum.errResult, err)
		}
	}
	t.Log(successMessage)
}

type ProductData2 struct {
	filter    string
	errResult int
}

func TestChannelReadByString(t *testing.T) {
	testData := []ProductData2{
		{"", http.StatusNotFound},
		{"||||", http.StatusAccepted},
		{"notfound", http.StatusAccepted},
		{"iñmfdpd", http.StatusAccepted},
	}

	for _, datum := range testData {
		channelProducts, errTest := make(chan model.Products), make(chan int)
		go product_repository.ChannelReadByString(datum.filter, channelProducts, errTest)
		products := <-channelProducts
		err := <-errTest
		if len(products) == 0 {
			t.Logf("0 data on products.")
		}
		if err != datum.errResult {
			t.Errorf("ReadByString(%s) FAILED, Expected %d, got %d", datum.filter, datum.errResult, err)
		}
	}

}

func TestReadByString(t *testing.T) {
	testData := []ProductData2{
		{"", http.StatusNotFound},
		{"||||", http.StatusAccepted},
		{"notfound", http.StatusAccepted},
		{"iñmfdpd", http.StatusAccepted},
	}

	for _, datum := range testData {
		_, err := product_repository.ReadByString(datum.filter)
		if err != datum.errResult {
			t.Errorf("ReadByString(%s) FAILED, Expected %d, got %d", datum.filter, datum.errResult, err)
		}
	}
}

func TestReadProducts(t *testing.T) {
	products, err := product_repository.ReadProducts()

	if err != http.StatusAccepted {
		t.Errorf("Error in ReadProducts(), status: %d", err)
	}
	products.PrintSlice()
	t.Log(successMessage)
}
