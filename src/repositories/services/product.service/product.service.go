package product_service

import (
	"log"
	"strconv"
	"strings"

	product_repository "github.com/HansBukerG/wm-back-end/src/repositories/product.repository"
	"github.com/HansBukerG/wm-back-end/src/utils"

	model "github.com/HansBukerG/wm-back-end/src/models"
)

func Read()(model.Products,error){
	products,err:= product_repository.ReadProducts()

	if err != nil {
		log.Printf("There is an error in call ReadProducts():" + err.Error())
		return nil,err
	}
	utils.PrintSlice(products)
	return products,err
}

func SearchByString(search string) (model.Products, error) {
	var products model.Products
	var product model.Product
	var err error

	if strings.Trim(search, " ") == "" {
		return nil, err
	}

	id_int, err := strconv.Atoi(search)
	if err == nil { //ITS A NUMBER
		product, err = readById(id_int)
		products = append(products, &product)
	} else { // ITS NOT A NUMBER
		if len(search) > 3 {
			products, err = readByString(strings.ToLower(search))
		} else {
			return nil, err
		}
	}
	if utils.IsPalindrome(search) {
		products = utils.ApplyDiscount(products)
	}

	utils.PrintSlice(products)
	return products, err
}

func readById(id int) (model.Product, error) {
	product, err := product_repository.ReadById(id)
	return product, err
}

func readByString(search string) (model.Products, error) {

	var field string
	var field2 string

	field = "brand"
	field2 = "description"

	channelProductsByBrand, errChan := make(chan model.Products), make(chan error)
	channelProductsByDescription, errChan2 := make(chan model.Products), make(chan error)
	var products model.Products
	var err error

	go product_repository.ChannelReadByString(field, search, channelProductsByBrand, errChan)
	go product_repository.ChannelReadByString(field2, search, channelProductsByDescription, errChan2)

	productsByBrand, errBrand := <-channelProductsByBrand, <-errChan
	productsByDescription, errDescription := <-channelProductsByDescription, <-errChan2

	if errBrand != nil {
		log.Printf("Error in call ChannelReadByString() for Brand: " + errBrand.Error())
		return nil, errBrand
	}
	if errDescription != nil {
		log.Printf("Error in call ChannelReadByString() for Description: " + errDescription.Error())
		return nil, errDescription
	}
	products = utils.UnifySlices(productsByBrand, productsByDescription)

	return products, err
}
