package pit

import (
	"testing"
	"strconv"
)

func Test_convert_int_string(t *testing.T){
	a := 1
	if "1"!= strconv.Itoa(a) {
		t.Error()
	}
	s := "1"
	i, err := strconv.Atoi(s)
	if 1 != i || err !=nil{
		t.Error()
	}
}