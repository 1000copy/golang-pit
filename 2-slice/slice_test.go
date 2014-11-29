package slice_test

import (
"testing" 
// ;"fmt"
)


func Test_slice(t *testing.T){
	slice :=[]byte{1,2,3,4,5,6}
	a := slice[2:5]	
	if len(a) != 3 {
		t.Error("a len expected 3")
	}
	if a[0] != 3 || a[1]!=4  || a[2] != 5 {
		t.Error("index 2,3,4 value error")
	}
	// b := 
	if !testEq(a ,[]byte {3,4,5}) {
		t.Error("index 2,3,4 array content error")
	}
}

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