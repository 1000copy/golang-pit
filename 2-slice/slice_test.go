package slice_test

import (
"testing" 
// "fmt"
"strconv"
)
func Test_valuecopy(t *testing.T){
	var a [10]int = [10]int{1, 2, 3}
    b := a
    b[0] = 2
    if (a[0] == b[0]){
    	t.Error("value copy error")
    }
}
func Test_slicecreate(t *testing.T){
	var a [10]int = [10]int{1, 2, 3,4,5}
	// create method 1
    b := a[0:3]
    // create method 2
    c:= make([]int, 10, 20)  	
    if len(b) != 3 {
    	t.Error("length error")
    }
    if cap(c) != 20 || len(c) !=10 {
    	t.Error("length or len error")
    }
}
func Test_cap(t*testing.T){
	data := make([]int, 10, 20)
    data[0] = 1
    data[1] = 2
    if len(data) != 10{
    	t.Error("len error")
    }
    if cap(data) != 20{
    	t.Error("cap error")
    }
    // t.Log("length:", len(data), "cap:", cap(data), ":", data)
    // t.Logf("length:%d cap:%d data:%s",len(data),cap(data),to_str(data[:]))
}
func to_str(a []int) string{
	r:=""
	for i := range a {
		// r = r+ string(a[i])
		if i == 0 {
			r += strconv.Itoa(a[i])
		}else{
			r += ","+strconv.Itoa(a[i])
		}
	}
	return "["+r+"]"
}
func Test_slice(t *testing.T){
	slice :=[]byte{1,2,3,4,5,6}
	a := slice[2:5]	
	if len(a) != 3 {
		t.Error("a len expected 3")
	}
	if a[0] != 3 || a[1]!=4  || a[2] != 5 {
		t.Error("index 2,3,4 value error")
	}
	if !testEq(a ,[]byte {3,4,5}) {
		t.Error("index 2,3,4 array content error")
	}
	// a[0] = 254
	a[0] = 254
	if (slice[2] != 254){
		t.Error("ref expected!")
	}
	c := make([]byte,len(a))
	copy(a,c)
	c[0] = 254
	if (slice[2] == 254){
		t.Error("value ,value expected .")
	}
	if testEq(slice[:],[]byte{1,2,3,4,5,6}){
		t.Error("all slice is error")
	}
    if testEq(slice[:3],[]byte{1,2,3}){
		t.Error("slice from 0-3 ")
	}
	if testEq(slice[3:],[]byte{4,5,6}){
		t.Error("slice from 3")
	}
	if testEq(append(a,7,8),[]byte{1,2,3,4,5,6,7,8}){
		t.Error("append failure")
	}
}
// more ref : 
// http://www.cnblogs.com/yjf512/archive/2012/06/14/2549929.html
// http://studygolang.com/topics/74
func testEq(a, b []byte) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}