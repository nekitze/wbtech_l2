package pattern

/*
Стратегия - поведенческий паттерн. Определяет семейство алгоритмов,
икапсулирует каждый из них и делает их взаимозаменяемыми.
Стратегия позволяет изменять алгоритмы независимо от клиентов, которые
ими пользуются.

Применяется, когда:
- нужно использовать разные вариации какого-то алгоритма внутри одного объекта;
- есть множество похожих классов, отличающихся только некоторым поведением;
- не хотим раскрывать детали реализации алгоритмов для других классов;
- различные вариации алгоритмов реализованы в виде развесистого условного оператора,
а каждая ветка такого оператора представляет собой вариацию алгоритма;

Плюсы:
- замена алгоритма на лету;
- инкапсуляция алгоритмов;
- уход от наследования к делегированию;
- реализует прицип открытости/закрытости;
Минусы:
- клиент должен знать, в чём разница между стратегиями, чтобы выбрать подходящую;
- усложнения программы за счёт дополнительных классов;

Пример реального применения: выбор алгоритма обработки при получении данных в различных;
форматах (JSON, XML, etc..); выбор алгоритма в зависимости от способа оплаты на сайте;
*/

type Payment interface {
	Pay() error
}

type cardPayment struct {
	cardNumber string
	cvv        string
}

func NewCardPayment(cardNumber, cvv string) Payment {
	return &cardPayment{
		cardNumber: cardNumber,
		cvv:        cvv,
	}
}

func (p *cardPayment) Pay() error {
	return nil
}

type payPalPayment struct {
}

func NewPayPalPayment() Payment {
	return &cardPayment{}
}

func (p *payPalPayment) Pay() error {
	return nil
}

type qiwiPayment struct {
}

func NewQIWIPayment() Payment {
	return &cardPayment{}
}

func (p *qiwiPayment) Pay() error {
	return nil
}

// processOrder не должен знать, какой именно алгоритм будет применён
func processOrder(product string, payment Payment) {
	err := payment.Pay()
	if err != nil {
		return
	}
}

func someMethod1() {
	product := "vehicle"

	payWay := 3

	// Определяем алгоритм в зависимости от способа оплаты
	var payment Payment
	switch payWay {
	case 1:
		payment = NewCardPayment("5536 9138 0516", "221")
	case 2:
		payment = NewPayPalPayment()
	case 3:
		payment = NewQIWIPayment()
	}

	processOrder(product, payment)
}
