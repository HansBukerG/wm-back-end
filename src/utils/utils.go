package utils

import (
	"log"
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
		} else {
			if HasDiscount(item) {
				products = RemoveItem(products, item.Id)
				products = append(products, item)
			}
		}
	}
	return products
}

func HasDiscount(product *model.Product) bool {
	if product.Original_price > 0 {
		return true
	}
	return false
}

func RemoveItem(slice model.Products, id int) model.Products {
	for index, value := range slice {
		if value.Id == id {
			slice = append(slice[:index], slice[index+1:]...)
			break
		}
	}
	return slice
}

func Find(product *model.Product, products model.Products) bool {

	for _, item := range products {
		if item.Id == product.Id {
			return true
		}
	}
	return false
}

func LookForPalindromes(product *model.Product) bool {
	return (IsPalindrome(strconv.Itoa(product.Id)) ||
		checkFilter(product.Brand) ||
		checkFilter(product.Description))
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

func checkFilter(text string) bool {
	descriptionFields := strings.Fields(text)
	for _, value := range descriptionFields {
		if IsPalindrome(value) {
			return true
		}
	}
	return false
}

func ApplyDiscount(products model.Products) model.Products {
	for _, item := range products {
		item.Discount_percentaje = 50
		item.Original_price = item.Price
		item.Price = item.Price / 2
	}
	return products
}

func ApplyDiscountToProduct(product *model.Product) {
	product.Discount_percentaje = 50
	product.Original_price = product.Price
	product.Price = product.Price / 2
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

func PrintSlice(slice model.Products) {
	//uncomment for detailed info
	// for _, product := range slice {
	// 	log.Printf("Product with id: %d ", product.Id)
	// }
	log.Printf("Collection returned with %d values!", len(slice))
}

func CheckValue(search string) int {
	_, err := strconv.Atoi(search)
	search = strings.Trim(search, " ")
	if err == nil { //ITS A NUMBER
		return 1
	} else { // NOT A NUMBER
		if len(search) >= 3 {
			return 2
		} else { // Dont accomplish the requeriments
			return 0
		}
	}
}
