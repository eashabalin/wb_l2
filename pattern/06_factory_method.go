package pattern

import "fmt"

// Фабричный метод — это порождающий паттерн проектирования, который определяет общий интерфейс
// для создания объектов в суперклассе, позволяя подклассам изменять тип создаваемых объектов.

// Данный шаблон делегирует создание объектов наследникам родительского класса.
// Это позволяет использовать в коде программы не конкретные классы, а манипулировать абстрактными объектами
// на более высоком уровне.

type Transport interface {
	Deliver()
}

type Truck struct{}

func (t *Truck) Deliver() {
	fmt.Println("delivering by truck...")
}

type Ship struct{}

func (s *Ship) Deliver() {
	fmt.Println("delivering by ship...")
}

// Logistics – базовый создатель объекта Transport
type Logistics struct{}

func (l *Logistics) planDelivery() {
	fmt.Println("delivery planned")
}

// CreateTransport по умолчанию создаёт объект Truck. В дочерних структурах структуры Logistics метод может
// быть переопределён для создания других объектов интерфейса Transport.
func (l *Logistics) CreateTransport() Transport {
	return &Truck{}
}

// RoadLogistics – дочерний создатель объекта Transport
type RoadLogistics struct {
	Logistics
}

func (rl *RoadLogistics) CreateTransport() Transport {
	return &Truck{}
}

// SeaLogistics – дочерний создатель объекта Transport
type SeaLogistics struct {
	Logistics
}

// CreateTransport – метод переопределён, чтобы возвращать Ship
func (rl *SeaLogistics) CreateTransport() Transport {
	return &Ship{}
}
