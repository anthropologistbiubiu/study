package main

import (
	"fmt"
)

type Button struct {
	command Command
}

func (b *Button) Press() {
	b.command.execute()
}

type Command interface {
	execute()
}

type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

type Device interface {
	on()
	off()
}

type Tv struct {
}

func (this *Tv) on() {
	fmt.Println("tv is on")
}

func (this *Tv) off() {
	fmt.Println("tv is off")
}
func main() {
	tv := &Tv{}
	oncommand := &OnCommand{
		device: tv,
	}
	offcommand := &OffCommand{
		device: tv,
	}
	onButton := &Button{
		command: oncommand,
	}
	onButton.Press()

	offButton := &Button{
		command: offcommand,
	}
	offButton.Press()

}
