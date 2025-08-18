package main

import (
	"fmt"
	"time"
)


func greet(name string) {
	fmt.Println( "Hello, " + name + "!")
}
func greetUserWithWelcome(name string, doneChan chan bool) {
	time.Sleep(2*time.Second) // Simulate a delay
	//return "Welcome, " + name + "!"
	fmt.Println("Welcome, " + name + "!")
	doneChan <- true
}
func sumOfTwoNumbers(a, b int) {
	fmt.Println(a+b);
}

func main(){
	greet("kaushal yadav")
	done := make(chan bool)
	go greetUserWithWelcome("kaushal yadav", done)
	<-done // Wait for the goroutine to finish
	fmt.Println("Greeted user in a goroutine")
	// Call the sumOfTwoNumbers function
	sumOfTwoNumbers(10,20)

}
