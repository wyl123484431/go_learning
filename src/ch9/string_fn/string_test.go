package string_fn

import (
	"testing"
	"strings"
	"strconv"
)

func TestStringFn(t *testing.T)  {
	s:= "A,B,C"
	parts := strings.Split(s,",")   // 字符串的切割
	for _,part:=range parts{
		t.Log(part)
	}

	t.Log(strings.Join(parts,"-"))   //字符串的连接
}

func TestConv(t *testing.T)  {
	s := strconv.Itoa(10)   //字符串的转换
	t.Log("str :" + s)
	if i,err:=strconv.Atoi("10");
	 err == nil {
		t.Log(10 + i)
	}

}
