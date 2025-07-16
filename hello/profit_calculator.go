package main

import (
	"fmt"
)

func main() {
	var revenue, expense, taxRate float64
	fmt.Print("Enter the total revenue: ")
	fmt.Scanf("%f", &revenue)
	fmt.Print("Enter the expense : ")
	fmt.Scanf("%f", &expense)	
	fmt.Print("Enter the taxRate : ")
	fmt.Scanf("%f", &taxRate)	
	ebt := calculateEarningBeforeTax(revenue, expense)
	fmt.Printf("The Earning before tax is: $%.2f\n", ebt)
	eat := calculateEarningAfterTax(ebt, taxRate)
	fmt.Printf("The Earning after tax is: $%.2f\n", eat)
	fmt.Println("Thank you for using the profit calculator!")
	fmt.Println("Have a great day!")
	fmt.Println("Goodbye!")
	}

func calculateEarningBeforeTax(revenue, expenses float64) float64 {
	return revenue - expenses
}
func calculateEarningAfterTax(ebt, taxRate float64) float64 {
	return ebt * (1 - taxRate)
}