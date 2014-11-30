package pit

import (
"testing"
"fmt"
"reflect"
)
func Test_pointer_runtime_reflect(test *testing.T){
	// defer recover
	defer func(){
		if x:= recover();x !=nil{
			v := reflect.ValueOf(x)
			t := reflect.TypeOf(x)
			if (t.String() !="runtime.errorString" )||
			   (v.String() != "invalid memory address or nil pointer dereference") {
				test.Error("errorString expected")
			}
			
		}
	}()
	var a *int
	*a  = 1 // panic 
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
func inplace_ref_fmt(){
	fmt.Print()
}

