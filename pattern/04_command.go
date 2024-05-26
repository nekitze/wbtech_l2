package pattern

import "fmt"

/*
Команда - поведенческий паттерн, который инкапуслирует запрос как
объект, позволяя задавать параметры клиентов для обработки соответствующих
запросов, ставить запросы в очередь или протоколировать их, а также
поддерживать отмену операций.

Применяется, когда хотим:
- параметризировать объекты выполняемым действием;
- определять, ставить в очередь и выполнять запросы в рантайме;
- поддерживать отмену операций;
- поддерживать протоколирование изменений;
- структурировать систему на основе высокоуровневых операций, построенных из
примитивных;

Плюсы:
- убирает прямую зависимость между объектами, вызывающими операции, и объектами,
которые их непосредственно выполняют;
- позволяет реализовать простую отмену и повтор операций;
- позволяет реализовать отложенный запуск операций;
- позволяет собирать сложные команды из простых;
Минусы:
- усложняет код программы введением множества дополнительных классов;

Пример реального применения: обработка объектов фоторедактора последовательностью команд;
*/

type Command interface {
	Execute()
}

type Receiver struct {
	Name string
}

func (r *Receiver) Action() {
	fmt.Println("Action has been taken by", r.Name)
}

type ConcreteCommand struct {
	receiver *Receiver
}

func (c *ConcreteCommand) Execute() {
	c.receiver.Action()
}

type Invoker struct {
	commands []Command
}

func (i *Invoker) AddCommand(command Command) {
	i.commands = append(i.commands, command)
}

func (i *Invoker) UndoCommand() {
	if len(i.commands) != 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

func (i *Invoker) Execute() {
	for _, command := range i.commands {
		command.Execute()
	}
}
