package product_service

import (
	"strconv"
	"strings"

	product_repository "github.com/HansBukerG/wm-back-end/src/repositories/product.repository"
	"github.com/HansBukerG/wm-back-end/src/utils"

	model "github.com/HansBukerG/wm-back-end/src/models"
)

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

	return products, err
}

func readById(id int) (model.Product, error) {
	product, err := product_repository.ReadById(id)
	return product, err
}

func readByString(search string) (model.Products, error) {

	field_brand := "brand"
	field_description := "description"

	productsByBrand, err := product_repository.ReadByString(field_brand, search)
	if err != nil {
		return nil, err
	}
	productsByDescription, err := product_repository.ReadByString(field_description, search)
	if err != nil {
		return nil, err
	}
	products := utils.UnifySlices(productsByBrand, productsByDescription)

	return products, err
}
