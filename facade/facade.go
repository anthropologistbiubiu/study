package facade

import(
	"fmt"
)
type AModel interface{
	TestA()string
}

type aModel struct{}

func(a *aModel)TestA()string{
	return "AModel"
}
func NewAModel()AModel{
	return &aModel{}
}

type BModel interface{
	TestB()string
}
type bModel struct{}

func(b *bModel)TestB()string{
	return "BModel"
}
func NewBMode()BModel{
	return &bModel{}
}

type API interface{
	Test()string
}
type ABMdel struct{
	a AModel
	b BModel	
}
func (c *ABMdel)Test()string{
	aStr := c.a.TestA()
	bStr :=c.b.TestB()
	return fmt.Sprintf("%s%s",aStr,bStr)
}
func NewAPI()API{
	return  &ABMdel{
		a: NewAModel(),
		b: NewBMode(),
	}
}