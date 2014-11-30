package self_expression

// 何为代码的自解释？

import (
"fmt"
"testing"
)

func inc(a *int)int{
  (*a)++;
  return *a
}
//  执行代码，以代码直接表达期望
func Test_self_express(t *testing.T){
	var a int;
	a  = 1
	inc(&a)
	if a != 2 {
		t.Error("inc expected error")
	}
}
// 验证结果，打印方式。需要消费眼睛的注意力，运行时去了解代码是否正确，并且需要文字说明期望
func Test_eye_consume_express(t *testing.T){
	var a int;
	a  = 1
	inc(&a)
	fmt.Println(a)
}

// 云泥之别。再次体现 ！