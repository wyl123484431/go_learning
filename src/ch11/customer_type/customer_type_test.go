package customer_type

import (
	"time"
	"fmt"
	"testing"
)

type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv {        // 计算执行时间
	return func(n int) int{
		start :=time.Now()
		ret:= inner(n)
		fmt.Println("time apent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {                // 执行方法睡眠一秒
	time.Sleep(time.Second*1)
	return op
}

func TestFn(t *testing.T)  {
	tsSf := timeSpent(slowFun)
	t.Log(tsSf(10))
}
