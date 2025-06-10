package main

import (
	"fmt"
	"sync"
)


   func print(i int, wg *sync.WaitGroup){

	     defer wg.Done()
	  fmt.Println(i);
   }

  func main (){

	   var wg sync.WaitGroup

      for i:=1;i<=10;i++ {


		  wg.Add(1)

		 go  print(i,&wg)
	  }
	 
	   wg.Wait()
  }