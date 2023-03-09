package main

import "fmt"

type Collection interface {
	CreateIterator() Iterator
}

type useCollection struct {
	users []*user
}

func (this *useCollection) CreateIterator() *userIterator {
	return &userIterator{
		users: this.users,
	}
}

type user struct {
	name string
	age  int
}

type userIterator struct {
	index int
	users []*user
}
type Iterator interface {
	hasNext() bool
	getNext() Iterator
}

func (this *userIterator) hasNext() bool {
	if this.index < len(this.users) {
		return true
	}
	return false
}
func (this *userIterator) getNext() *user {
	if this.hasNext() {
		user := this.users[this.index]
		this.index++
		return user
	}
	return nil

}
func main() {
	user1 := &user{
		name: "user1",
		age:  10,
	}
	user2 := &user{
		name: "user2",
		age:  10,
	}
	userCollecion := &useCollection{
		users: []*user{user1, user2},
	}
	useriterator := userCollecion.CreateIterator()
	for useriterator.hasNext() {
		fmt.Printf("%+v\n", useriterator.getNext())
	}

}
