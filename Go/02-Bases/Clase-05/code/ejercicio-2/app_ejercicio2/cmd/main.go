package main

import (
	"app_ejercicio_2/app_ejercicio2/internal"
	"fmt"
)

func main() {
	fmt.Println("Inicio del programa")
	// Obtener el precio de un producto pequeño
	smallProduct, err := internal.Factory("small", 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	priceSmallProduct, err := smallProduct.GetPrice()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("El precio del producto pequeño es: ", priceSmallProduct)
	// Obtener el precio de un producto mediano
	product, err := internal.Factory("medium", 1000)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Hacer un type assertion para convertir Product a MediumProduct
	mediumProduct, ok := product.(*internal.MediumProduct)
	if !ok {
		fmt.Println("El producto no es de tipo MediumProduct")
		return
	}

	// Set the IsStore field to true
	mediumProduct.IsStore = true

	// Get the price of the medium product
	priceMediumProduct, err := mediumProduct.GetPrice()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("El precio del producto mediano es: ", priceMediumProduct)

	// Obtener el precio de un producto grande
	product2, err := internal.Factory("large", 10000)

	if err != nil {
		fmt.Println(err)
		return
	}
	// Hacer una afirmación de tipo para convertir Product a LargeProduct
	largeProduct, ok := product2.(*internal.LargeProduct)
	if !ok {
		fmt.Println("El producto no es de tipo LargeProduct")
		return
	}
	// Ahora puedes acceder a IsStore
	largeProduct.IsStore = true
	// Set transportCost to 2500
	largeProduct.TransportCost = 2500

	priceLargeProduct, err := largeProduct.GetPrice()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("El precio del producto grande es: ", priceLargeProduct)
}
