package main

import ("fmt"; "reflect")

func Var(){
	var a = "Initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var e int
	fmt.Println(e)

	var d = true
	fmt.Println(d)

	f := true
	fmt.Println(f)

	fmt.Println("Type of f variable: ",reflect.TypeOf(f))
}