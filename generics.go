package main

import "fmt"


 type  number interface {

	 int | float32
 }


    func multiple [T number] ( a T) T{

		 return  a*a;
	}
	 
 func main (){


       fmt.Println(multiple(6))
	 
 }