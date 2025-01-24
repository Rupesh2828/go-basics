package main

import (
	"fmt"
)

// func main() {
// 	fmt.Println("Hello GO..")
// 	agency:= "Fast Track"
// 	fmt.Println("Welcome to ",agency)

// 	var totalCars int = 50;
// 	fmt.Println("Our fleet consists of ", totalCars, "cars")

// 	startingPrice:= 29

// 	fmt.Println("Our prices are at %V", startingPrice, "Take your price")

// 	//boolean

// 	// isIncluded:= true

// }

// func main(){
// 	var name string
// 	fmt.Println("Name:", name)
// 	var count int
// 	fmt.Println("Count:", count)
// 	var price float64
// 	fmt.Println("Price: ", price)
// 	var insured bool
// 	fmt.Println("Insured:", insured)
	

// }

// func main(){

// 	str:= "Hello GO"
// 	length:=len(str)
// 	fmt.Println(length)

// 	//below comparison is true, go is case insesitive.
// 	str1:="Go Programming"
// 	str2:= "go programming"

// 	strings.EqualFold(str1,str2)
// 	fmt.Println(strings.EqualFold(str1,str2))

//slicing

// str3:= "Hello, World!"
// wIndex:= strings.Index(str3,"W")

// substr:= str3[wIndex:12]  //fetch from W as we have above, from that to till 12, counts till 11
// fmt.Println(substr)
// fmt.Println(str3[wIndex:])
// fmt.Println(str3[:wIndex])

// fmt.Println(strings.Replace(str3, "World", "GO", 1))

// 	str4:= "Go Propgramming"

// 	fmt.Println(strings.ToLower(str4))
// 	fmt.Println(strings.Contains(str4,"Go"))

// }

// func main(){
// 	const Agency = "Helo"
// 	//for multiple constants declaration
// 	const (
// 		Founded = 2008
// 		Founder="Rupesh"

// 	)
// 	fmt.Println(Agency)
// 	fmt.Println(Founder)

// 	const (
// 		Economy = iota
// 		Compact
// 		Standard
// 	)
// 	fmt.Println("Economy", Economy)
// 	fmt.Println("Compact", Compact)
// 	fmt.Println("Standard", Standard)

// }

// for user input

// func main(){
// var name string
// fmt.Println("What is your name?")
// fmt.Scanln(&name)

// fmt.Println("Hello,", name)
// }

//POINTERS.

// func sayHello(s string){
// 	s = "Hello"

// }
// func sayHelloPointer(s *string){
// 	*s = "Hello"

// }

// func main(){
// 	var greet = "Hello, Go"
// 	arrayofindex:=5
// 	fmt.Println(arrayofindex)
// 	sayHello(greet)
// 	fmt.Println("Afer sayHello: ", greet)

// 	sayHelloPointer(&greet)
// 	fmt.Println("Afer sayHelloPointer: ", greet)

// }

//conditional operator.

// func main() {
// 	 isFast:= true
// 	 isSlow:= true

// 	if isFast {
// 		fmt.Println("You are going too fast, slow down.")
// 	} else if isSlow{
// 		fmt.Println("Thank u for not speeding")

// 	}
// }


// func main() {
//     day := "Tuesday"

//     // Switch statement
//     switch day {
//     case "Monday":
//         fmt.Println("Start of the workweek.")
// 		fallthrough
//     case "Tuesday":
//         fmt.Println("Second day of the week.")
// 		fallthrough
//     case "Friday":
//         fmt.Println("Almost the weekend!")
//     case "Saturday", "Sunday": // Multiple cases
//         fmt.Println("It's the weekend!")
//     default: // Default case
//         fmt.Println("Just another weekday.")
//     }
// }

// func main(){
// 	for i:=0; i<5; i++{
// 		fmt.Println(i)
// 	}
// }


//ARRAYS IN GO.

func main(){

	var bodyTypes [3]string

	// arr := [5]string{"1", "2", "3", "4"}

	thisIsArray := []bool{true, false}

	fmt.Println(thisIsArray) 


	bodyTypes[0] = "Rupesh Jadhav"

	fmt.Println((bodyTypes))

}

