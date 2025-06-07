package main

import "fmt"


     func counter (cnt int ) func() int{

		 return  func() int {
                cnt++
			  return cnt ;
		 }
	 } 

  func main (){

	  inc := counter(0);
	  var number int  =inc();
	   fmt.Println(number);
	  
  }