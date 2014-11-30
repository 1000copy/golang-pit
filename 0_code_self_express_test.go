package pit

// 何为代码的自解释？

import (
"fmt"
"testing"
)

func add(a *int)int{
  (*a)++;
  return *a
}
//  执行代码，以代码直接表达期望
func Test_self_express(t *testing.T){
	var a int;
	a  = 1
	add(&a)
	if a != 2 {
		t.Error("add expected error")
	}
}
// 验证结果，打印方式。并且需要文字说明期望.消费眼睛的注意力,我觉得不够好
func Test_eye_consume_express(t *testing.T){
	var a int;
	a  = 1
	add(&a)
	fmt.Println(a)
}

// 云泥之别。再次体现 ！