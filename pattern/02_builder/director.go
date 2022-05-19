package main

//Содержит в себе интерфейс с методами, которые направляет директор строителю
type director struct {
	builder iBuilder
}

//Создается объект директора, который содержит конкретный тип объекта, релизующий интерфейс(нормальный дом, иглу)
func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildHouse() house {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}
