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

func (this *VendingMatchine) requestItem() error {
	fmt.Println("v is requesting item")
	return this.currentItemState.requestItem()
}
func (this *VendingMatchine) addItemCount(count int) error {
	fmt.Println("v is add item")
	return this.currentItemState.addItem(count)
}
func (this *VendingMatchine) insertMoney(money int) error {
	fmt.Println("v is insert into money")
	return this.currentItemState.insertMoney(money)
}
func (this *VendingMatchine) dispenseItem() error {
	fmt.Println("v is dispenseItem  ")
	return this.currentItemState.dispenseItem()
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
	addItem(int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}

type hasItemState struct {
	vendingMatchine *VendingMatchine
}

func (this *hasItemState) addItem(count int) error {
	return fmt.Errorf("item is enough")
}
func (this *hasItemState) requestItem() error {
	if this.vendingMatchine.itemCount == 0 {
		this.vendingMatchine.SetCurrentState(this.vendingMatchine.noItemState)
		return fmt.Errorf("No enough Item Present,please supplyment item")
	}
	fmt.Println("request Item success!")
	this.vendingMatchine.SetCurrentState(this.vendingMatchine.requestItemState)
	return nil
}
func (this *hasItemState) insertMoney(money int) error {
	return fmt.Errorf("please request Item first")
}
func (this *hasItemState) dispenseItem() error {
	return fmt.Errorf("please request Item first")
}

type requestItemState struct {
	vendingMatchine *VendingMatchine
}

func (this *requestItemState) addItem(count int) error {
	return fmt.Errorf("item has already requested please insert into money")
}
func (this *requestItemState) requestItem() error {
	return fmt.Errorf("item has already requested please insert into money")
}
func (this *requestItemState) insertMoney(money int) error {
	if money < this.vendingMatchine.itemPrice {
		return fmt.Errorf("insert money is not enough")
	}
	fmt.Println("payment success!")
	this.vendingMatchine.SetCurrentState(this.vendingMatchine.hasMoneyState)
	return nil
}
func (this *requestItemState) dispenseItem() error {
	return fmt.Errorf("item request is in porcess")
}

type hasMoneyState struct {
	vendingMatchine *VendingMatchine
}

func (this *hasMoneyState) addItem(count int) error {
	return fmt.Errorf("waiting for dispenseItem ")
}
func (this *hasMoneyState) requestItem() error {
	return fmt.Errorf("waiting for dispenseItem ")
}
func (this *hasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("waiting for dispenseItem ")
}
func (this *hasMoneyState) dispenseItem() error {
	this.vendingMatchine.IncreaseItemCount(-1)
	if this.vendingMatchine.itemCount == 0 {
		this.vendingMatchine.SetCurrentState(this.vendingMatchine.noItemState)
	} else {
		this.vendingMatchine.SetCurrentState(this.vendingMatchine.hasItemState)
	}
	return nil
}

type noItemState struct {
	vendingMatchine *VendingMatchine
}

func (this *noItemState) addItem(count int) error {
	fmt.Println("we are supplement item")
	this.vendingMatchine.IncreaseItemCount(this.vendingMatchine.itemCount)
	return nil
}
func (this *noItemState) requestItem() error {
	return fmt.Errorf("item out of stock")
}
func (this *noItemState) insertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}
func (this *noItemState) dispenseItem() error {
	return fmt.Errorf("item out of stock")
}
func main() {
	v := NewVendingMatchine(10, 0)
	err := v.requestItem()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = v.insertMoney(8)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = v.dispenseItem()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println()
	v.addItemCount(1)
	err = v.requestItem()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = v.insertMoney(8)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = v.dispenseItem()
	if err != nil {
		fmt.Println(err.Error())
	}
}
