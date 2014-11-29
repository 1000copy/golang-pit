golang-pit
==========

golang features code and doc

21-centery clang ,big 8 suprises ... 

我被蛊惑，我要开坑。以后如何，不管了。/////\\\\\\\\\


configuration enviroments
-------------

Crawle to github.com and new repository 1000copy/golang-pit . Then back to my mac ,type command in terminal:

	$ git config --global user.name "1000copy"
	$ git config --global user.email 1000copy@gmail.com

	$ pwd
	/Users/chuanjun
	$ mkdir golang
	$ mkdir golang/src
	$ mkdir golang/pkg
	$ mkdir golang/bin
	$ export GOPATH=/Users/chuanjun/golang
	$ echo $GOPATH
	/Users/chuanjun/golang
	$ go get github.com/1000copy/golang-pit


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