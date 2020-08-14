package handlers

import (
	"github.com/gorilla/mux"
	"github.com/jnprogrammer/go_microservices/product-api/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

//how standard GO manages sever requests before using a framework like Gorilla or gin
func (products *Products) ServeHTTP(responsewriter http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		products.GetProducts(responsewriter, request)
		return
	}
	// handle an update
	if request.Method == http.MethodPost {
		products.AddProduct(responsewriter, request)
		return
	}

	if request.Method == http.MethodPut {
		products.logger.Println("PUT", request.URL.Path)
		// expect the id in the uri
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(request.URL.Path, -1)

		if len(g) != 1 {
			products.logger.Println("Invalid URI more than one ID ")
			http.Error(responsewriter, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			products.logger.Println("Invalid URI more than one capture group")
			http.Error(responsewriter, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		_, err := strconv.Atoi(idString) // I need to handel this error gracefully
		if err != nil {
			products.logger.Println("Invalid URI unable to convert to number", idString)
			http.Error(responsewriter, "Invalid URI", http.StatusBadRequest)
			return
		}
		products.UpdateProducts(responsewriter, request)
		return
	}

	//catch all for not impalement handlers
	responsewriter.WriteHeader(http.StatusMethodNotAllowed)
}

func (products *Products) GetProducts(responsewriter http.ResponseWriter, request *http.Request) {
	productlist := data.GetProducts()
	err := productlist.ToJSON(responsewriter)
	if err != nil {
		http.Error(responsewriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (products *Products) AddProduct(responsewriter http.ResponseWriter, request *http.Request) {
	products.logger.Println("Handle post product")

	prod := &data.Product{}
	err := prod.FromJSON(request.Body)

	if err != nil {
		http.Error(responsewriter, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

func (products Products) UpdateProducts(responsewriter http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(responsewriter, "Unable to convert id", http.StatusBadRequest)
		return
	}

	products.logger.Println("Handle update product", id)

	prod := &data.Product{}

	err = prod.FromJSON(request.Body)

	if err != nil {
		http.Error(responsewriter, "Unable to unmarshal json", http.StatusBadRequest)
	}

	if err == data.UpdateProduct(id, prod) {
		if err == data.ErrProductNotFound {
			http.Error(responsewriter, "Product not found", http.StatusInternalServerError)
		}
		return
	}
	if err != nil {
		http.Error(responsewriter, "internal error is found", http.StatusInternalServerError)
		return
	}
}
