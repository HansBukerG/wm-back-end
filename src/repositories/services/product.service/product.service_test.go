package product_service_test

import (
	"net/http"
	"testing"

	product_service "github.com/HansBukerG/wm-back-end/src/repositories/services/product.service"
)

var successMessage = "Success!"

func TestRead(t *testing.T) {
	_, err := product_service.Read()

	if err != http.StatusAccepted {
		t.Errorf("Error in Read(), status: %d ", err)
	}
}

type dataFilterInt struct {
	filter int
	result int
}

func TestReadById(t *testing.T) {
	productTest := []dataFilterInt{
		{0, http.StatusNotFound},
		{1, http.StatusAccepted},
		{2, http.StatusAccepted},
		{3, http.StatusAccepted},
		{4, http.StatusAccepted},
	}

	for _, datum := range productTest {
		_, result := product_service.ReadById(datum.filter)
		if result != datum.result {
			t.Errorf("ReadById(%d) FAILED, Expected %d, got %d", datum.filter, datum.result, result)
		}
	}
}

type dataFilterString struct {
	filter string
	result int
}

func TestReadByString(t *testing.T) {
	productTest := []dataFilterString{
		{"", http.StatusNotFound},
		{"0", http.StatusNotFound},
		{"asdf", http.StatusAccepted},
		{"dsa", http.StatusAccepted},
	}

	for _, datum := range productTest {
		_, result := product_service.ReadByString(datum.filter)

		if result != datum.result {
			t.Errorf("ReadByString(%s) FAILED, Expected %d, got %d", datum.filter, datum.result, result)
		}
	}
}
