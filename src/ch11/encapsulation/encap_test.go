package encapsulation

import (
	"testing"
	"fmt"
	"unsafe"
)

type Employee struct {  // 结构体  封装数据  struct 实例
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


func (e Employee) string() string  {  // 在实例被调用的时，实例成员会进行复制
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d",e.Id,e.Name,e.Age)
}


/*func (e *Employee) string() string  {  // 避免内存拷贝  指针  避免开销
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d",e.Id,e.Name,e.Age)
}*/

func TestStructOperations(t *testing.T)  {
	e := &Employee{"0","Bob",20}
	fmt.Printf("Address is %x\n",unsafe.Pointer(&e.Name))    // 打印地址值
	t.Log(e.string())
}