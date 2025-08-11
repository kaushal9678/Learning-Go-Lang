package main

import "fmt"

func main() {
	arrayOfPrices := [4]float64{10.99, 10.43, 5.89, 25.49}
	fmt.Println(arrayOfPrices)
	//arrayOfSlices := arrayOfPrices[1:]
	// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
	arrayOfHobbies := [3]string{"Reading", "Cycling", "Cooking"}
	fmt.Println(arrayOfHobbies)
	// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
	fmt.Println(arrayOfHobbies[0]) // First element
	fmt.Println("1:3",arrayOfHobbies[1:3]) // Second and third element combined as a new list	

// Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
	sliceOfHobbies1 := arrayOfHobbies[0:2] // First and second elements
	sliceOfHobbies2 := arrayOfHobbies[:2]  // First and second elements
	fmt.Println("sliceOfHobbies1", sliceOfHobbies1)
	fmt.Println("sliceOfHobbies2", sliceOfHobbies2)
	// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
	sliceOfHobbies1 = arrayOfHobbies[1:3] // Second and last elements
	fmt.Println("sliceOfHobbies1 [1:3]", sliceOfHobbies1)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	dynamicGoals := []string{"Learn Go", "Build a web app"}
	fmt.Println(dynamicGoals)
	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	dynamicGoals[1] = "Master Go"
	dynamicGoals = append(dynamicGoals, "Contribute to open source")
	fmt.Println(dynamicGoals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
	type Product struct {
		Title string
		ID    int
		Price float64
		
	}

	dynamicProducts := []Product{
		{Title: "Laptop", ID: 1, Price: 999.99},
		{Title: "Smartphone", ID: 2, Price: 499.99},
	}
	fmt.Println(dynamicProducts)
	dynamicProducts = append(dynamicProducts, Product{Title: "Tablet", ID: 3, Price: 299.99})
	fmt.Println(dynamicProducts)
}
// arrays/list.go	

// Time to practice what you learned!



// 3)



