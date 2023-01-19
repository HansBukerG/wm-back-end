package product_service_test

import (
	"context"
	"fmt"
	"testing"

	model "github.com/HansBukerG/wm-back-end/src/models"
	"github.com/HansBukerG/wm-back-end/src/repositories/services/product.service"
	"go.mongodb.org/mongo-driver/mongo"
)
var ctx = context.Background()
func TestCreate(t *testing.T){
	product := model.Product{
		Id:          5000,
		Brand:       "test",
		Description: "test",
		Image:       "test",
		Price:       100000,
	}

	err:= product_service.Create(product)

	if err != nil {
		t.Error("La prueba de persistencia de datos de producto ha fallado")
		t.Fail()
	}else{
		t.Log("La prueba finalizo con exito!")
	}
}

func TestRead(t *testing.T){
	
	products,err := product_service.Read()

	if err != nil {
		t.Error("Error in query for products")
		t.Fail()
	}
	if len(products) == 0 {
		t.Error("Query has 0 values inside")
		t.Fail()
	}else{
		for _, product := range products{
			fmt.Printf("t: id: %v,brand: %v,description: %v,price: %v, found_item: %v\n", product.Id,product.Brand,product.Description,product.Price,product.Found_item)
		}
		t.Log("Success!")
	}
}
func TestReadById(t *testing.T){
	idTest := "1"

	product,err := product_service.ReadById(idTest)

	if err != nil { //important
		if err == mongo.ErrNoDocuments{
			t.Log("404: No documents found")
			t.Log("La prueba finalizo con exito!")
		}else{
			t.Error("La prueba de persistencia de datos de producto ha fallado")
			t.Fail()
		}
	}else{
		fmt.Printf("Product: id: %v,brand: %v,description: %v,price: %v, found_item: %v\n", product.Id,product.Brand,product.Description,product.Price,product.Found_item)
		t.Log("La prueba finalizo con exito!")
	}
}

func TestReadByString(t *testing.T){
	/*
	Example filters
	*/
	// filter := "ooy"
	// filter := "asdfdsa" //palindrome
	filter := "vqhev"
	
	products,err := product_service.ReadByString(filter)

	if err != nil {
		t.Error("Error in query for products")
		t.Fail()
	}
	if len(products) == 0 {
		t.Error("Query has 0 values inside")
		t.Fail()
	}else{
		for _, product := range products{
			fmt.Printf("t: id: %v,brand: %v,description: %v,price: %v, found_item: %v\n", product.Id,product.Brand,product.Description,product.Price,product.Found_item)
		}
		t.Log("Success!")
	}
}
func TestUpdate(t *testing.T){
	idTest := 5000
	testProduct := model.Product{
		Brand:       "test updated",
		Description: "test updated",
		Image:       "test updated",
		Price:       100000,
	}

	err:= product_service.Update(testProduct,idTest)

	if err != nil {
		t.Error("La prueba de persistencia de datos de producto ha fallado")
		t.Fail()
	}else{
		t.Log("La prueba finalizo con exito!")
	}
}
func TestDelete(t *testing.T){
	idTest := 5000

	err := product_service.Delete(idTest)
	if err != nil {
		t.Error("La prueba de persistencia de datos de producto ha fallado")
		t.Fail()
	}else{
		t.Log("La prueba finalizo con exito!")
	}
}