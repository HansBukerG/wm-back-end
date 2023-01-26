package utils

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	model "github.com/HansBukerG/wm-back-end/src/models"
)

func UnifySlices(brand, description model.Products) model.Products {
	var products model.Products
	if len(brand) == 0 {
		return description
	}
	if len(description) == 0 {
		return brand
	}

	products = brand

	for _, item := range description {
		if !Find(item, products) {
			products = append(products, item)
		}
	}

	return products
}

func Find(product *model.Product, products model.Products) bool {

	for _, item := range products {
		if item.Id == product.Id {
			return true
		}
	}
	return false
}

func IsPalindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func ApplyDiscount(products model.Products) model.Products {
	for _, item := range products {
		item.Discount_percentaje = 50
		item.Original_price = item.Price
		item.Price = item.Price / 2
	}
	return products
}

// func ApplyDiscountToProduct(product *model.Product) *model.Product{
// 	product.Discount_percentaje = 50
// 	product.Original_price = product.Price
// 	product.Price = product.Price / 2
// 	return product
// }

func EmptyProduct() *model.Product {
	product := model.Product{
		Id:          0,
		Brand:       "",
		Description: "",
		Image:       "",
		Price:       0,
	}

	return &product
}

func PrintSlice(slice model.Products) {
	log.Printf("Collection with %d values!", len(slice))
}

func CheckProducts(products model.Products, err error)(model.Products,int){
	var status int
	if err != nil {
		log.Printf("There is an error in call: " + err.Error())
		status = http.StatusBadRequest
		return nil, status
	}
	if len(products) == 0 {
		log.Printf("Return with 0 data.")
		status = http.StatusNoContent
		return nil, status
	} 
	status = http.StatusAccepted
	return products,status
}

func CheckValue(search string)(int){
	_, err := strconv.Atoi(search)
	search = strings.Trim(search, " ")
	if err == nil { //ITS A NUMBER
		return 1
	}else{ // NOT A NUMBER
		if len(search) > 3 { 
			return 2
		}else{ // Dont accomplish the requeriments
			return 0
		}
	}
}