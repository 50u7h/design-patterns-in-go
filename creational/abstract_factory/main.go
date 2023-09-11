package main

import "fmt"

func main() {
	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike")
	nikeShoe := nikeFactory.makeShoe()
	nikeShort := nikeFactory.makeShort()
	adidasShoe := adidasFactory.makeShoe()
	adidasShort := adidasFactory.makeShort()
	printShoeDetails(nikeShoe)
	printShortDetails(nikeShort)
	printShoeDetails(adidasShoe)
	printShortDetails(adidasShort)
}

func printShoeDetails(s iShoe) {
	fmt.Printf("Shoe Brand: %s", s.getBrand())
	fmt.Println()
	fmt.Printf("Shoe Size: %d", s.getSize())
	fmt.Println()
	fmt.Println()
}

func printShortDetails(s iShort) {
	fmt.Printf("Short Brand: %s", s.getBrand())
	fmt.Println()
	fmt.Printf("Short Size: %s", s.getSize())
	fmt.Println()
	fmt.Println()
}
