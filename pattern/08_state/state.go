package _8_state

type State interface {
	addItem(int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}
