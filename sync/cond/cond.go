package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	condSyncTest()
}

func condSyncTest() {
	//var cond sync.Cond
	var wg sync.WaitGroup
	m := new(sync.Mutex)
	cond := sync.NewCond(m)
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			cond.L.Lock()
			fmt.Printf("准备好了：%v", i)
			fmt.Println(",ready")
			cond.Wait()
			fmt.Println(i, "号到了")
			cond.L.Unlock()
		}(i)
	}
	time.Sleep(5 * time.Second)
	fmt.Println("裁判已经就位，准备发令枪")
	fmt.Println("比赛开始，大家准备跑")
	cond.Broadcast() //统一控制
	/*
		for i := 0; i < 5; i++ {
			cond.Signal() //发令枪响,一个个发
			fmt.Printf("一个一个跑，%d\n", i)
		}
	*/
	wg.Wait()
}

/*
package main



import (
	"log"
	"sync"
	"time"
)

var done = false

func read(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		c.Wait()
	}
	log.Println(name, "starts reading")
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	log.Println(name, "starts writing")
	time.Sleep(time.Second)
	c.L.Lock()
	done = true
	c.L.Unlock()
	log.Println(name, "wakes all")
	c.Broadcast()
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)
	write("writer", cond)
	time.Sleep(time.Second * 3)
}

*/
