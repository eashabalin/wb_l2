package pattern

import "fmt"

// Посетитель — это поведенческий паттерн проектирования, который позволяет добавлять в программу новые операции,
// не изменяя классы объектов, над которыми эти операции могут выполняться.
// Представим, что надо сделать экспорт данных для объектов разных типов: Dot, Circle, Rectangle, CompoundShape
// https://refactoring.guru/ru/design-patterns/visitor

type Shape interface {
	Move(x, y int)
	Draw()
	Accept(v Visitor)
}

type Dot struct {
	X, Y int
}

func (d *Dot) Move(x, y int) {
	fmt.Printf("Moved to (%d, %d)\n", x, y)
}

func (d *Dot) Draw() {
	fmt.Println(".")
}

func (d *Dot) Accept(v Visitor) {
	v.DoForDot(d)
}

type Circle struct {
	X, Y, Radius int
}

func (c *Circle) Move(x, y int) {
	fmt.Printf("Moved to (%d, %d)\n", x, y)
}

func (c *Circle) Draw() {
	fmt.Println("o")
}

func (c *Circle) Accept(v Visitor) {
	v.DoForCircle(c)
}

type Rectangle struct {
	X, Y, A, B int
}

func (r *Rectangle) Move(x, y int) {
	fmt.Printf("Moved to (%d, %d)\n", x, y)
}

func (r *Rectangle) Draw() {
	fmt.Println("п")
}

func (r *Rectangle) Accept(v Visitor) {
	v.DoForRectangle(r)
}

type CompoundShape struct {
	X, Y, R1, R2 int
}

func (cs *CompoundShape) Move(x, y int) {
	fmt.Printf("Moved to (%d, %d)\n", x, y)
}

func (cs *CompoundShape) Draw() {
	fmt.Println("S")
}

func (cs *CompoundShape) Accept(v Visitor) {
	v.DoForCompoundShape(cs)
}

type Visitor interface {
	DoForDot(d *Dot)
	DoForCircle(c *Circle)
	DoForRectangle(r *Rectangle)
	DoForCompoundShape(cs *CompoundShape)
}

type XMLExportVisitor struct{}

func (v *XMLExportVisitor) DoForDot(d *Dot) {
	fmt.Printf("<dot x=%d y=%d>\n", d.X, d.Y)
}

func (v *XMLExportVisitor) DoForCircle(c *Circle) {
	fmt.Printf("<circle x=%d y=%d r=%d>\n", c.X, c.Y, c.Radius)
}

func (v *XMLExportVisitor) DoForRectangle(r *Rectangle) {
	fmt.Printf("<rectangle x=%d y=%d a=%d b=%d>\n", r.X, r.Y, r.A, r.B)
}

func (v *XMLExportVisitor) DoForCompoundShape(cs *CompoundShape) {
	fmt.Printf("<cshape x=%d y=%d r1=%d r2=%d>\n", cs.X, cs.Y, cs.R1, cs.R2)
}
