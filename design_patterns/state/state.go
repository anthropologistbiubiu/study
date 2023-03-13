package main

import "fmt"

type VendingMatchine struct {
	hasItemState     State
	requestItemState State
	hasMoneyState    State
	noItemState      State
	currentItemState State
	itemPrice        int
	itemCount        int
}

func (this *VendingMatchine) SetCurrentState(s State) {
	this.currentItemState = s
}

func (this *VendingMatchine) IncreaseItemCount(count int) {
	fmt.Println("vendingMatchine increase item amount")
	this.itemCount = +count
}

func NewVendingMatchine(itemPrice, itemCount int) *VendingMatchine {

	vend := &VendingMatchine{
		itemPrice: itemPrice,
		itemCount: itemCount,
	}
	HasItemState := &hasItemState{
		vendingMatchine: vend,
	}
	RequestItemState := &requestItemState{
		vendingMatchine: vend,
	}
	HasMoneyState := &hasMoneyState{
		vendingMatchine: vend,
	}
	NoItemState := &noItemState{
		vendingMatchine: vend,
	}
	vend.SetCurrentState(HasItemState)
	vend.hasItemState = HasItemState
	vend.noItemState = NoItemState
	vend.requestItemState = RequestItemState
	vend.hasItemState = HasMoneyState
	return vend
}

type State interface {
	addItemState(int) error
	requestItemState() error
	insertItemState()
	dispenseItemState()
}

type hasItemState struct {
	vendingMatchine *VendingMatchine
}

func (this *hasItemState) addItemState(count int) error {
	fmt.Println("add itemCount")
	this.vendingMatchine.IncreaseItemCount(count)
	return nil
}
func (this *hasItemState) requestItemState() error {
	if this.vendingMatchine.itemCount == 0 {
		this.vendingMatchine.SetCurrentState(this.vendingMatchine.noItemState)
		return fmt.Errorf("No Item Present")
	}
	fmt.Println("request Item!")
	this.vendingMatchine.itemCount = this.vendingMatchine.itemCount - 1
	return nil
}
func (this *hasItemState) insertItemState() {

}
func (this *hasItemState) dispenseItemState() {

}

type requestItemState struct {
	vendingMatchine *VendingMatchine
}

func (this *requestItemState) addItemState() error {

}
func (this *requestItemState) requestItemState() {

}
func (this *requestItemState) insertItemState() {

}
func (this *requestItemState) dispenseItemState() {

}

type hasMoneyState struct {
	vendingMatchine *VendingMatchine
}

func (this *hasMoneyState) addItemState() error {

}
func (this *hasMoneyState) requestItemState() {

}
func (this *hasMoneyState) insertItemState() {

}
func (this *hasMoneyState) dispenseItemState() {

}

type noItemState struct {
	vendingMatchine *VendingMatchine
}

func (this *noItemState) addItemState() error {

}
func (this *noItemState) requestItemState() {

}
func (this *noItemState) insertItemState() {

}
func (this *noItemState) dispenseItemState() {

}
