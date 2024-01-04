// Script created to define the struct Product and its methods
package internal

import "errors"

type Product interface {
	GetPrice() (float64, error)
}

// __________________________________________________SMALL_PRODUCTS____________________________________________________
// Struct Small Product
type SmallProduct struct {
	price float64
}

// Method GetPrice of SmallProduct
func (s *SmallProduct) GetPrice() (float64, error) {
	return s.price, nil
}

// __________________________________________________MEDIUM_PRODUCTS____________________________________________________
// Struct Medium Product

type MediumProduct struct {
	price   float64
	IsStore bool // if isStore is true, the price is 3% more than the price
}

// Method GetPrice of MediumProduct
func (m *MediumProduct) GetPrice() (float64, error) {
	if m.IsStore { // if isStore is true, the price is 3% more than the price
		return (m.price * 0.03) + m.price, nil
	}
	return 0, errors.New("Error: unexpected error")
}

// Method se

// __________________________________________________LARGE_PRODUCTS____________________________________________________s
// Struct Large Product

type LargeProduct struct {
	price         float64 // if isStore is true, the price is 6% more than the price
	TransportCost float64
	IsStore       bool // Add the missing isStore field
}

// Method GetPrice of LargeProduct
func (l *LargeProduct) GetPrice() (float64, error) {
	if l.IsStore { // if isStore is true, the price is 6% more than the price
		return (l.price * 0.06) + l.price + l.TransportCost, nil
	}
	return 0, errors.New("Error: unexpected error") // Fix the return statement
}
