package controller

import (
	"encoding/json"
	"log"

	// "log"
	"net/http"
	"strconv"
	"strings"

	model "github.com/HansBukerG/wm-back-end/src/models"
	product_service "github.com/HansBukerG/wm-back-end/src/repositories/services/product.service"
	"github.com/HansBukerG/wm-back-end/src/utils"
	"github.com/gorilla/mux"
)

func GetProductsDefault(writer http.ResponseWriter, request *http.Request) {
	var status int
	products, err := product_service.Read()
	RegisterFound, status := utils.CheckProducts(products, err)
	response, _ := json.Marshal(RegisterFound)
	executeResponse(writer, status, response)
}

func GetProductByString(writer http.ResponseWriter, request *http.Request) {
	varsRequest := mux.Vars(request)
	searchString := varsRequest["searchString"]
	var products model.Products
	var err error

	switch checkValue := utils.CheckValue(searchString); checkValue {
	case 1:
		searchString = strings.Trim(searchString, " ")
		id_int, _ := strconv.Atoi(searchString)
		products, err = product_service.ReadById(id_int)
	case 2:
		searchString = strings.Trim(searchString, " ")
		searchString = strings.ToLower(searchString)
		products, err = product_service.ReadByString(searchString)
	case 0:
		products, err = nil, nil
	}

	if utils.IsPalindrome(searchString) {
		log.Printf("Discount applied to products!")
		products = utils.ApplyDiscount(products)
	}

	utils.PrintSlice(products)
	RegisterFound, status := utils.CheckProducts(products, err)

	response, _ := json.Marshal(RegisterFound)
	executeResponse(writer, status, response)
}

func executeResponse(writer http.ResponseWriter, status int, response []byte) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write(response)
}
