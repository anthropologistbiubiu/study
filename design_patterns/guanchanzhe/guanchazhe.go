package main

import "fmt"

type Observer interface {
	update(itemName string)
	getId() string
}
type subject interface {
	regesiter(o Observer)
	deregesiter(o Observer)
	notifyAll()
}

type Item struct {
	observerList []Observer
	name         string
	inStock      bool
}

func (i *Item) regesiter(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) deregesiter(o Observer) {

}

func (i *Item) notifyAll() {
	i.inStock = true
	for _, v := range i.observerList {
		v.update(i.name)
	}
}

type observer struct {
	name string
}

func (o *observer) update(itemName string) {
	fmt.Printf("send email to %+v item:%+v is instock\n", o.name, itemName)
}
func (o *observer) getId() string {
	return o.name
}

func NewObserver(name string) Observer {
	return &observer{
		name: name,
	}
}
func NewItem(name string) subject {
	return &Item{
		name: name,
	}
}

func main() {

	item := NewItem("iphone13")

	o1 := NewObserver("oberver1")
	o2 := NewObserver("oberver2")
	item.regesiter(o1)
	item.regesiter(o2)
	item.notifyAll()
}
