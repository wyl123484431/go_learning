package _map

import "testing"

// map的value值 是个方法
func TestMapWithFunValue(t *testing.T)  {
	m := map[int]func(op int)int{}
	m[1] =func(op int) int {return op}
	m[2] =func(op int) int {return op*op}
	m[3] =func(op int) int {return op*op*op}

	t.Log(m[1](2),m[2](2),m[3](2)) // 2 4 8
	t.Log(m)
}

// 使用map 来实现 set    go语言没有内置的set
func TestMapForSet(t *testing.T)  {
	mySet := map[int]bool{}
	mySet[1] = true

	n :=1
	if mySet[n] {     //没有 添加
		t.Log("n is existing",n)
	}else {
		t.Logf("%d is not existing",n)
	}

	mySet[3] = true
	t.Log(len(mySet))

	delete(mySet,1)   //删除map中 值
	n =1
	if mySet[n] {
		t.Log("n is existing",n)
	}else {
		t.Logf("%d is not existing",n)
	}
}
