package main

import (
	"example.com/pricing-calculator/cmdmanager"
	"example.com/pricing-calculator/prices"
)

func main() {
	taxRates := []float64{0,0.07, 0.1, 0.15}
	
	for _, taxRate := range taxRates {
		//fm := filemanager.New("prices.txt", fmt.Sprintf("tax_included_prices_%.0f.json", taxRate * 100),);
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm,taxRate)
		priceJob.Process()
	}

}
