package golang

import (
//  invalid import path: "testing "
//  "testing "
"testing"
)
// expected declaration, found '{'

func Test_pass(t *testing.T){
	t.Log("pass")
	// t.Error("就是这么有钱，就是这么任性！")
}
