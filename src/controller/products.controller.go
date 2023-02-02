package controller

import (
	"encoding/json"
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

	RegisterFound.SortSlice()

	response, _ := json.Marshal(RegisterFound)
	executeResponse(writer, status, response)
}

func GetProductByString(writer http.ResponseWriter, request *http.Request) {
	varsRequest := mux.Vars(request)
	filterValue := strings.Trim(varsRequest["searchString"], " ")
	var productsResponse model.Products
	var err error

	filterValue = strings.ToLower(filterValue)
	filterList := strings.Fields(filterValue)

	for _, filter_item := range filterList {
		var productRange model.Products

		switch checkValue := utils.CheckValue(filter_item); checkValue {
		case 1:
			id_int, _ := strconv.Atoi(filter_item)
			productRange, err = product_service.ReadById(id_int)
		case 2:
			productRange, err = product_service.ReadByString(filter_item)
		case 0:
			productRange, err = nil, nil
		}

		productsResponse = utils.UnifySlices(productsResponse, productRange)
	}


	
	productsResponse.SortSlice()

	utils.PrintSlice(productsResponse)
	RegisterFound, status := utils.CheckProducts(productsResponse, err)

	response, _ := json.Marshal(RegisterFound)
	executeResponse(writer, status, response)
}

func executeResponse(writer http.ResponseWriter, status int, response []byte) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write(response)
}
