package _8_state

import "fmt"

type NoItemState struct {
	vendingMachine *VendingMachine
}

func (s *NoItemState) requestItem() error {
	return fmt.Errorf("item out of stock")
}

func (s *NoItemState) addItem(count int) error {
	s.vendingMachine.incrementItemCount(count)
	s.vendingMachine.setState(s.vendingMachine.hasItem)
	return nil
}

func (s *NoItemState) insertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (s *NoItemState) dispenseItem() error {
	return fmt.Errorf("item out of stock")
}
