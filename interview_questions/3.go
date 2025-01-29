package main

import "fmt"

type Shape interface {
	Area()
}

type Rect struct {
	len, wid float64
}

type Circle struct {
	rad float64
}

func (c Circle) Area() {
	fmt.Println(3.14 * c.rad * c.rad)
}

func (r Rect) Area() {
	fmt.Println(r.len * r.wid)
}

func CalcParm(s Shape) {
	c, ok := s.(Circle)
	if ok {
		c.Area()
	}
	r, ok := s.(Rect)
	if ok {
		r.Area()
	}

	switch sh := s.(type) {
	case Rect:
		sh.Area()
	case Circle:
		sh.Area()
	}
}

func InterfaceImpl() {

	sh := []Shape{
		Rect{4, 5},
		Circle{6},
	}

	for _, s := range sh {
		s.Area()
		CalcParm(s)
	}

}

// func main() {
// 	InterfaceImpl()
// }
