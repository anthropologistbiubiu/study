package simpleFactory

import(
	"testing"
)

func Test1(t *testing.T){
	api := NewFactory("hi")
	say := api.say("hi")
	if say != "hi"{
		t.Fatal("testing error")
	}
}
func Test2(t *testing.T){
	api := NewFactory("hello")
	say := api.say("hello")
	if say != "hello"{
		t.Fatal("testing error")
	}
}