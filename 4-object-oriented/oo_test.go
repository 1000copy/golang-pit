package pit
import (
"testing"
)

type Rect struct {
	width ,height float64
}
// func 
func area (r Rect)float64{
	return r.width*r.height;
}

func Test_area_func(t *testing.T){
	r := Rect{10,5}
	if 50.0!= area(r){
		t.Error("area calculate error")
	}
}

// or encapulating 封装

func (r Rect)area()float64{
	return r.width*r.height
}

// * or not ,it is just ok

func (r *Rect)area_p()float64{
	return r.width*r.height
}

func Test_area_method(t *testing.T){
	r := Rect{10,5}
	if 50.0!= r.area(){
		t.Error("area calculate error")
	}
	if 50.0!= r.area_p(){
		t.Error("area calculate error")
	}
}
// * , change object state

func (r Rect)Set_width (newWidth float64){
	r.width = newWidth
}


func (r *Rect)Set_width_p(newWidth float64){
	r.width = newWidth
}

func Test_area_change_property(t *testing.T){
	r := Rect{10,5}
	if 10!= r.width{
		t.Error()
	}
	r.Set_width(11)
	if 10!= r.width{
		t.Error()
	}
	r.Set_width_p(11)
	if 11!= r.width{
		t.Error()
	}
}

// 继承 + 重载
type Shape struct{}
type Circle struct{
	radius float64
	Shape
}
type Rectangle struct{
	width,height float64
	Shape
}

func (s Shape )area()float64{
  return 0
}
func (s Circle )area()float64{
  pi := 3.14
  return pi*s.radius*s.radius
}
func (s Rectangle)area()float64{
  return s.width*s.height
}

func Test_inherited (t *testing.T){
	s := Shape{}
	if 0 != s.area(){
		t.Error(s.area())
	}
	
	r := Rectangle {10.0,5.0,Shape{}}
	c := Circle{10.0,Shape{}}
	if 50 != r.area(){
		t.Error(r.area())
	}
	if 314 != c.area(){
		t.Error(c.area())
	}
}
//  simple type extension method

type Vchcode string
//  object slices




