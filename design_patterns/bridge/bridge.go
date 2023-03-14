package main

import "fmt"

type Computer interface {
	setPrinter(Printer)
	Print()
}

type Mac struct {
	printer Printer
}

func (this *Mac) SetPrinter(printer Printer) {
	fmt.Println("Printer Request For Mac")
	this.printer = printer
}

func (this *Mac) Print() {
	this.printer.PrintFile()
}

type Windows struct {
	printer Printer
}

func (this *Windows) setPrinter(printer Printer) {
	fmt.Println("Printer Request For Windows")
	this.printer = printer
}

func (this *Windows) Print() {
	this.printer.PrintFile()
}

type Printer interface {
	PrintFile()
}
type Hp struct {
}

func (this *Hp) PrintFile() {
	fmt.Println("Printing By Hp Printer")
}

type Xiaomi struct {
}

func (this *Xiaomi) PrintFile() {
	fmt.Println("Printing by Xiaomi Printer")
}

func main() {
	mac := &Mac{}
	windows := Windows{}
	hp := &Hp{}
	xiaomi := &Xiaomi{}
	mac.SetPrinter(hp)
	mac.Print()
	mac.SetPrinter(xiaomi)
	mac.Print()
	fmt.Println()

	windows.setPrinter(xiaomi)
	windows.Print()
	windows.setPrinter(hp)
	windows.Print()
}
