package main

import "fmt"

func main() {

	w := Wheel{3}
	w.Log()

	c := Car{Wheel: w, Size: 232}
	c.Log()
}

type Car struct {
	Wheel
	Size int
}

type Wheel struct {
	Size int
}

func (w Wheel) Log() {
	fmt.Println(w.Size)
}

func (c Car) Log() {
	fmt.Println(c.Size)
}
