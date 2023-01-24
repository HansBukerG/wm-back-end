package product_repository

import (
	"context"

	model "github.com/HansBukerG/wm-back-end/src/models"
	"github.com/HansBukerG/wm-back-end/src/utils"

	"github.com/HansBukerG/wm-back-end/src/database"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
)

var collection = database.GetCollection("products")
var ctx = context.Background()

func ReadById(id int) (model.Product, error) {
	filter := bson.M{"id": id}
	var product model.Product
	err := collection.FindOne(ctx, filter).Decode(&product)
	return product, err
}

func ChannelReadByString(field string, search string, size int, productsChan chan model.Products, errChan chan error) {
	var products model.Products
	var err error
	for i := 0; i < size; i++ {
		products = append(products, utils.EmptyProduct())
	}
	productsChan <- products
	errChan <- err
}

func ReadByString(field string, search string) (model.Products, error) {
	var products model.Products

	filter := bson.M{field: bson.M{"$regex": search, "$options": "im"}}

	productList, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for productList.Next(ctx) {
		var product model.Product
		err = productList.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}
