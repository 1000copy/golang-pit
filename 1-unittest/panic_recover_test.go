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

func inplace_ref_fmt(){
	fmt.Print()
}

