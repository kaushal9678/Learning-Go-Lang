package main

import (
	"fmt"
	"time"
)


func greet(name string, done chan bool) {
	fmt.Println( "Hello, " + name + "!")
	done <- true
}
func greetUserWithWelcome(name string, doneChan chan bool) {
	time.Sleep(2*time.Second) // Simulate a delay
	//return "Welcome, " + name + "!"
	fmt.Println("Welcome, " + name + "!")
	doneChan <- true
	close(doneChan)// Close the channel to signal completion
}
func sumOfTwoNumbers(a, b int, done chan bool) {
	fmt.Println(a+b);
	done <- true
}

func main(){
	dones := make([]chan bool, 3);
	dones[0] = make(chan bool)
	go greet("kaushal yadav", dones[0])
	dones[1] = make(chan bool)
	go greetUserWithWelcome("kaushal yadav", dones[1])
	// Wait for the goroutine to finish
	fmt.Println("Greeted user in a goroutine")
	// Call the sumOfTwoNumbers function
	dones[2] = make(chan bool)
	go sumOfTwoNumbers(10,20, dones[2])

	for _, done := range dones{
		<-done
		fmt.Println("Done with the goroutine")
		if done == dones[0] {
			fmt.Println("Greeted user in main function")
		} else if done == dones[1] {
			fmt.Println("Greeted user in a goroutine")
		} else if done == dones[2] {
			fmt.Println("Sum of two numbers is calculated in a goroutine")
		}
		if done == dones[len(dones)-1] {
			fmt.Println("All goroutines are done")
			break
		}
		fmt.Println("Waiting for the next goroutine to finish")
		time.Sleep(1 * time.Second) // Simulate some processing time
		
	}

}
