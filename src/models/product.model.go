package model

import (
	"log"
	"sort"

	"go.mongodb.org/mongo-driver/bson/primitive"
	// "github.com/HansBukerG/wm-back-end/src/utils"
)

type Product struct {
	Id_object           primitive.ObjectID `bson:"_id,omitempty"`
	Id                  int                `json:"id"`
	Brand               string             `json:"brand"`
	Description         string             `json:"description"`
	Image               string             `json:"image"`
	Price               int                `json:"price"`
	Discount_percentaje int                `json:"discount_percentaje"`
	Original_price      int                `json:"original_price"`
}

func (product *Product) SetEmptyProduct() {
	product = &Product{
		Id_object:           [12]byte{0},
		Id:                  0,
		Brand:               "",
		Description:         "",
		Image:               "",
		Price:               0,
		Discount_percentaje: 0,
		Original_price:      0,
	}
}

func (product *Product) HasDiscount() bool {
	return product.Original_price > 0
}

func (product *Product) ApplyDiscountToProduct() {
	product.Discount_percentaje = 50
	product.Original_price = product.Price
	product.Price = product.Price / 2
}

// type Product *product
type Products []*Product

func (products Products) SortSlice() {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Discount_percentaje > products[j].Discount_percentaje
	})
}

func (products Products) RemoveItem(id int) Products {
	for index, value := range products {
		if value.Id == id {
			products = append(products[:index], products[index+1:]...)
			break
		}
	}
	return products
}

func (products Products) Find(product *Product) bool {
	for _, item := range products {
		if item.Id == product.Id {
			return true
		}
	}
	return false
}

func (products Products) AddProducts(newProducts Products) Products {
	if len(newProducts) != 0 {
		for _, item := range newProducts {
			if !products.Find(item) {
				products = append(products, item)
			} else {
				if item.HasDiscount() {
					products = products.RemoveItem(item.Id)
					products = append(products, item)
				}
			}
		}
	}
	return products
}

func (slice Products) PrintSlice() {
	// uncomment for detailed info
	// for _, product := range slice {
	// 	log.Printf("Product with id: %d ", product.Id)
	// }
	log.Printf("Collection returned with %d values!", len(slice))
}
