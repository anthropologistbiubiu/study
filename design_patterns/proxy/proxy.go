package main

import "fmt"

type Server interface {
	handleRequest(url, method string) error
}

type Nginx struct {
	application   *Application
	maxAllowVisit int
	ratelimiter   map[string]int
}

func NewNginx() *Nginx {
	return &Nginx{
		application:   &Application{},
		maxAllowVisit: 2,
		ratelimiter:   make(map[string]int),
	}
}
func (this *Nginx) AllowVisit(url string) error {
	if _, ok := this.ratelimiter[url]; !ok {
		this.ratelimiter[url] = 1
		return nil
	}
	this.ratelimiter[url] += 1
	if this.ratelimiter[url] > this.maxAllowVisit {
		return fmt.Errorf("visited is limited")
	}
	return nil

}
func (this *Nginx) handleRequest(url, method string) error {
	allow := this.AllowVisit(url)
	if allow == nil {
		return this.application.handleRequest(url, method)
	}
	return fmt.Errorf("403 not allowed")
}

type Application struct {
}

func (this *Application) handleRequest(url, method string) error {
	if url == "/apple" && method == "post" {
		fmt.Println("apple")
		return nil
	}
	if url == "/orange" && method == "post" {
		fmt.Println("orange")
		return nil
	}
	return fmt.Errorf("404 not found\n")
}

func main() {
	nginx := NewNginx()
	err := nginx.handleRequest("/apple", "post")
	if err != nil {
		fmt.Println(err)
	}
	err = nginx.handleRequest("/apple", "post")
	if err != nil {
		fmt.Println(err)
	}
	err = nginx.handleRequest("/apple", "post")
	if err != nil {
		fmt.Println(err)
	}
	err = nginx.handleRequest("/app", "post")
	if err != nil {
		fmt.Println(err)
	}

	err = nginx.handleRequest("/orange", "post")
	if err != nil {
		fmt.Print(err)
	}
	err = nginx.handleRequest("/orange", "get")
	if err != nil {
		fmt.Print(err)
	}
}
