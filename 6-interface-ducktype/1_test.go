package pit

import fmt "fmt"
import "testing"
// 不需要继承接口，因此很省事
type Zero struct{

}
type One struct{}
func (z Zero)Value() int{
	return 0 
}
type Num interface{
	Value() int
}
func (z One)Value()int{
	return 1 
}

func Test_1(t *testing.T){
	var a [2]Num
	a[0] = One{}
	a[1] = Zero{}
	// for _,obj := range a{
	// 	fmt.Println(obj.Value())
	// }
	if a[0].Value() != 1 || a[1].Value() !=0 {
		t.Error()
	}
	if false {
		fmt.Println()	
	}
}
