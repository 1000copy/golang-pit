package pit

// 何为代码的自解释？

import (
"testing"
)

func inc(a *int)int{
  (*a)++;
  return *a
}
func Test_value_to_pointer(t *testing.T){
	var a int;
	a  = 1
	inc(&a)
	if a != 2 {
		t.Error("inc expected error")
	}
}

func Test_pointer(t *testing.T){
	b := 1
	a  := &b
	inc(a)
	if *a != 2 {
		t.Error("inc expected result",*a)
	}
}
