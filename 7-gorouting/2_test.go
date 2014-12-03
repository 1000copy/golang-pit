package pit

import fmt "fmt"
import "testing"
import "runtime"
func sum(a[]int,c chan int){
	sum := 0 
	for _,v := range a{
		sum+=v 
	}
	c <- sum
}

func Test_1(t *testing.T){
	fmt.Println(runtime.NumGoroutine())
	a := []int{1,2,3,4,5,6}
	c := make(chan int)
	m := len(a)/2
	go sum(a[:m],c)
	go sum(a[m:],c)
	fmt.Println(runtime.NumGoroutine())
	x,y := <-c,<-c // 两个channel取完才会继续。否则就租塞于此
	if x != 6 || y != 15 {
		t.Error()
	}
}

func Test_2(t *testing.T){
	fmt.Println(runtime.NumGoroutine())
	a := []int{1,2,3,4,5,6}
	c := make(chan int)
	m := len(a)/2
	
	go sum(a[:m],c)
	go sum(a[m:],c)
	fmt.Println(runtime.NumGoroutine())
	x,y := <- c ,<- c
	if (x+y)!= 21{
		t.Error()
	}
	
	fmt.Println("")
}