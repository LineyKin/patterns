package stuff

// Директор. Класс, задающий порядок строительства объекта

type Director struct {
	builder IBuilder
}

func NewDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) BuildHouse() House {
	d.builder.SetDoorType()
	d.builder.SetNumFloor()
	d.builder.SetWindowType()

	return d.builder.GetHouse()
}
