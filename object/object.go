package main

import "fmt"

type report interface {
	send()
}

func send(r report) {
	r.send()
}

type user struct {
	name string
	age  int
}

func (this *user) send() {
	fmt.Printf(" %s's age is %d\n", this.name, this.age)
}

type admin struct {
	name string
	age  int
}

func (this *admin) send() {
	fmt.Printf(" %s's age is %d\n", this.name, this.age)
}

func main() {

	u := &user{
		name: "sunweiming",
		age:  18,
	}
	send(u)

	a := &admin{
		name: "yujinling",
		age:  18,
	}
	send(a)
	var p report
	p = u
	p.send()
	p = a
	p.send()
}
