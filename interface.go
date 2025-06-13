package main

import "fmt"

    
 type  Animal  interface { //by making Animal interface we can now make diff strct for diif animal lile lion ,tiger,etc
	 call()
 }

  
  type  Forest struct {
                              //main struct to take animal of type interface means have call method 
	 Living_Being Animal
  }

  func (b Forest) callname (){

	  b.Living_Being.call();
  }


  type Lion struct{

	 Name string 
	 gender string
	 Eat  string
  }

  func (l Lion ) call (){
	fmt.Println( l);
  }

  type Zebra struct{

	 Name string 
	 gender string
	 Eat  string
  }

  func (Z Zebra ) call (){
	fmt.Println( Z);
  }



  

 func main (){

	
    //    L:= Lion {
	// 	 Name : "Lion",
	// 	 gender: "male",
	// 	 Eat : "Meat",
	//    }
       Z:= Zebra{
		 Name : "Zebra",
		 gender: "male",
		 Eat : "Grass",
	   }

	   F:=Forest{Z}


	   F.callname()
	 
 }