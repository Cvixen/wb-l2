package main

type circle struct {
	radius int
}

//Добавляем в каждую фигуру метод ассепт, в который добавляем интерфейс визитор,
// который в дальнейшем будет содержать объект с конкретной реализацией
func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

func (c *circle) getType() string {
	return "Circle"
}
