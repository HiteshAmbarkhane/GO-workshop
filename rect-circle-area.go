package main

import "fmt"
import "strconv"
import "os"

type Rectangle struct {
	length, width float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 { return r.length * r.width }
func (c Circle) Area() float64    { return 3.14 * c.radius * c.radius }

func (r Rectangle) Perimeter() float64 { return 2 * (r.length + r.width) }
func (c Circle) Perimeter() float64    { return 3.14 * c.radius * 2 }

func main() {
	a, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	if len(os.Args) > 2 {
		b, err := strconv.ParseFloat(os.Args[2], 64)
		if err != nil {
			fmt.Printf("couldn't convert number: %v\n", err)
			return
		}
		fmt.Printf("Its a Rectangle with a = %v and b = %v\n", a, b)
		r := Rectangle{a, b}
		fmt.Printf("Rect area = %v perimeter = %v\n", r.Area(), r.Perimeter())
	} else {
		fmt.Printf("Its a Circle with radius = %v\n", a)
		c := Circle{a}
		fmt.Printf("Circle area = %v perimeter = %v\n", c.Area(), c.Perimeter())
	}
}
