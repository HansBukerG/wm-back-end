package controller

import (
	"encoding/json"
	"log"
	"net/http"

	product_service "github.com/HansBukerG/wm-back-end/src/repositories/services/product.service"
	"github.com/HansBukerG/wm-back-end/src/utils"
	"github.com/gorilla/mux"
)

func GetProductsDefault(writer http.ResponseWriter, request *http.Request) {
	var status int
	products, err := product_service.Read()
	if err != nil {
		log.Printf("There is an error in call Read(): " + err.Error())
		status = http.StatusNotFound
		products = append(products, utils.EmptyProduct())
	} else {
		status = http.StatusAccepted
	}

	utils.PrintSlice(products)
	response, _ := json.Marshal(products)
	executeResponse(writer, status, response)
}

func GetProductByString(writer http.ResponseWriter, request *http.Request) {
	var status int
	//firt step, i receive the data
	varsRequest := mux.Vars(request)
	searchString := varsRequest["searchString"]

	//i call product_service to get my data from MongoDB
	RegisterFound, err := product_service.SearchByString(searchString)
	if err != nil {
		// RegisterFound = append(RegisterFound, utils.EmptyProduct())
		status = http.StatusNotFound
	}
	if len(RegisterFound) == 0 {
		// RegisterFound = append(RegisterFound, utils.EmptyProduct())
		status = http.StatusNotFound

	} else {
		status = http.StatusAccepted
	}

	response, _ := json.Marshal(RegisterFound)
	executeResponse(writer, status, response)
}

func executeResponse(writer http.ResponseWriter, status int, response []byte) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write(response)
}
