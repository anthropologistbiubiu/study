package mian

func main() {

	var workChan = make(chan int)
	<-workChan
	println("end")
}
