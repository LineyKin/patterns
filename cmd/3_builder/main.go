package main

import (
	"fmt"
	"patterns/cmd/3_builder/stuff"
)

func main() {
	normalBuilder := stuff.GetBuilder("normal")
	iglooBuilder := stuff.GetBuilder("igloo")

	director := stuff.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	fmt.Printf("Normal house door type: %s\n", normalHouse.DoorType)
	fmt.Printf("Normal house window type: %s\n", normalHouse.WindowType)
	fmt.Printf("Normal house num floor: %d\n", normalHouse.Floor)

	director = stuff.NewDirector(iglooBuilder)
	iglooHouse := director.BuildHouse()

	fmt.Printf("Igloo house door type: %s\n", iglooHouse.DoorType)
	fmt.Printf("Igloo house window type: %s\n", iglooHouse.WindowType)
	fmt.Printf("Igloo house num floor: %d\n", iglooHouse.Floor)
}
