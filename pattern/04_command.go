package pattern

// Команда — это поведенческий паттерн, позволяющий заворачивать запросы или простые операции в отдельные объекты.
// Это позволяет откладывать выполнение команд, выстраивать их в очереди, а также хранить историю и делать отмену.

import "fmt"

type Button struct {
	Command Command
}

func (b *Button) Press() {
	b.Command.execute()
}

type Command interface {
	execute()
}

type OnCommand struct {
	Device Device
}

func (c *OnCommand) execute() {
	c.Device.on()
}

type OffCommand struct {
	Device Device
}

func (c *OffCommand) execute() {
	c.Device.off()
}

type Device interface {
	on()
	off()
}

type TV struct {
	isRunning bool
}

func (tv *TV) on() {
	tv.isRunning = true
	fmt.Println("Turning TV on")
}

func (tv *TV) off() {
	tv.isRunning = false
	fmt.Println("Turning TV off")
}
