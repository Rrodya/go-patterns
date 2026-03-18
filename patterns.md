# Поведенческие паттерны (Behavioral)

### Паттерны уровня класса

Описывают взаимодействие между классами и их подклассами. Такие отношения выражаются наследованным и реализацией классов. Базовый класс определяет интерфейс, а подклассы - реализацию

### Паттерны уровня объекта

Описываются взаимодействия между объектами. Такие отношения выражаются связями - ассоциацией, агрегацией и композицией.
Структура строятся путем объеденения объектом некоторых классов.

- Ассоциация объекты двух ссылаются один на другой. Просто знают друг о друге

```go
type Recipe struct {
    Name string
}

type Chef struct {
    Name string
    FavoriteRecipe *Recipe
}
```

- Агрегация частная форма ассоциации. Объект-контейнер содержит другие объекта, но они могут жить отдельно

```go
type Employee struct {
	Name string
	Position string
}

type Department struct {
	Name string
	Empoyees []*Empoyees
}

func (d *Department) AddEmployee(e *Employee) {
	d.Employees = append(d.Employees, e)
}

func main() {
	emp1 := &Employee{Name: "Anna", Position: "Developer"}
	emp1 := &Employee{Name: "Boris", Position: "QA"}

	it := &Department{Name: "IT"}
	it.AddEmployee(emp1)
	it.AddEmployee(emp2)

	hr := &Department{Name: "HR"}
	hr.AddEmployee(emp1)

	it = nil

	fmt.Printf("Anna still exists %+v\n", emp1)
}
```

- Композиция тоже самое что и агрегация только составные объекты не могут существовать отдельно от объекта контейнера и если контейнер будет уничтожить то и все содержимое будет уничтожено

```go
type Engine struct {
	HorsePower int
}

type Car struct {
	Brand string
	Engine Engine
}

func NewCar(brand string) *Car {
	return &Car{
		Brand string
		Engine Engine{Horsepower: 150}
	}
}

func main() {
	car := NewCar("Toyota")
	fmt.Printf("Car: %s, Engine: %d h.p\n", car.Brand, car.Engine.HorsePower)

	car = nil
}
```
