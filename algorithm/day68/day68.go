package main

import (
	"fmt"
	"testing"
	"unsafe"
)

type student struct {
	Name string `json:"name"`
}

func f1() *[]student {
	var ms = []student{}
	s1 := student{Name: "sun"}
	s2 := student{Name: "sun"}
	s3 := student{Name: "sun"}
	s4 := student{Name: "sun"}
	s5 := student{Name: "sun"}
	ms = append(ms, s1, s2, s3, s4, s5)
	fmt.Println("f1", unsafe.Sizeof(ms))
	fmt.Println("f1", &ms[0])
	return &ms
}

func usef1() {
	or := f1()
	fmt.Println(unsafe.Sizeof(or))
	fmt.Println("f1", &or)
}
func BenchmarkParallelf1(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			usef1()
		}
	})
}
func f2() []*student {
	var ms = []*student{}
	s1 := new(student)
	s1.Name = "sun"
	s2 := new(student)
	s2.Name = "sun"
	s3 := new(student)
	s3.Name = "sun"
	s4 := new(student)
	s4.Name = "sun"
	s5 := new(student)
	s5.Name = "sun"
	ms = append(ms, s1, s2, s3, s4, s5)
	s1.Name = "sun"
	fmt.Println("f2", unsafe.Sizeof(ms))
	fmt.Println("f2", &ms)
	return ms
}

func usef2() {
	or := f2()
	fmt.Println(unsafe.Sizeof(or))
	fmt.Println("f2", &or)
}
func BenchmarkParallelf2(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			usef2()
		}
	})
}

func main() {
	usef1()
	usef2()
}
