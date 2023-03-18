package main

import "fmt"

/*
	func main() {
		defer second()
		first()
	}

	func first() {
		fmt.Println("first")
	}

	func second() {
		fmt.Println("second")
	}

	func main() {
		defer second()
		first()
		third()
	}

	func first() {
		fmt.Println("first")
	}

	func second() {
		fmt.Println("second")
	}

	func third() {
		defer first()
		fmt.Println("third")
	}

	func main() {
		defer second()
		defer third()
		first()
	}

	func first() {
		fmt.Println("first")
	}

	func second() {
		fmt.Println("second")
	}

	func third() {
		fmt.Println("third")
	}

//panic & recover

	func main() {
		f()
	}

	func final_print(msg string) {
		fmt.Println(msg)
	}

	func f() {
		fmt.Println("f.1")
		g()
		fmt.Println("f.2")
	}

	func g() {
		defer final_print("g.defer()")
		fmt.Println("g.1")
		h()
		fmt.Println("g.2")
	}

	func h() {
		defer final_print("h.defer()")
		fmt.Println("h.1")
		panic("panic in h()")
		fmt.Println("h.2")
	}
*/
func main() {
	f()
}

func final_print(msg string) {
	fmt.Println(msg)
}

func f() {
	fmt.Println("f.1")
	g()
	fmt.Println("f.2")
}

func g() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	defer final_print("g.defer()")
	fmt.Println("g.1")
	h()
	fmt.Println("g.2")
}

func h() {
	defer final_print("h.defer()")
	fmt.Println("h.1")
	panic("panic in h()")
	fmt.Println("h.2")
}
