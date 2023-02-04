package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	model "github.com/HansBukerG/wm-back-end/src/models"
	product_service "github.com/HansBukerG/wm-back-end/src/repositories/services/product.service"
	"github.com/HansBukerG/wm-back-end/src/utils"
	"github.com/gorilla/mux"
)

func GetProductsDefault(writer http.ResponseWriter, request *http.Request) {
	products, status := product_service.Read()
	products.SortSlice()
	response, _ := json.Marshal(products)
	executeResponse(writer, status, response)
}

func GetProductByString(writer http.ResponseWriter, request *http.Request) {
	varsRequest := mux.Vars(request)
	filterValue := strings.Trim(varsRequest["searchString"], " ")
	var productsResponse model.Products
	var statusResponse int
	var statusFlag bool
	statusFlag = false

	filterValue = strings.ToLower(filterValue)
	filterList := strings.Fields(filterValue)

	for _, filter_item := range filterList {

		var productRange model.Products
		var statusRequest int

		switch checkValue := utils.CheckValue(filter_item); checkValue {
		case 1:
			id_int, _ := strconv.Atoi(filter_item)
			productRange, statusRequest = product_service.ReadById(id_int)
		case 2:
			productRange, statusRequest = product_service.ReadByString(filter_item)
		case 0:
			productRange, statusRequest = nil, http.StatusNotFound
		}

		if statusRequest == http.StatusAccepted {
			log.Printf("filter: %s has returned with data!", filter_item)
			productsResponse = utils.UnifySlices(productsResponse, productRange)
			statusResponse = statusRequest
			statusFlag = true
		} else {
			log.Printf("Filter: %s has returned with 0 data", filter_item)
			if !statusFlag {
				statusResponse = statusRequest
			}
		}
	}

	productsResponse.SortSlice()

	utils.PrintSlice(productsResponse)
	response, _ := json.Marshal(productsResponse)
	executeResponse(writer, statusResponse, response)
}

func executeResponse(writer http.ResponseWriter, status int, response []byte) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write(response)
}
