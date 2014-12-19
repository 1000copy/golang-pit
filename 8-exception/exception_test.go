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
//今天测试了下，发现尽管golang作者和社区一再啰嗦，其实 panic recover
//可以实现try catch finally。只不过，标准库不这么干，因此要全盘使用
//exception handle模式，需要包装标准库再使用。

//  类似C的错误处理代码
// （check，check，check style 的）,吓尿了
// http://weibo.com/1644105187/Bz9fQjKv6
// actually ...

// panic/recover 就是try catch finally !

