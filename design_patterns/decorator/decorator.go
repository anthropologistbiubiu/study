package main

import "fmt"

type IPizza interface {
	getPrice() int
}

type VeggeMania struct {
}

func (this *VeggeMania) getPrice() int {
	return 15

}

type tomatoTopping struct {
	pizza IPizza
}

func (this *tomatoTopping) getPrice() int {
	c := this.pizza.getPrice()
	return c + 15
}

type chessTopping struct {
	pizza IPizza
}

func (this *chessTopping) getPrice() int {
	c := this.pizza.getPrice()
	return c + 16
}

func main() {
	pizza := &VeggeMania{}

	tomatoPizza := &tomatoTopping{
		pizza: pizza,
	}

	fmt.Printf("tomatoPizza is %+v nums is %d\n", tomatoPizza, tomatoPizza.getPrice())
	fmt.Println(tomatoPizza.pizza.getPrice())
}
