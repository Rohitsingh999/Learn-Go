package main

import "fmt"

  
 func rec( runes []rune, idx int,mid int)  {


	  if idx==mid {
		return 
	  }


        n :=len(runes)

		runes[idx],runes[n-1-idx] =runes[n-1-idx], runes[idx]

		 rec(runes,idx+1,mid);
	   
	

 }
   

 func main (){

    
	  str := "Rohit"

     runes:= []rune(str)
     mid:= len(str)/2;

       fmt.Println(str);
	   rec(runes,0,mid);
       str =string(runes)
	   fmt.Println(str);
     
 }