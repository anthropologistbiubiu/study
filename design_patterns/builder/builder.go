package main

import "fmt"

type IBuilder interface {
	setWindowsType()
	setFloorType()
	setFloorNum()
	getHouse() House
}

type House struct {
	windowsType string
	floorType   string
	floorNum    int
}

type NormalBuilder struct {
	windowsType string
	floorType   string
	floorNum    int
}

func (this *NormalBuilder) setWindowsType() {
	this.windowsType = "wood windows"
}
func (this *NormalBuilder) setFloorType() {
	this.floorType = "wood windows"
}
func (this *NormalBuilder) setFloorNum() {
	this.floorNum = 1
}
func (this *NormalBuilder) getHouse() House {
	return House{
		windowsType: this.windowsType,
		floorType:   this.floorType,
		floorNum:    this.floorNum,
	}
}

func NewNormalBuilder() IBuilder {
	return &NormalBuilder{}
}

type IglooBuilder struct {
	windowsType string
	floorType   string
	floorNum    int
}

func (this *IglooBuilder) setWindowsType() {
	this.windowsType = "Igloo windows"
}
func (this *IglooBuilder) setFloorType() {
	this.floorType = "Igloo windows"
}
func (this *IglooBuilder) setFloorNum() {
	this.floorNum = 1
}
func (this *IglooBuilder) getHouse() House {
	return House{
		windowsType: this.windowsType,
		floorType:   this.floorType,
		floorNum:    this.floorNum,
	}
}
func NewIglooBuilder() IBuilder {
	return &IglooBuilder{}
}

type Director struct {
	builder IBuilder
}

func NewDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (this *Director) setBuilder(b IBuilder) {
	this.builder = b
}

func main() {

	normal := NewNormalBuilder()
	normal.setWindowsType()
	normal.setFloorType()
	normal.setFloorNum()
	director := NewDirector(normal)
	house := director.builder.getHouse()
	fmt.Printf("house : %+v\n", house)
}
