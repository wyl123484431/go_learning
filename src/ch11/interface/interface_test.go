package _interface

import "testing"

type Programmer interface {    // 接口
	WriteHellowWorld() string
}

type GoProgrammer struct {     // 实现

}

func (g *GoProgrammer) WriteHellowWorld() string{    // 鸭子模式
	return "fmt.Println(\"Hellow World\")"
}

func TestClient(t *testing.T)  {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHellowWorld())
}

