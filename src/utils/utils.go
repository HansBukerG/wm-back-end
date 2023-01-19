package utils

import model "github.com/HansBukerG/wm-back-end/src/models"

func UnifySlices(brand,description model.Products) model.Products {
	var products model.Products
	if len(brand) == 0{
		return description
	}
	if len(description) == 0{
		return brand
	}

	products = brand

	for _, item := range description{
		if !find(item,products){
			products = append(products, item)
		}
	}

	return products
}

func find(product *model.Product, products model.Products ) bool{

	for _, item := range products{
		if item.Id == product.Id{
			return true
		}
	}
	return false
}