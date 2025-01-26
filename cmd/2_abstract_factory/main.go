package main

// https://refactoring.guru/ru/design-patterns/abstract-factory/go/example

import (
	"fmt"
	"patterns/cmd/2_abstract_factory/stuff"
)

// Клиентский код

func main() {
	// абстрактная фабрика возвращает конкретные фабрики
	adidasFactory, _ := stuff.GetSportsFactory("adidas")
	nikeFactory, _ := stuff.GetSportsFactory("nike")

	// конкретные фабрики возвращают конкретные продукты
	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s stuff.IShoe) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}

func printShirtDetails(s stuff.IShirt) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}
