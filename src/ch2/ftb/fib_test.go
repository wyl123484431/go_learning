package ftb

import (
	"testing"
	"fmt"
)

func TestFibList(t *testing.T) {
	var (
		a = 1
		b = 1
	)
	fmt.Print(a, " ")
	for i := 0; i < 5; i++ {
		fmt.Println(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}



func TestExchange(t *testing.T){
	a :=1
	b :=2
	tmp :=a
	a=b
	b=tmp
	t.Log(a,b)
}
