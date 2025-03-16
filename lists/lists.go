package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A Book"}
	prices := []float64{10.99, 9.99, 5.99, 20.00}
	fmt.Println(prices)

	productNames[2] = "A Carpet"

	fmt.Println(productNames)
	fmt.Println(prices[2])

	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}
