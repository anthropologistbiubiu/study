package strategy

import "testing"

func Test_strategy(t *testing.T) {
	c := NewContext()
	plantA := NewStrategyA()
	c.strategy = plantA
	c.Execute()

	plantB := NewStrategyB()
	c.strategy = plantB
	c.Execute()
}
