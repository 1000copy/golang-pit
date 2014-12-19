package pit

import (
  "testing"
  "fmt"
)
// try catch 底层抛出，高层捕获
func callstack_level_1 (){
  defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
   callstack_level_2 ()
}
func callstack_level_2 (){
	callstack_level_3()
}
func callstack_level_3 (){
	panic("low level error") 
}
// try finally 在level_0 内不管发生了什么，finally section 都会执行
func callstack_level_0 (){
	panic("panic directly") ;
}
// try finally
func Test_tryfinally(t *testing.T){
	// finally section
	defer func() {
        if r := recover(); r != nil {
            fmt.Println("finally section", r)
        }
    }()
    //do something
    callstack_level_0()
}
// try catch
func Test_trycatch(t *testing.T){
    callstack_level_1()
}


