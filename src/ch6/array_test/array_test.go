package array_test

import "testing"

func TestArrayInit(t *testing.T)  {
	var arr [3]int
	arr1 :=[4]int{1,2,3,4}
	arr3 :=[...]int{1,2,3,4,5}
	arr1[1]=5
	t.Log(arr[1],arr[2])
	t.Log(arr1,arr3)
}

func TestArrayTravel(t *testing.T)  {
	arr3 := [...]int{1,3,4,5}
	// 多维数组
	arr4 := [2][2]int{{1,2},{3,4}}
	for _,e := range arr4 {
		t.Log(e)
	}
	for i:=0; i<len(arr3); i++{
		t.Log(arr3[i])
	}
	// 站位不关注下标
	for _,e := range arr3 { //indx
		t.Log(e)
	}
}

// 数组的截断
func TestArraySection(t *testing.T){
	arr3 := [...]int{1,2,3,4,5}

	arr3_sec := arr3[3:]
	t.Log(arr3_sec)
}