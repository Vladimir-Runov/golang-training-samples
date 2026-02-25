package main

import (
	"fmt"
)

// Определяем интерфейс Shape
type Shape interface {
	Info() string
}

// Определяем базовый класс Figure
type Figure struct {
	x, y float64 // Координаты центра
}

// Определяем класс Circle, который содержит Figure
type Circle struct {
	Figure
	radius float64 // Радиус круга
}

// Метод для отображения информации о круге
func (c Circle) Info() string {
	return fmt.Sprintf("Circle - Center: (%f, %f), Radius: %f,S=%f", c.x, c.y, c.radius, 3.14*c.radius*c.radius)
}

// Определяем класс Triangle, который содержит Figure
type Triangle struct {
	Figure
	side float64 // Длина стороны треугольника
}

// Метод для отображения информации о треугольнике
func (t Triangle) Info() string {
	return fmt.Sprintf("Triangle - Center: (%f, %f), Side Length: %f, S=%f", t.x, t.y, t.side, 732050*4*t.side*t.side)
}

// Определяем класс Square, который содержит Figure
type Square struct {
	Figure
	side float64 // Длина стороны квадрата
}

// Метод для отображения информации о квадрате
func (s Square) Info() string {
	return fmt.Sprintf("Square - Center: (%f, %f), Side Length: %f,S=%f", s.x, s.y, s.side, s.side*s.side)
}

func main() {
	// Создаем массив объектов типа Shape
	figures := []Shape{
		Circle{Figure{0, 0}, 5},
		Triangle{Figure{1, 1}, 3},
		Square{Figure{2, 2}, 4},
	}

	// Выводим информацию о каждой фигуре
	for _, figure := range figures {
		fmt.Println(figure.Info())
	}
}
