package encapsulation

import "testing"

type Employee struct {  // 结构体
	Id string
	Name string
	Age int
}

func TestCreateEmployeeObj(t *testing.T)  {
	e := Employee{"0","Bob",20}
	e1 := Employee{Name:"Mike",Age:30}
	e2 := new(Employee) //返回指针
	e2.Id="2"
	e2.Name="Rose"
	e2.Age=22
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)   // &{2 Rose 22} *指针类型
	t.Logf("e is %T",e)   // %T 是类型   &e   为指针类型
	t.Logf("e2 is %T",e2)  //e2 is *encapsulation.Employee   *指针类型

}
