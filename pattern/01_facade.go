package pattern

import "fmt"

/*
Фасад (структурный паттерн) - предоставляет унифицированный интерфейс вместо набора интерфейсов
некоторой подсистемы. Фасад определяет интерфейс более высокого
уровня, который упрощает использование подсистемы.

Применяется, когда:
- нужно представить простой интерфейс к сложной подсистеме;
- нужно уменьшить количество зависимостей между клиентом и сложной системой;
- нужно разложить подсистему на отдельные слои;

Плюсы:
- изоляция клиентов от компонентов сложной подсистемы.
Минусы:
- риск появления божественного объекта, привязанного ко всем классам программы;
- ограничивает возможности клиента;
*/

// Compiler - простой интерфейс к сложной подсистеме компилятора.
type Compiler struct {
	Scanner
	Parser
	Validator
	CodeGenerator
}

func (s *Compiler) Compile() {
	result := s.Scan() + s.Parse()
	if s.IsValid(result) {
		fmt.Println(result + s.Generate())
	} else {
		fmt.Println("Compilation error!")
	}
}

type Scanner struct {
}

func (s *Scanner) Scan() string {
	return "Program "
}

type Parser struct {
}

func (s *Parser) Parse() string {
	return "successfully "
}

type Validator struct {
}

func (s *Validator) IsValid(code string) bool {
	if len(code) != 0 {
		return true
	}
	return false
}

type CodeGenerator struct {
}

func (s *CodeGenerator) Generate() string {
	return "compiled!"
}
