package main

import "fmt"

/*
Строитель - порождающий паттерн проектирования, который позволяет создавать
сложные объекты пошагово. Строитель даёт возможность использовать один и
тот же код строительства для получения разных представлений объектов.

Применяется, когда:
- появляется конструктор с большим количеством опциональных параметров;
- нужно собирать сложные составные объекты;
- создание нескольких представлений объекта состоит из одинаковых этапов,
  которые отличаются в деталях;

Плюсы:
- пошаговое создание объектов;
- переиспользование кода создания различных объектов;
- изоляция сложного кода сборки объекта от основной бизнес-логики;
Минусы:
- усложнение кода дополнительными классами;

Пример реального применения: StringBuilder
*/

type Database struct {
	User     string
	Password string
	Url      string
	Name     string
	Driver   string
}

type DatabaseBuilderI interface {
	User(val string) DatabaseBuilderI
	Password(val string) DatabaseBuilderI
	Url(val string) DatabaseBuilderI
	Name(val string) DatabaseBuilderI
	Driver(val string) DatabaseBuilderI

	Connect() Database
}

type databaseBuilder struct {
	database Database
}

func NewDatabaseBuilder() DatabaseBuilderI {
	return databaseBuilder{database: Database{}}
}

func (b databaseBuilder) User(val string) DatabaseBuilderI {
	b.database.User = val
	return b
}

func (b databaseBuilder) Password(val string) DatabaseBuilderI {
	b.database.Password = val
	return b
}

func (b databaseBuilder) Url(val string) DatabaseBuilderI {
	b.database.Url = val
	return b
}

func (b databaseBuilder) Name(val string) DatabaseBuilderI {
	b.database.Name = val
	return b
}

func (b databaseBuilder) Driver(val string) DatabaseBuilderI {
	b.database.Driver = val
	return b
}

func (b databaseBuilder) Connect() Database {
	fmt.Printf("Connected to db: %v\n", b)
	return b.database
}

// PostgresDatabaseBuilder :
// Можно создать конкретный билдер для Postgres
type PostgresDatabaseBuilder struct {
	databaseBuilder
}

func NewPostgresDatabaseBuilder() DatabaseBuilderI {
	return databaseBuilder{}.Name("postgres").Url("pgsql://localhost:5432/postgres").Driver("Postgres")
}

// Director :
// Этому объекту можно задать билдер, который он будет использовать
// для создания объектов с нужными параметрами
type Director struct {
	builder DatabaseBuilderI
}

func NewDirector(b DatabaseBuilderI) *Director {
	return &Director{builder: b}
}

func (d *Director) SetBuilder(b DatabaseBuilderI) {
	d.builder = b
}

func (d *Director) ConnectDatabase() Database {
	return d.builder.Connect()
}

// Пример использования
func main() {
	// Создаём подключение полностью с нуля
	mysqlBuilder := NewDatabaseBuilder().
		User("mysql").
		Password("2222").
		Url("mysql://localhost:5432/postgres").
		Name("postgres").
		Driver("Postgres")
	mysqlDB := mysqlBuilder.Connect()

	// Используем готовый билдер для Postgres
	// Можем изменить некоторые опции, например пользователя
	postgresBuilder := NewPostgresDatabaseBuilder()
	postgresBuilder.User("anotherUser")
	postgresDB := postgresBuilder.Connect()

	fmt.Println("My Databases:", mysqlDB, postgresDB)

	// Использование Director
	// Создаем, передавая билдер
	director := NewDirector(mysqlBuilder)
	mysqlDBFromDirector := director.ConnectDatabase()

	// Устанавливаем другой билдер
	director.SetBuilder(postgresBuilder)
	postgresDBFromDirector := director.ConnectDatabase()

	fmt.Println("Director's Databases:", mysqlDBFromDirector, postgresDBFromDirector)
}
