package patterns

import "fmt"

/*
	Паттерн "Строитель" используется для создания сложных объектов пошагово.
	Он позволяет создавать различные представления объекта, используя один и тот же процесс конструирования.

	Плюсы:
	Разделение кода
	Контроль над процессом построения
	Поддержка различных представлений

	Минусы:
	Избыточность
	Сложность кода
*/

// ? Продукт

type Car struct {
	Chassis  string
	Body     string
	Paint    string
	Interior string
}

// ? Интерфейс Строителя

type CarBuilder interface {
	fixChassis()
	fixBody()
	paint()
	fixInterior()
	getCar() Car
}

// ? Строитель для классического автомобиля

type ClassicCarBuilder struct {
	car Car
}

func GetClassicCarBuilder() *ClassicCarBuilder {
	return &ClassicCarBuilder{Car{}}
}

func (c *ClassicCarBuilder) fixChassis() {
	fmt.Println("Assembling chassis of the classical model")
	c.car.Chassis = "Classic Chassis"
}

func (c *ClassicCarBuilder) fixBody() {
	fmt.Println("Assembling body of the classical model")
	c.car.Body = "Classic Body"
}

func (c *ClassicCarBuilder) paint() {
	fmt.Println("Painting body of the classical model")
	c.car.Paint = "Classic White Paint"
}

func (c *ClassicCarBuilder) fixInterior() {
	fmt.Println("Setting up interior of the classical model")
	c.car.Interior = "Classic interior"
}

func (c *ClassicCarBuilder) getCar() Car {
	return c.car
}

// ? Строитель для современного автомобиля

type ModernCarBuilder struct {
	car Car
}

func GetModernCarBuilder() *ModernCarBuilder {
	return &ModernCarBuilder{Car{}}
}

func (c *ModernCarBuilder) fixChassis() {
	fmt.Println("Assembling chassis of the modern model")
	c.car.Chassis = "Modern Chassis"
}

func (c *ModernCarBuilder) fixBody() {
	fmt.Println("Assembling body of the modern model")
	c.car.Body = "Modern Body"
}

func (c *ModernCarBuilder) paint() {
	fmt.Println("Painting body of the modern model")
	c.car.Paint = "Modern White Paint"
}

func (c *ModernCarBuilder) fixInterior() {
	fmt.Println("Setting up interior of the modern model")
	c.car.Interior = "Modern interior"
}

func (c *ModernCarBuilder) getCar() Car {
	return c.car
}

// ? Директор

type Director struct {
	builder CarBuilder
}

func NewDirector(b CarBuilder) *Director {
	return &Director{b}
}

func (d *Director) constructCar() {
	d.builder.fixChassis()
	d.builder.fixBody()
	d.builder.paint()
	d.builder.fixInterior()
}

func CheckBuilder() {
	classicBuilder := GetClassicCarBuilder()
	director := NewDirector(classicBuilder)
	director.constructCar()
	classicCar := classicBuilder.getCar()
	fmt.Printf("Car built: %+v\n", classicCar)

	modernBuilder := GetModernCarBuilder()
	director = NewDirector(modernBuilder)
	director.constructCar()
	modernCar := modernBuilder.getCar()
	fmt.Printf("Car built: %+v\n", modernCar)
}
