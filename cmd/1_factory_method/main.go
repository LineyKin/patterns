package main

import (
	"fmt"
	"patterns/cmd/1_factory_method/stuff"
)

// Фабричный метод — это порождающий паттерн проектирования,
// который определяет общий интерфейс для создания объектов в суперклассе,
// позволяя подклассам изменять тип создаваемых объектов.

// https://refactoring.guru/ru/design-patterns/factory-method/go/example

func main() {
	ak47, _ := stuff.GetGun("ak47")
	musket, _ := stuff.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g stuff.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
