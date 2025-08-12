package main

import "fmt"
func main() {
	prices := []float64{10,20,30}
	taxRates := []float64{0,0.07, 0.1, 0.15}
	result := make(map[float64][]float64)
	for _, taxRate := range taxRates{
		taxIncludedPrices := make([]float64, len(prices))
		for priceIndex, price := range prices {
			taxIncludedPrices[priceIndex] = price *(1 + taxRate) 
		}
		result[taxRate] = taxIncludedPrices;
	}
	fmt.Println("Prices with tax included:", result);

	for taxRate, taxIncludedPrice := range result{
		fmt.Print("\ntaxRate, Tax Included Prices: ", taxRate, taxIncludedPrice)
		for _, prices := range taxIncludedPrice{
			fmt.Println("\nPrice with tax included:", prices)
		}

	}
}
