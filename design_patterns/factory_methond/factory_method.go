package main

import "fmt"

type IGun interface {
	setName(name string)
	getName() string
	setPow(pow int)
	getPow() int
}

type Gun struct {
	name string
	pow  int
}

func (this *Gun) setName(name string) {
	this.name = name
}

func (this *Gun) getName() string {
	return this.name
}
func (this *Gun) setPow(pow int) {
	this.pow = pow
}
func (this *Gun) getPow() int {
	return this.pow
}

type Ak47 struct {
	gun IGun
}

func NewAK47() IGun {
	return &Gun{
		name: "AK47",
		pow:  1,
	}
}

type musket struct {
	Gun IGun
}

func NewMusket() IGun {
	return &Gun{
		name: "Musket",
		pow:  1,
	}
}

func Factory(method string) IGun {
	if method == "ak" {
		return NewAK47()
	} else if method == "musk" {
		return NewMusket()
	}
	return nil
}

func main() {
	ak := Factory("ak")
	musk := Factory("musk")
	fmt.Println(ak)
	fmt.Println(musk)
}
