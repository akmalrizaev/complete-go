package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A Book"}
	prices := [4]float64{10.99, 9.99, 5.99, 20.00}
	fmt.Println(prices)

	productNames[2] = "A Carpet"

	fmt.Println(productNames)
	fmt.Println(prices[2])
}
