package main

import "fmt"


 func main() {
//    fmt.Println("pointers");

//    var ptr *int 

//    fmt.Println(ptr)

 number := 23

   var ptr  = &number
    fmt.Println(*ptr);
	fmt.Println("address",ptr);

	*ptr= *ptr*2
	 fmt.Println(*ptr);
	 fmt.Println(number);
	 fmt.Println("address",ptr);

 }