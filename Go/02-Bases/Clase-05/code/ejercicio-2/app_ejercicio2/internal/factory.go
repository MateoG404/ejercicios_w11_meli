// Script created to return a new instance of the struct Product
package internal

import "errors"

const (
	SmallProductConst  = "small"
	MediumProductConst = "medium"
	LargeProductConst  = "large"
)

func Factory(typeProduct string, price float64) (Product, error) {
	// Switch to return the price of the type of product
	switch typeProduct {
	case SmallProductConst:
		return &SmallProduct{price: price}, nil
	case MediumProductConst:
		return &MediumProduct{price: price}, nil
	case LargeProductConst:
		return &LargeProduct{price: price}, nil
	default:
		return nil, errors.New("error: unknown product type")
	}
}
