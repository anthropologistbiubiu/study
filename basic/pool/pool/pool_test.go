package pool

import (
	"sync"
	"testing"
)

type A struct {
	Name string
}

func (a *A) Reset() {
	a.Name = ""
}

var pool = sync.Pool{
	New: func() interface{} {
		return new(A)
	},
}

func BenchmarkWithoutPool(b *testing.B) {
	var a *A
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			a = new(A)
			a.Name = "tink"
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	var a *A
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			a = pool.Get().(*A)
			a.Reset()
			a.Name = "tink"
			pool.Put(a) // 一定要记得放回操作，否则退化到每次都需要New操作
		}
	}
}

//pool]$ go test -bench .
//goos: linux
//goarch: amd64
//pkg: pool
//cpu: Intel(R) Xeon(R) Platinum 8269CY CPU @ 2.50GHz
//BenchmarkWithoutPool-8              3736            277442 ns/op          160001 B/op      10000 allocs/op
//BenchmarkWithPool-8                 7821            151826 ns/op               0 B/op          0 allocs/op
//PASS
//ok      pool    2.277s
