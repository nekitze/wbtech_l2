package pattern

import "log"

/*
Фабричный метод - порождающий паттерн проектирования, который определяет общий
интерфейс для создания объектов в суперклассе, позволяя подклассам изменять тип
создаваемых объектов.

Применяется, когда:
- классу заранее неизвестно, объекты каких классов ему нужно создавать;
- класс спроектирован так, чтобы объекты, которые он создает, специфицировались
подклассами;
- класс делегирует свои обязанности одному из нескольких вспомогательных
подкалссов, и нужно локализовать знание о том, какой класс принимает эти
обязанности на себя;

Плюсы:
- избавляет главный класс от привязки к конкретным типам объектов;
- выделяет код производства объектов в одно место, упрощая поддержку кода;
- упрощает добавление новых типов объектов в программу;
Минусы:
- может привести к созданию больших параллельных иерархий классов, так как
для каждого типа объекта надо создать свой подкласс создателя;

Пример реального применения: добавление новых типов товаров в интернет-магазин,
когда каждая категория товаром может иметь свои особенности в процессе создания
соответствующих объектов;
*/

type action string

const (
	A action = "A"
	B action = "B"
	C action = "C"
)

type Product interface {
	Use() string
}

type Creator interface {
	CreateProduct(action action) Product
}

type ConcreteCreator struct{}

func NewCreator() Creator {
	return &ConcreteCreator{}
}

func (p *ConcreteCreator) CreateProduct(action action) Product {
	var product Product

	switch action {
	case A:
		product = &ConcreteProductA{string(action)}
	case B:
		product = &ConcreteProductB{string(action)}
	case C:
		product = &ConcreteProductC{string(action)}
	default:
		log.Fatalln("Unknown Action")
	}

	return product
}

type ConcreteProductA struct {
	action string
}

func (p *ConcreteProductA) Use() string {
	return p.action
}

type ConcreteProductB struct {
	action string
}

func (p *ConcreteProductB) Use() string {
	return p.action
}

type ConcreteProductC struct {
	action string
}

func (p *ConcreteProductC) Use() string {
	return p.action
}
