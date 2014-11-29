unit test in go 
-------------

source code here : [golang_test.go](golang_test.go) 

run unit test by type command in terminal:

	$ cd src/github.com/1000copy/golang-pit/
	$ go test 
	PASS
	ok  	github.com/1000copy/golang-pit	0.091s

or -v to see verbose log :

	$ go test -v
	=== RUN Test_pass
	--- PASS: Test_pass (0.00s)
		golang-pit_test.go:11: pass
	PASS
	ok  	github.com/1000copy/golang-pit	0.036s

or unit test the whole package :

	$ pwd
	/Users/chuanjun/golang
	$ go test github.com/1000copy/golang-pit
	ok  	github.com/1000copy/golang-pit	0.022s