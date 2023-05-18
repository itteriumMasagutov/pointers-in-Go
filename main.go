package main

import (
	"fmt"
)

func main() {

	//Creating a pointer
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)

	//dereferencing a pointer
	y := 125
	f := &y
	fmt.Println("address of f is", f)
	fmt.Println("equal y is", *f)
	*f++
	fmt.Println("new value of i is", *f)

	//Zero value of a pointer is nil
	d := 12
	var m *int
	m = &d
	if m == nil {
		fmt.Println("m equal nil")
	} else if m != nil{
		fmt.Println("m not equal nil")
	}
	//pointer to an array as a argument to a function
	q := [3]string{"Almaz", "Dima", "Oleg"}
	modify1(&q)
	fmt.Println("pointer to array", q)

	o := [3]string{"Boris", "Dima", "Olga"}
	modify2(&o)
	fmt.Println("pointer to array", o)

	//slice as an argument to a function
	i := [3]int{10, 4, 7}
	modify3(i[:])
	fmt.Println("pointer to slice", i)

	//passing pointer to a function
	u := 60
	fmt.Println("value of k before function call is", u)
	k := &u
	change(k)
	fmt.Println("value of k after function call is", u) // or just *k
}

func modify1(arr *[3]string) {
(*arr)[1] = "Mike"
}

func modify2(arr *[3]string) {
	arr[0] = "Sasha"
}

func modify3(arr []int) {
	arr[0] = 12
}

func change(val *int) {
	*val = 35
}