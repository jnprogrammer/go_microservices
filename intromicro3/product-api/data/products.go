package data

import "time"

// Product defines the structure for an API product

type Product struct {
	ID          int
	Name        string
	Description string
	price       float32
	SKU         string
	CreatedOn   string
	UpdatedOn   string
	DeleteOn    string
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
		ID:          1,
		Name:        "Esspresso",
		Description: "Short and strong Coffee without milk",
		price:       1.00,
		SKU:         "ess100",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
