package product_service

import (
	"strconv"

	product_repository "github.com/HansBukerG/wm-back-end/src/repositories/product.repository"
	"github.com/HansBukerG/wm-back-end/src/utils"

	model "github.com/HansBukerG/wm-back-end/src/models"
)

func SearchByString(search string)(model.Products, error){
	var products model.Products
	var product model.Product
	var err error
	if len(search) <= 3{
		//In this case scenario i will point to the readById function
		product, err =  readById(search)
		products = append(products, &product)
	}else{
		products,err = readByString(search)
	}
	if err != nil {
		return nil, err
	}

	return products,err
}

func readById(id string) (model.Product, error){
	id_int,err := strconv.Atoi(id)
	if err != nil {
		id_int = -1
	}
	product,err := product_repository.ReadById(id_int)
	return product, err
}

func readByString(search string) (model.Products, error){

	field_brand := "brand"
	field_description := "description"
	productsByBrand, err := product_repository.ReadByString(field_brand,search)
	if err != nil {
		return nil, err
	}
	productsByDescription, err := product_repository.ReadByString(field_description,search)
	if err != nil {
		return nil, err
	}
	products := utils.UnifySlices(productsByBrand,productsByDescription)
	
	return products,err
}