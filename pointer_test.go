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

func Test_pointer_malloc(test *testing.T){
	// declare var like this :
	var a *int
	a = new (int)
	*a  = 1 // panic silent
	if *a != 1 {
		test.Error("1 expected")
	}
	// or like this:
	b:= new(int)
	*b = 2
	if *b != 2 {
		test.Error("2 expected")	
	}
}