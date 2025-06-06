package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("input")
  reader:= bufio.NewReader(os.Stdin)

  fmt.Println("enter the anything")
  //comma ok //erorr ok syntax

  input,_:= reader.ReadString('\n');
  fmt.Println(input);

 //convert string to int
 
  numrating,err:= strconv.ParseFloat(strings.TrimSpace(input),64)

   if err!=nil{
	 fmt.Println(err)
   }else{
	 fmt.Println("added 1 to your rating:",numrating+1)
   }



    


}
