package main

import "fmt"

func dynamic_array() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[1])

	updatedPrices := append(prices, 5.99)
	fmt.Println(updatedPrices)

}
