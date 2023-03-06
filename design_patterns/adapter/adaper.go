package main

import "fmt"

type client struct {
}

type Computer interface {
	InsertIntoLightningPort()
}

func (c *client) InsertLightningConnectorIntoComputer(com Computer) {
	com.InsertIntoLightningPort()
}

type Mac struct {
}

func (this *Mac) InsertIntoLightningPort() {
	fmt.Println("Insert Into Mac Port")
}

type Windows struct {
}

func (this *Windows) InsertIntoLightningPort() {
	fmt.Println("Insert Into Windows Port")
}

type WindowsAdaptor struct {
	WindowMachine *Windows
}

func (this *WindowsAdaptor) InsertIntoLightningPort(com Computer) {
	fmt.Println("convert windows signal into lighting signal")
	com.InsertIntoLightningPort()

}
func main() {
	mac := &Mac{}
	mac.InsertIntoLightningPort()
	windows := &Windows{}
	WindowsAdaptor := &WindowsAdaptor{}
	WindowsAdaptor.InsertIntoLightningPort(windows)

}
