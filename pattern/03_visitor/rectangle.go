package main

type rectangle struct {
	l int
	b int
}

//Добавляем в каждую фигуру метод ассепт, в который добавляем интерфейс визитор,
// который в дальнейшем будет содержать объект с конкретной реализацией
func (t *rectangle) accept(v visitor) {
	v.visitForrectangle(t)
}

func (t *rectangle) getType() string {
	return "rectangle"
}
