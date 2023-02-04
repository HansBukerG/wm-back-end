package product_service

import (
	"net/http"

	product_repository "github.com/HansBukerG/wm-back-end/src/repositories/product.repository"

	model "github.com/HansBukerG/wm-back-end/src/models"
)

func Read() (model.Products, int) {
	products, err := product_repository.ReadProducts()
	return products, err
}

func ReadById(id int) (model.Products, int) {
	var products model.Products
	product, err := product_repository.ReadById(id)
	if err != http.StatusAccepted {
		return nil, err
	} else {
		products = append(products, &product)
	}
	return products, err
}

func ReadByString(filter string) (model.Products, int) {
	var products model.Products
	var err int

	channelProducts, errChan := make(chan model.Products), make(chan int)

	go product_repository.ChannelReadByString(filter, channelProducts, errChan)

	products, err = <-channelProducts, <-errChan

	if err != http.StatusAccepted {
		return nil, http.StatusNotFound
	}
	if len(products) == 0 {
		return nil, http.StatusNotFound
	}

	return products, err
}
