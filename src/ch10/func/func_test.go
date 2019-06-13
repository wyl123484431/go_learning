package _func

import (
	"math/rand"
	"testing"
	"time"
	"fmt"
)

func returnMultiValues()(int,int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func (op int)int {        // 计算执行时间
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
	a, _ := returnMultiValues()
	t.Log(a)  // 1 7
	tsSf := timeSpent(slowFun)
	t.Log(tsSf(10))
}

func Sum(ops...int) int {    //可变参数    转化为数组
	ret :=0
	for _, op:= range ops{
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T)  {
	t.Log(Sum(1,2,3,4,5))  //15
}

func Clear()  {
	fmt.Println("Clear resources.")

}

// 延迟执行
func TestDefer(t *testing.T)  {
	defer Clear()    // 相当于  finally
	fmt.Println("Start")
	panic("err")   // 报错了   还是会执行
}






