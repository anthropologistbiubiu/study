package strategy

import "fmt"

type Strategy interface {
	Execute()
}

type strategyA struct {
}

func (s *strategyA) Execute() {
	println("A plan execute")
}

func NewStrategyA() Strategy {
	return &strategyA{}
}

type strategyB struct {
}

func (s *strategyB) Execute() {
	fmt.Println("B plant excute")
}

func NewStrategyB() Strategy {
	return &strategyB{}
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) Execute() {
	c.strategy.Execute()
}

func NewContext() *Context {
	return &Context{}
}
