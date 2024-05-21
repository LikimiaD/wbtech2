package patterns

import "fmt"

/*
	Паттерн "Команда" превращает запросы в объекты,
	позволяя параметризовать методы с различными запросами,
	ставить запросы в очередь или протоколировать их,
	а также поддерживать отмену операций.

	Плюсы:
	Инкапсуляция запросов
	Поддержка отмены операций
	Логирование и очереди
	Упрощенное тестирование

	Минусы:
	Сложный паттерн
	Перегрузка
*/

// ? Интерфейс команды

type Command interface {
	execute()
	undo()
}

// ? Получатель

type Light struct {
	Active bool
}

func (l *Light) on() {
	l.Active = true
	fmt.Println("Light is ON")
}

func (l *Light) off() {
	l.Active = false
	fmt.Println("Light is OFF")
}

// ? Команда для включения света

type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) execute() {
	c.light.on()
}

func (c *LightOnCommand) undo() {
	c.light.off()
}

// ? Команда для выключения света

type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) execute() {
	c.light.off()
}

func (c *LightOffCommand) undo() {
	c.light.on()
}

// ? Инициатор

type RemoteControl struct {
	history []Command
}

func (r *RemoteControl) setCommand(command Command) {
	r.history = append(r.history, command)
	command.execute()
}

func (r *RemoteControl) undo() {
	if len(r.history) > 0 {
		lastCommand := r.history[len(r.history)-1]
		lastCommand.undo()
		r.history = r.history[:len(r.history)-1]
	} else {
		fmt.Println("No commands to undo")
	}
}

func CheckCommand() {
	// ? Получатель
	light := &Light{}

	// ? Команды
	lightOn := &LightOnCommand{light}
	lightOff := &LightOffCommand{light}

	// ? Инициатор
	remote := &RemoteControl{}

	// ? Клиент
	remote.setCommand(lightOn)
	remote.setCommand(lightOff)

	// ? Отмена последней команды
	remote.undo()
	remote.undo()
}
