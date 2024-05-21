package patterns

import "fmt"

/*
	Паттерн "Цепочка вызовов" позволяет передавать запросы по цепочке обработчиков.
	Каждый обработчик решает, обработать запрос или передать его следующему обработчику в цепочке.

	Плюсы:
	Уменьшение связанности
	Гибкость
	Расширяемость

	Минусы:
	Неопределенность
	Трудности в отладке
*/

// ? Интерфейс Обработчика

type Handler interface {
	setNext(handler Handler)
	handle(request int)
}

// ? Базовый Обработчик

type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) setNext(handler Handler) {
	h.next = handler
}

func (h *BaseHandler) handleNext(request int) {
	if h.next != nil {
		h.next.handle(request)
	}
}

// ? Младшая поддержка

type JuniorSupport struct {
	BaseHandler
}

func (h *JuniorSupport) handle(request int) {
	if request <= 1 {
		fmt.Println("Junior Support handles request", request)
	} else {
		fmt.Println("Junior Support passes request", request)
		h.handleNext(request)
	}
}

// ? Старшая поддержка

type SeniorSupport struct {
	BaseHandler
}

func (h *SeniorSupport) handle(request int) {
	if request <= 2 {
		fmt.Println("Senior Support handles request", request)
	} else {
		fmt.Println("Senior Support passes request", request)
		h.handleNext(request)
	}
}

// ? Поддержка менеджера

type ManagerSupport struct {
	BaseHandler
}

func (h *ManagerSupport) handle(request int) {
	if request <= 3 {
		fmt.Println("Manager handles request", request)
	} else {
		fmt.Println("Request", request, "cannot be handled")
	}
}

func CheckChain() {
	junior := &JuniorSupport{}
	senior := &SeniorSupport{}
	manager := &ManagerSupport{}

	junior.setNext(senior)
	senior.setNext(manager)

	requests := []int{1, 2, 3, 4}

	for _, req := range requests {
		fmt.Println("Handling request", req)
		junior.handle(req)
		fmt.Println()
	}
}
