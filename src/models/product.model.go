package model

import (
	"sort"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

// type Product *product
type Products []*Product

func(products Products) SortSlice(){
	sort.Slice(products, func(i, j int) bool {
		return products[i].Discount_percentaje > products[j].Discount_percentaje
	})
}