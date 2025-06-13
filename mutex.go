package main

import (
	"fmt"
	"sync"
)



   type variable struct {

	 View int
	 mu sync.Mutex
   }

   func (v *variable ) inc (wg *sync.WaitGroup){
    
	 defer func (){ 
		    wg.Done() 
		    v.mu.Unlock() 
			}()
	 v.mu.Lock()
	     v.View+=1 
    
		 fmt.Println(v.View)
		 
   }


 func main (){


	  var wg sync.WaitGroup

	    v:= variable{View: 0}
	 
	   for i:=1 ;i<=100;i++{
           wg.Add(1)
		   go v.inc(&wg)
	   }

	  
      
        wg.Wait()

 }