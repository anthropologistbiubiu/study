package main

import "fmt"

type IdCard string

func (this *IdCard) GetBirtyDay() IdCard {
	return (*this)[6:]
}

func NewIdCard(nums string) IdCard {
	return IdCard(nums)
}
func main() {

	card := NewIdCard("12345678910")
	fmt.Println(card.GetBirtyDay())

}
