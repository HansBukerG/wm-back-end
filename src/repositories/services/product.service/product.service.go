package product_service

import (
	"log"

	product_repository "github.com/HansBukerG/wm-back-end/src/repositories/product.repository"
	"github.com/HansBukerG/wm-back-end/src/utils"

	model "github.com/HansBukerG/wm-back-end/src/models"
)

func Read() (model.Products, error) {
	products, err := product_repository.ReadProducts()

	if err != nil {
		log.Printf("There is an error in call ReadProducts():" + err.Error())
		return nil, err
	}
	utils.PrintSlice(products)
	return products, err
}

func ReadById(id int) (model.Products, error) {
	var products model.Products
	product, err := product_repository.ReadById(id)
	if err != nil {
		log.Printf("Error in call ReadById(): " + err.Error())
		return nil, err
	} else {
		products = append(products, &product)
	}
	return products, err
}

func ReadByString(filter string) (model.Products, error) {
	var products model.Products
	var err error

	channelProducts, errChan := make(chan model.Products), make(chan error)

	go product_repository.ChannelReadByString(filter, channelProducts, errChan)

	products, err = <-channelProducts, <-errChan

	if err != nil {
		log.Printf("Error in query for products")
		return nil, err
	}
	if len(products) == 0 {
		log.Printf("Query has found 0 documents")
		return nil, nil
	}

	return products, err
}
