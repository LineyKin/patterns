package stuff

// конкретный строитель иглу (дом эскимосов)

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) SetWindowType() {
	b.windowType = "Snow window"
}

func (b *IglooBuilder) SetDoorType() {
	b.doorType = "Snow door"
}

func (b *IglooBuilder) SetNumFloor() {
	b.floor = 1
}

func (b *IglooBuilder) GetHouse() House {
	return House{
		DoorType:   b.doorType,
		WindowType: b.windowType,
		Floor:      b.floor,
	}
}
