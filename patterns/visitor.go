package patterns

import (
	"fmt"
	"math"
)

/*
	Паттерн "Посетитель" позволяет добавлять новые операции к объектам без изменения этих объектов.
	Он позволяет определить новую операцию без изменения классов элементов, над которыми она выполняется.

	Плюсы:
	Упрощение добавления нового функционала
	Разделение логики
	Поддержка открытости\закрытости

	Минусы:
	Разрыв связей
	Тесная связанность
	Сложность при добавлении новых элементов
*/

// ? Интерфейс Посетителя

type Visitor interface {
	visitCircle(*Circle)
	visitRectangle(*Rectangle)
}

// ? Интерфейс Фигуры

type Shape interface {
	accept(Visitor)
}

// ? Круг

type Circle struct {
	Radius float64
}

func (c *Circle) accept(v Visitor) {
	v.visitCircle(c)
}

// ? Прямоугольник

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) accept(v Visitor) {
	v.visitRectangle(r)
}

// ? Посетитель для вычисления площади

type CalculateArea struct {
	area float64
}

func (a *CalculateArea) visitCircle(c *Circle) {
	a.area = math.Pi * math.Pow(c.Radius, 2)
	fmt.Printf("Area of Circle: %.2f\n", a.area)
}

func (a *CalculateArea) visitRectangle(r *Rectangle) {
	a.area = r.Width * r.Height
	fmt.Printf("Area of Rectangle: %.2f\n", a.area)
}

// ? Посетитель для вычисления периметра

type PerimeterCalculator struct {
	perimeter float64
}

func (p *PerimeterCalculator) visitCircle(c *Circle) {
	p.perimeter = 2 * math.Pi * c.Radius
	fmt.Printf("Perimeter of Circle: %.2f\n", p.perimeter)
}

func (p *PerimeterCalculator) visitRectangle(r *Rectangle) {
	p.perimeter = 2 * (r.Width + r.Height)
	fmt.Printf("Perimeter of Rectangle: %.2f\n", p.perimeter)
}

func CheckVisitor() {
	circle := &Circle{Radius: 5}
	rectangle := &Rectangle{Width: 4, Height: 6}

	areaCalculator := &CalculateArea{}
	perimeterCalculator := &PerimeterCalculator{}

	// ? Расчет площади
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	// ? Расчет периметра
	circle.accept(perimeterCalculator)
	rectangle.accept(perimeterCalculator)
}
