package pattern

// Фасад — это структурный паттерн проектирования, который предоставляет простой интерфейс к сложной системе классов,
// библиотеке или фреймворку.

import "fmt"

type IgnitionSystem struct{}

func (is *IgnitionSystem) TurnOn() {
	fmt.Println("Ignition turned on")
}

func (is *IgnitionSystem) TurnOff() {
	fmt.Println("Ignition turned off")
}

type FuelPump struct{}

func (fp *FuelPump) Pump() {
	fmt.Println("Fuel pumped")
}

type Starter struct{}

func (s *Starter) Spin() {
	fmt.Println("Starter is spinning")
}

func (s *Starter) StopSpinning() {
	fmt.Println("Starter stopped spinning")
}

type EngineFacade struct {
	ignitionSystem IgnitionSystem
	fuelPump       FuelPump
	starter        Starter
}

func (e *EngineFacade) Start() {
	e.ignitionSystem.TurnOn()
	e.fuelPump.Pump()
	e.starter.Spin()
	fmt.Println("Engine started")
	e.starter.StopSpinning()
}

func (e *EngineFacade) Stop() {
	e.ignitionSystem.TurnOff()
	fmt.Println("Engine stopped")
}
