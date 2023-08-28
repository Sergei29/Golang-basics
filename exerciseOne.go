package main

import "fmt"

type Triangle struct {
	height float64
	base   float64
}

type Square struct {
	sideLength float64
}

type Shape interface {
	getArea() float64
}

func exerciseOne() {
	s := Square{sideLength: 5.5}
	t := Triangle{base: 10, height: 3}

	printArea(s)
	printArea(t)

}

func printArea(sh Shape) {
	fmt.Println(sh.getArea())
}

func (tr Triangle) getArea() float64 {
	return 0.5 * tr.base * tr.height
}

func (sq Square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}
