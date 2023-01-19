package product_repository

import (
	"context"

	model "github.com/HansBukerG/wm-back-end/src/models"

	"github.com/HansBukerG/wm-back-end/src/database"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
)

var collection = database.GetCollection("products")
var ctx = context.Background()

func Create(product model.Product) error {
	var err error

	_, err = collection.InsertOne(ctx, product)

	if err != nil {
		return err
	}
	return nil
}

func Read() (model.Products, error) {
	var products model.Products

	filter := bson.D{}
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

func ReadById(id int) (model.Product,error) {
	filter := bson.M{"id":id}
	var product model.Product
	err := collection.FindOne(ctx,filter).Decode(&product)
	return product, err
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

func Update(product model.Product, id int) error {
	filter := bson.M{"id": id}

	update := bson.M{
		"$set": bson.M{
			"brand":       product.Brand,
			"description": product.Description,
			"image":       product.Image,
			"price":       product.Price,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}
	return nil
}

func Delete(id int) error {
	filter :=  bson.M{"id": id}

	_,err :=  collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}
	return nil
}
