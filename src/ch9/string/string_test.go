package string

import "testing"

func TestSting(t *testing.T)  {
	var s string    // 是个字节的切片
	t.Log(s)  // 默认值为 ""
	s ="hellow"
	t.Log(len(s))   // 6
	// s[1]="3" //string是不可改变的byte slice
	s ="\xE4\xB8\xA5" //可以存储任何二进制数据
	t.Log(s)
	s = "中"
	t.Log(len(s)) // 3     byte数
	c:=[]rune(s)     //rune  数据类型   可以转化为unicode
	t.Log(len(c)) // 1
	t.Logf("中 unicode %x", c[0])  //4e2d
	t.Logf("中 UTF8 %x", s)        //e4 b8 ad      三个字节
}
