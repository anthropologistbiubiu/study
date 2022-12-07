package simpleFactory

import (
	"fmt"
)

type API interface{
	say(str string)string
}

type HIAPI struct{

}

func (t *HIAPI)say(hi string)string{
	return fmt.Sprintf("%s",hi)
}

type HelloAPI struct{
}

func (t* HelloAPI) say(hello string)string{
	return fmt.Sprintf("%s",hello)
}


func NewFactory(say string)API{
	if say == "hello"{
		return &HelloAPI{}
	}
	if say== "hi"{
		return &HIAPI{}
	}
	return nil
}