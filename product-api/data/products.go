package data

import (
	"encoding/json"
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
	UpdatedOn   string  `json:"-"` //come back later and recieve the data
	DeleteOn    string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products { //returns our product list
	return productList
}

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
