package main

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

//Возвращяет определенный объект, реализующий интерфес билдера
func getBuilder(builderType string) iBuilder {
	if builderType == "normal" {
		return &normalBuilder{}
	}
	if builderType == "igloo" {
		return &iglooBuilder{}
	}
	return nil
}
