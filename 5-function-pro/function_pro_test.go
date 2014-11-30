package pit

import fmt "fmt"
import "testing"

type Stringy func() string

func foo() string{
  return "Stringy function"
}

func takesAFunction(foo Stringy){
  fmt.Printf("takesAFunction: %v\n", foo())
}

func returnsAFunction()Stringy{
  return func()string{
    fmt.Printf("Inner stringy function\n");
    return "bar" // have to return a string to be stringy
  }
}

func Test_main(t*testing.T){
  takesAFunction(foo);
  var f Stringy = returnsAFunction();
  f();
  var baz Stringy = func()string{
    return "anonymous stringy\n"
  };
  fmt.Printf(baz());
}

// map

func map_func(array []int, fun func(int)int){
    for index, element := range array{
        array[index] = fun(element)
    }
}
func Test_m1(t*testing.T){
    v := []int{1, 2 ,3 ,4 , 5}
    fmt.Println(v)
    map_func(v, func(s int)int{return s * s})
    fmt.Println(v)
}


// filter

func filter_func(array []int, fun func(int)bool)[]int{
    rlist := []int{}
    for index, element := range array{
        if fun(element){
          rlist = append(rlist,array[index] )
        }
    }
    return rlist
}
// filter + foo,语意是整体的，一次写在一起。
// 这是delphi等支持函数变量，却不支持匿名函数的语言的差别。
// delphi是无法把apply +foo书写在一起的
func Test_m2(t*testing.T){
    v := []int{1, 2 ,3 ,4 , 5}
    fmt.Println(v)
    r := filter_func(v, func(s int)bool{return s > 1})
    fmt.Println(r)
}
// lambda express : referece local var
func Test_m3(t*testing.T){
    v := []int{1, 2 ,3 ,4 , 5}
    fmt.Println(v)
    from :=  4;
    r := filter_func(v, func(s int)bool{return s > from})
    fmt.Println(r)
}
//  但是。。因为go不支持泛型，因此，对map/reduce的表现是不完整的。


