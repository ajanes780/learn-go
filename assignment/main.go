package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct{}
type square struct{}

func main() {
	t := triangle{}
	s := square{}
	fmt.Println(t.getArea())
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * 0.5
}

func (s square) getArea() float64 {
	return 0.5 * 0.5
}
