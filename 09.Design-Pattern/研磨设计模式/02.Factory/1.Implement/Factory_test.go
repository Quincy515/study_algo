package Implement

import (
	"fmt"
	"testing"
)

func TestFactoryPlus(t *testing.T) {
	var fac OperatorFactory
	fac = PlusOperatorFactory{}
	op := fac.Create()
	op.SetLeft(20)
	op.SetRight(10)
	fmt.Println("plus result: ", op.Result())
}

func TestFactorySub(t *testing.T) {
	var fac OperatorFactory
	fac = SubOperatorFactory{}
	op := fac.Create()
	op.SetLeft(20)
	op.SetRight(10)
	fmt.Println("sub result: ", op.Result())
}
