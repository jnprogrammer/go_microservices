package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Product defines the structure for an API product

type Product struct {
	ID          int     `json:"id"` //struct tags look up in docs
	Name        string  `json:"name"`
	Description string  `json:"description"`
	price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"` //come back later and receive the data
	DeleteOn    string  `json:"-"`
}

func (product *Product) FromJSON(reader io.Reader) error {
	e := json.NewDecoder(reader)
	e.Decode(product)

	return e.Decode(product)
}

type Products []*Product

/*
ToJSON serializes the contents of the collection to JSON
NewEncoder provides better performance than json.Unmarshal as it doesn't
have to buffer the output into an in memory slice of bytes
this reduces allocations and the overheads of the service

https://golang.org./pkg/encoding/json/#NewEncoder
*/

func (products *Products) ToJSON(writer io.Writer) error {
	e := json.NewEncoder(writer)
	return e.Encode(products)
}

func GetProducts() Products { //returns our product list
	return productList
}

func UpdateProduct(id int, product *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	product.ID = id
	productList[pos] = product

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, 0, ErrProductNotFound
}

func AddProduct(product *Product) {
	product.ID = getNextID()
	productList = append(productList, product)
}

func getNextID() int {
	lastproduct := productList[len(productList)-1]
	return lastproduct.ID + 1
}

// productList is a Hard coded list of products for this
// example
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy Milky coffee",
		price:       1.50,
		SKU:         "lat150",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Esspresso",
		Description: "Short and strong Coffee without milk",
		price:       1.00,
		SKU:         "ess100",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
