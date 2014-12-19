golang-pit
==========

golang features code and doc

21-centery clang ,big 8 suprises ... 

我被蛊惑，我要开坑。 

How to startup leanly ?
-------------
- clone the repository
- go test 

configuration enviroments for users
-------------

Crawle to github.com and new repository 1000copy/golang-pit . Then back to my mac ,type command in terminal:

	$ git config --global user.name "YOUR USERNAME"
	$ git config --global user.email YOUREMAIL

	$ mkdir golang
	$ mkdir golang/src
	$ mkdir golang/pkg
	$ mkdir golang/bin
	$ export GOPATH=GOLANG_ABSOLUTLY_DIRICTORY
	$ echo $GOPATH
	$ go get github.com/1000copy/golang-pit

configuration enviroments for developers
-------------
	$mkdir golang
	$mkdir golang/src
	$cd golang/src
	$GIT TO HERE
	$ export GOPATH=GOLANG_ABSOLUTLY_DIRICTORY
	$ go test ./...

Golang tests in sub-directory
----------
simple answer is :
	go test ./...

	3个点！
	如果是go  test  package_name ,就只会执行当前包的第一层目录

google: golang test in sub directory insite:stackoverflow.com

