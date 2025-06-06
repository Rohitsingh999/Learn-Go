package main

import "fmt"

func main() {

	// fmt.Println("arrays")

	// var array [5]string

	//  array [0]="Rohit"
	//  array [1]="Mohit"
	//  array [4]="Nohit"

	//Slice
	// var arr = [3]string{"a", "b", "c"}
	// var slc = [...]string{"d", "e", "f"}

    // fmt.Println(arr)
	// fmt.Println(slc)

	// var slc1 = append(slc[1:2],"rohit ","singh ","rawat",".")
	// fmt.Println(slc1);

	 var mp = map[int]string{

		 1: "rohit",
		 2: "Mohit",
		 3: "Pohit",
	 }

	  for key,val := range mp{

		 fmt.Println(key,"->",val);
	  }

	  
	 
	

}