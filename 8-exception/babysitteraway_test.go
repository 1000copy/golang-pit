package pit

import (
  "testing"
  "fmt"
  "io/ioutil"
)

func check(e error){
  if e!=nil{
    panic(e)
  }
}
func Test_babysitter(t *testing.T){
    filename := "babysitteraway_test.go"
    data,err := ioutil.ReadFile(filename)
    check(err)
    fmt.Println(string(data)[:10])
    data,err = ioutil.ReadFile(filename)
    check(err)
    fmt.Println(string(data)[:10])
    data,err = ioutil.ReadFile(filename)
    check(err)
    fmt.Println(string(data)[:10])    
}
//

type File struct {

}
func (f File)ReadFile(file string)[]byte{
  data,err := ioutil.ReadFile(file)
  check(err)
  return data
}
func Test_babysitter_keep_away(t *testing.T){
    // panic: open tmp/dat: no such file or directory [recovered]
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    filename := "babysitteraway_test.go"
    file := File {};
    data := file.ReadFile(filename)
    fmt.Println(string(data)[:10])
    data = file.ReadFile(filename)
    fmt.Println(string(data)[:10])
    data = file.ReadFile(filename)
    fmt.Println(string(data)[:10])

}








// keep babysitter programing away !