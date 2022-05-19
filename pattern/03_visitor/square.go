package main

type square struct {
	side int
}

//Добавляем в каждую фигуру метод ассепт, в который добавляем интерфейс визитор,
// который в дальнейшем будет содержать объект с конкретной реализацией
func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}

func (s *square) getType() string {
	return "Square"
}
