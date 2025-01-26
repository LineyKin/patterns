package stuff

// конкретный строитель нормального дома

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) SetWindowType() {
	b.windowType = "Wooden window"
}

func (b *NormalBuilder) SetDoorType() {
	b.doorType = "Wooden door"
}

func (b *NormalBuilder) SetNumFloor() {
	b.floor = 2
}

func (b *NormalBuilder) GetHouse() House {
	return House{
		DoorType:   b.doorType,
		WindowType: b.windowType,
		Floor:      b.floor,
	}
}
