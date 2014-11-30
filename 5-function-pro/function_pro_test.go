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
//  有人说：但是。。因为go不支持泛型，因此，对map/reduce的表现是不完整的。
//  那么，interface{},这样的万用类型，是否可以解决此问题？
//  结果：虽然有点丑，但是确实如果设想的那样，可以做。
// map interface{}

func map_(array []interface{}, fun func(interface{})interface{}){
    for index, element := range array{
        array[index] = fun(element)
    }
}
func Test_m4(t*testing.T){
    v := []int{1, 2 ,3,6}
    vv := make([]interface{},4)
    for i := range v{
    	vv[i] = v[i]
    }
    fmt.Println(v)
    map_(vv, 
    	// func(s interface{})interface{}{return int(s) *int(s) })
    	func(s interface{})interface{}{return s.(int) *s.(int) })
    fmt.Println(vv)
}

// atom simplize 

type T interface{}

func map_a(array []T, fun func(T)T){
    for index, element := range array{
        array[index] = fun(element)
    }
}
func Test_m5(t*testing.T){
    v := []int{1, 2 ,3,9}
    vv := make([]T,4)
    for i := range v{
    	vv[i] = v[i]
    }
    fmt.Println(v)
    map_a(vv, 
    	// func(s interface{})interface{}{return int(s) *int(s) })
    	func(s T)T{return s.(int) *s.(int) })
    fmt.Println(vv)
}
// 除了[]int ,还可以[]string . 有点泛型的味道。

func Test_m6(t*testing.T){
    v := []string{"1", "2" ,"3","9"}
    vv := make([]T,4)
    for i := range v{
    	vv[i] = v[i]
    }
    fmt.Println(v)
    map_a(vv, 
    	// func(s interface{})interface{}{return int(s) *int(s) })
    	func(s T)T{return s.(string) +s.(string) })
    fmt.Println(vv)
}
// func t_array(v []T)[]T{
// 	r := make([]T,len(v))
//     for i := range v{
//     	r[i] = v[i]
//     }
//     return r
// }
func Test_m7(t*testing.T){
    v := []string{"1", "2" ,"3","1"}
    // vv := make([]T,len(v))
    vv := []T{}
    for _,vvv:=range v {
	    vv = append(vv,vvv)
	}
    fmt.Println(v)
    map_a(vv, 
    	// func(s interface{})interface{}{return int(s) *int(s) })
    	func(s T)T{return s.(string) +s.(string) })
    fmt.Println(vv)
}

// func Test_m8(t*testing.T){
//     v := []string{"1", "2" ,"3","1"}
//     // vv := make([]T,len(v))
//     vv := []T{}
// 	vv = append(vv,v...)
//     fmt.Println(v)
//     map_a(vv, 
//     	// func(s interface{})interface{}{return int(s) *int(s) })
//     	func(s T)T{return s.(string) +s.(string) })
//     fmt.Println(vv)
// }



