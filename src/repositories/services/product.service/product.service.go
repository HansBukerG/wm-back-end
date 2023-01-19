package product_service

import (
	"strconv"

	product_repository "github.com/HansBukerG/wm-back-end/src/repositories/product.repository"
	"github.com/HansBukerG/wm-back-end/src/utils"

	model "github.com/HansBukerG/wm-back-end/src/models"
)

func Create(product model.Product) error {
	err := product_repository.Create(product)
	if err != nil {
		return err
	}
	return nil
}

func Read() (model.Products, error){
	// var products model.Products

	products, err :=  product_repository.Read()

 	return products, err
}

func ReadById(id string) (model.Product, error){
	id_int,err := strconv.Atoi(id)
	if err != nil {
		id_int = -1
	}
	product,err := product_repository.ReadById(id_int)
	return product, err
}

func ReadByString(search string) (model.Products, error){
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

func Update(product model.Product, id int) error{

	err:= product_repository.Update(product,id)
	if err != nil {
		return err
	}

	return nil
}

func Delete(id int) error{
	err := product_repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}