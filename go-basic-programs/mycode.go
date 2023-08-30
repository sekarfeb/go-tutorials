package main

import ("fmt" ; "reflect")

func main(){
	fmt.Println("my code")
	var a=10
	b := 20.0004
	for i := 1 ; i <= 10; i++ {

	   fmt.Println(i)
	   fmt.Println(a,b)

	    fmt.Println("Type of a ",reflect.TypeOf(a))
	    fmt.Println("Type of b",reflect.TypeOf(b))

 	}
}
