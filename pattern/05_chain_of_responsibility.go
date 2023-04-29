package pattern

import "fmt"

// Цепочка обязанностей — это поведенческий паттерн проектирования, который позволяет передавать
// запросы последовательно по цепочке обработчиков. Каждый последующий обработчик решает,
// может ли он обработать запрос сам и стоит ли передавать запрос дальше по цепи.

type Patient struct {
	Name              string
	RegistrationDone  bool
	DoctorCheckUpDone bool
	MedicineDone      bool
	PaymentDone       bool
}

// Department – общий интерфейс обработчиков
type Department interface {
	Execute(*Patient)
	SetNext(Department)
}

// Reception – первый обработчик
type Reception struct {
	next Department
}

func (r *Reception) Execute(p *Patient) {
	if p.RegistrationDone {
		fmt.Printf("Patient %s registration already done\n", p.Name)
		r.next.Execute(p)
		return
	}
	fmt.Printf("Receprion registering patient %s\n", p.Name)
	p.RegistrationDone = true
	r.next.Execute(p)
}

func (r *Reception) SetNext(next Department) {
	r.next = next
}

// Doctor – второй обработчик
type Doctor struct {
	next Department
}

func (d *Doctor) Execute(p *Patient) {
	if p.DoctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.Execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.DoctorCheckUpDone = true
	d.next.Execute(p)
}

func (d *Doctor) SetNext(next Department) {
	d.next = next
}

// Medical – третий обработчик
type Medical struct {
	next Department
}

func (m *Medical) Execute(p *Patient) {
	if p.MedicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.Execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.MedicineDone = true
	m.next.Execute(p)
}

func (m *Medical) SetNext(next Department) {
	m.next = next
}

// Cashier – четвёртый обработчик
type Cashier struct {
	next Department
}

func (c *Cashier) Execute(p *Patient) {
	if p.PaymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient")
}

func (c *Cashier) SetNext(next Department) {
	c.next = next
}

func Run05() {
	cashier := &Cashier{}

	medical := &Medical{}
	medical.SetNext(cashier)

	doctor := &Doctor{}
	doctor.SetNext(medical)

	reception := &Reception{}
	reception.SetNext(doctor)

	patient := &Patient{Name: "Bob"}
	reception.Execute(patient)
}
