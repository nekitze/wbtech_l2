package pattern

import "fmt"

/*
Цепочка обязанностей - поведенческий паттерн, который позволяет передавать
запросы последовательно по цепочке обработчиков. Каждый последующий
обработчик решает, может ли он обработать запрос сам и стоит ли передавать
запрос дальше по цепи.

Применяется, когда:
- есть более одного объекта, способного обработать запрос, причем настоящий
обработчик заранее неизвестен и должен быть найден автоматически;
- нужно явно отправить запрос одному из нескольких объектов, не указывая явно,
какому именно;
- набор объектов, способных обработать запрос должен задаваться динамически;

Плюсы:
- уменьшает зависимость между клиентом и обработчиками;
Минусы:
- запрос может остаться никем не обработанным;

Пример реального применения: Filter Chain в сервисе авторизации;
*/

type Handler interface {
	SetNext(h Handler)
	Handle(request string) string
}

type AbstractHandler struct {
	next Handler
}

func (h *AbstractHandler) SetNext(next Handler) {
	h.next = next
}

func (h *AbstractHandler) Handle(request string) string {
	if h.next != nil {
		return h.next.Handle(request)
	}
	return ""
}

type HandlerA struct {
	AbstractHandler
}

func (h *HandlerA) Handle(request string) string {
	fmt.Println("Обработка запроса A:", request)
	return fmt.Sprintf("Ответ от A: %s", request)
}

type HandlerB struct {
	AbstractHandler
}

func (h *HandlerB) Handle(request string) string {
	fmt.Println("Обработка запроса B:", request)
	return fmt.Sprintf("Ответ от B: %s", request)
}
