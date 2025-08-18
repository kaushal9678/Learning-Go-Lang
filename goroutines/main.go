package main

import (
	"fmt"
	"time"
)


func greet(name string) {
	fmt.Println( "Hello, " + name + "!")
}
func greetUserWithWelcome(name string) {
	time.Sleep(2*time.Second) // Simulate a delay
	//return "Welcome, " + name + "!"
	fmt.Println("Welcome, " + name + "!")
}
func sumOfTwoNumbers(a, b int) {
	fmt.Println(a+b);
}

func main(){
	greet("kaushal yadav")
	greetUserWithWelcome("kaushal yadav")
	sumOfTwoNumbers(10,20)

}
