package product_repository

import (
	"context"

	model "github.com/HansBukerG/wm-back-end/src/models"
	// "github.com/HansBukerG/wm-back-end/src/utils"

	"github.com/HansBukerG/wm-back-end/src/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo"
)

var collection = database.GetCollection("products")
var ctx = context.Background()

func ReadById(id int) (model.Product,error) {
	filter := bson.M{"id":id}
	var product model.Product
	err := collection.FindOne(ctx,filter).Decode(&product)
	return product, err
}

func ChannelReadByString(field string, search string,channel chan model.Products, errChan chan error){
	products,err := ReadByString(field,search)
	channel <- products
	errChan <- err
}

func ReadByString(field string,search string) (model.Products,error){
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

func ReadProducts()(model.Products,error){
	filter := bson.D{}
	options:= options.Find().SetLimit(10).SetSkip(1)
	var products model.Products
	collectionRequest, err := collection.Find(ctx,filter,options)

	if err != nil {
		return nil, err
	}
	for collectionRequest.Next(ctx){
		var product model.Product
		err := collectionRequest.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products,err;
}