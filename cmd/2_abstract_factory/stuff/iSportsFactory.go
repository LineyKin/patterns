package stuff

import "fmt"

// Интерфейс абстрактной фабрики
type ISportsFactory interface {
	MakeShoe() IShoe   // интерфейс обувь
	MakeShirt() IShirt // интерфейс футболки
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	switch brand {
	case "adidas":
		return &Adidas{}, nil
	case "nike":
		return &Nike{}, nil
	default:
		return nil, fmt.Errorf("wrong brand type passed")
	}
}
