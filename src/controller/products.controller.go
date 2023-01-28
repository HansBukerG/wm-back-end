package controller

import (
	"encoding/json"
	// "log"

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
	filter_value := strings.Trim(varsRequest["searchString"]," ")
	var products_response model.Products
	var err error

	filter_value = strings.ToLower(filter_value)
	filter_list := strings.Fields(filter_value)

	for _, filter_item := range filter_list {
		var product_range model.Products

		switch checkValue := utils.CheckValue(filter_item); checkValue {
		case 1:
			id_int, _ := strconv.Atoi(filter_item)
			product_range, err = product_service.ReadById(id_int)
		case 2:
			product_range, err = product_service.ReadByString(filter_item)
		case 0:
			product_range, err = nil, nil
		}

		products_response = utils.UnifySlices(products_response,product_range)
	}

	utils.PrintSlice(products_response)
	RegisterFound, status := utils.CheckProducts(products_response, err)

	response, _ := json.Marshal(RegisterFound)
	executeResponse(writer, status, response)
}

func executeResponse(writer http.ResponseWriter, status int, response []byte) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write(response)
}
