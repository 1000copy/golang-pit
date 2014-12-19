golang-pit
==========

golang features code and doc

21-centery clang ,big 8 suprises ... 

我被蛊惑，我要开坑。 

How to startup leanly ?
-------------
- clone the repository
- go test 

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

Golang tests in sub-directory
----------
simple answer is :
	go test ./...

	3个点！
	如果是go  test  package_name ,就只会执行当前包的第一层目录

Golang - Run all go tests recursively including subpackages
In order to run all tests of your go app including sub packages just perform the following command:

go test ./…

Shawshank:golang-pit rita$ go test ./...  
ok  	golang-pit	0.006s
ok  	golang-pit/1-unittest	0.009s
ok  	golang-pit/2-slice	0.016s
ok  	golang-pit/3-simple-type_test	0.025s
ok  	golang-pit/4-object-oriented	0.008s
ok  	golang-pit/5-function-pro	0.011s
ok  	golang-pit/6-interface-ducktype	0.005s
ok  	golang-pit/7-gorouting	0.021s
ok  	golang-pit/8-exception	0.005s



http://stackoverflow.com/questions/19200235/golang-tests-in-sub-directory

