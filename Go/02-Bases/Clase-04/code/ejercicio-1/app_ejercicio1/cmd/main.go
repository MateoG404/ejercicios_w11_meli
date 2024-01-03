package main

import (
	"app_ejercicio1/app_ejercicio1/internal"
	"fmt"
)

func main() {
	fmt.Println("Inicio del programa")

	// Create new products and add to sliceProducts

	internal.SliceProducts = append(internal.SliceProducts, internal.Product{ID: 1, Name: "iPhone 12", Price: 999.99, Description: "iPhone 12 de 128 GB", Category: "Celulares"})
	internal.SliceProducts = append(internal.SliceProducts, internal.Product{ID: 1, Name: "iPhone 12", Price: 999.99, Description: "iPhone 12 de 128 GB", Category: "Celulares"})
	internal.SliceProducts = append(internal.SliceProducts, internal.Product{ID: 2, Name: "iPhone 12 Pro", Price: 1299.99, Description: "iPhone 12 Pro de 128 GB", Category: "Celulares"})
	internal.SliceProducts = append(internal.SliceProducts, internal.Product{ID: 3, Name: "iPhone 12 Pro Max", Price: 8923.2399, Description: "Samsungs 12 Pro Max de 128 GB", Category: "Celulares"})
	internal.SliceProducts = append(internal.SliceProducts, internal.Product{ID: 4, Name: "iPhone 12 Mini", Price: 999.99, Description: "iPhone 12 Mini de 128 GB", Category: "Celulares"})

	// Append new product to sliceProducts using method Save

	internal.SliceProducts[0].Save()

	// Print all products using method GetAll
	fmt.Println(" GetAll SliceProducts")
	internal.GetAll()

	// Print a product using function GetByID
	fmt.Println(" GetByID SliceProducts")
	fmt.Println(internal.GetByID(1))

	fmt.Println("Fin del programa")
}
