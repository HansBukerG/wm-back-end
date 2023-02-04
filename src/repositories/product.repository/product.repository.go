package product_repository

import (
	"context"
	"net/http"

	model "github.com/HansBukerG/wm-back-end/src/models"
	"github.com/HansBukerG/wm-back-end/src/utils"

	"github.com/HansBukerG/wm-back-end/src/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection = database.GetCollection("products")
var ctx = context.Background()

func ReadById(id int) (model.Product, int) {
	filter := bson.M{"id": id}
	var product model.Product
	err := collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		return *product.SetEmptyProduct(), http.StatusNotFound
	}
	if utils.LookForPalindromes(&product) {
		utils.ApplyDiscountToProduct(&product)
	}
	return product, http.StatusAccepted
}

func ChannelReadByString(search string, channel chan model.Products, errChan chan int) {
	products, err := ReadByString(search)
	channel <- products
	errChan <- err
	close(channel)
	close(errChan)
}

func ReadByString(search string) (model.Products, int) {
	var products model.Products

	filter := bson.D{
		{Key: "$or", Value: []interface{}{
			bson.M{"brand": bson.M{"$regex": search, "$options": "im"}},
			bson.M{"description": bson.M{"$regex": search, "$options": "im"}},
		},
		}}
	if len(search) == 0 {
		return nil, http.StatusNotFound
	}
	productList, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, http.StatusNotFound
	}

	for productList.Next(ctx) {
		var product model.Product
		decode := productList.Decode(&product)
		if decode != nil {
			return nil, http.StatusConflict
		}
		if utils.LookForPalindromes(&product) {
			utils.ApplyDiscountToProduct(&product)
		}
		products = append(products, &product)
	}
	return products, http.StatusAccepted
}

func ReadProducts() (model.Products, int) {
	filter := bson.D{}
	options := options.Find().SetLimit(10).SetSkip(1)
	var products model.Products
	collectionRequest, err := collection.Find(ctx, filter, options)

	if err != nil {
		return nil, http.StatusNotFound
	}
	for collectionRequest.Next(ctx) {
		var product model.Product
		err := collectionRequest.Decode(&product)
		if err != nil {
			return nil, http.StatusConflict
		}
		if utils.LookForPalindromes(&product) {
			utils.ApplyDiscountToProduct(&product)
		}
		products = append(products, &product)
	}
	return products, http.StatusAccepted
}
