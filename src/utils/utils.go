package utils

import (
	"log"

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

func PrintSlice(slice model.Products){
	for _, product := range slice{
		log.Printf("product: id: %d, brand: %s, description: %s, image: %s, price: %d", product.Id, product.Brand, product.Description, product.Image, product.Price)
	}
}