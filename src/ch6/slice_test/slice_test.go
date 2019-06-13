package slice_test

import "testing"



//切片
func TestSliceInit(t *testing.T)  {

	var s0 []int
	t.Log(len(s0),cap(s0))

	s0 = append(s0,1)
	t.Log(len(s0),cap(s0))

	s1 := []int{1,2,3,4}
	t.Log(len(s1),cap(s1))

	s2 :=make([]int,3,5) // 3是可访问长度  5是容量
	t.Log(len(s2),cap(s2))   // 3 5
	t.Log(s2[0],s2[1],s2[2]) //0 0 0 
	
	s2 = append(s2,1)
	
	t.Log(s2[0],s2[1],s2[2],s2[3])//0 0 0 1
	t.Log(len(s2),cap(s2)) // 4,5

}

func TestSliceGrowing(t *testing.T)  {
	s :=[]int{}
	for i:=0;i<10;i++{
		s = append(s,i)
		/*
			slice_test.go:30: 1 1
			slice_test.go:30: 2 2
			slice_test.go:30: 3 4
			slice_test.go:30: 4 4
			slice_test.go:30: 5 8
			slice_test.go:30: 6 8
			slice_test.go:30: 7 8
			slice_test.go:30: 8 8
			slice_test.go:30: 9 16
			slice_test.go:30: 10 16
		*/
		t.Log(len(s),cap(s))
	}
}

// 切片
func TestSliceShareMemory(t *testing.T)  {
	year :=[]string{"Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sep","Oct","Nov","Dec"}

	Q2 :=year[3:6]
	t.Log(Q2,len(Q2),cap(Q2)) //[Apr May Jun]   3 9

	summer :=year[5:8]
	t.Log(summer, len(summer), cap(summer)) // [Jun Jul Aug] 3 7

	summer[0] = "Unkonw"
	t.Log(Q2) //[Apr May Unkonw]
	t.Log(year) //[Jan Feb Mar Apr May Unkonw Jul Aug Sep Oct Nov Dec]
}

func TestSliceComparing(t *testing.T)  {
	a := []int{1,2,3,4}
	b := []int{1,2,3,4}
	t.Log(a,b)
	/*if a==b { //切片不可以比较
		t.Log("equal")
	}*/

}