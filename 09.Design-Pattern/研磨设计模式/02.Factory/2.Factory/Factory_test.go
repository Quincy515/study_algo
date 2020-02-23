package Factory

import (
	"fmt"
	"testing"
)

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetLeft(a)
	op.SetRight(b)
	return op.Result()
}

func TestFactory(t *testing.T) {
	var fac OperatorFactory

	fac = PlusOperatorFactory{}
	if compute(fac, 20, 10) != 30 {
		fmt.Println("error with factory method pattern")
	}

	fac = SubOperatorFactory{}
	if compute(fac, 20, 10) != 10 {
		fmt.Println("error with factory method pattern")
	}
}
