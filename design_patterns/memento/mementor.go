package main

type originator struct {
	state string
}

func (this *originator) setState(state string) {
	this.state = state
}

func (this *originator) restorstate(m mementor) {
	this.state = m.state
}

func (this *originator) CreateMementor() *mementor {
	return &mementor{
		state: this.state,
	}
}

func (this *originator) getState() string {
	return this.state
}

type mementor struct {
	state string
}

func (this *mementor) getSaveState() string {
	return this.state
}

type CareTaker struct {
	mementorArray []mementor
}

func (this *CareTaker) NewCareTaker() *CareTaker {
	return &CareTaker{
		mementorArray: make([]mementor, 0, 10),
	}
}
func (this *CareTaker) addMementor(m mementor) {
	this.mementorArray = append(this.mementorArray, m)

}
func (this *CareTaker) restorState(index int) mementor {
	return this.mementorArray[index]
}

func main() {

}
