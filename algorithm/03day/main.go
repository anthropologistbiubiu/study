package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var mapLock sync.Mutex

type IPVisited struct {
	ipvisited map[string]struct{}
}

func (this *IPVisited) Visited(ip string) bool {
	mapLock.Lock()
	defer mapLock.Unlock()
	if _, ok := this.ipvisited[ip]; ok {
		return false
	}
	fmt.Printf("%s is  visiting!\n", ip)
	this.ipvisited[ip] = struct{}{}
	go this.DeleteWaitForMinute(ip)
	return true
}

func (this *IPVisited) DeleteWaitForMinute(ip string) {
	fmt.Println("........")
	time.Sleep(time.Minute * 1)
	mapLock.Lock()
	delete(this.ipvisited, ip)
	mapLock.Unlock()
}
func NewAllowIp() *IPVisited {
	return &IPVisited{
		ipvisited: make(map[string]struct{}, 0),
	}
}
func main() {
	fmt.Println(runtime.NumCPU())
	wg := sync.WaitGroup{}
	ipm := NewAllowIp()
	for i := 0; i < 100; i++ {
		ipEnd := i
		ip := fmt.Sprintf("192.168.1.%d", ipEnd)
		for j := 0; j < 1; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if ipm.Visited(ip) {
					//fmt.Printf("%s is already visited! please waiting for a minute \n", ip)
				}
			}()
		}
	}
	wg.Wait()
}
