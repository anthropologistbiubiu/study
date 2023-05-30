package test

import (
	"testing"
)

type student struct {
	Name string `json:"name"`
}

func f1() *[]*student {
	var ms = make([]*student, 0, 10)
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
	/*
		s1 := student{Name: "sun"}
		s2 := student{Name: "sun"}
		s3 := student{Name: "sun"}
		s4 := student{Name: "sun"}
		s5 := student{Name: "sun"}
	*/
	ms = append(ms, s1, s2, s3, s4, s5)
	return &ms
}

func usef1() {
	or := f1()
	//fmt.Println(or)
	*or = append(*or, &student{Name: "sun"})
}
func BenchmarkParallelf1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		usef1()
	}
}
func f2() []*student {
	var ms = make([]*student, 0, 5)
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
	return ms
}

func usef2() {
	or := f2()
	//fmt.Println(or)
	or = append(or, &student{Name: "sun"})
}
func BenchmarkParallelf2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		usef2()
	}
}
