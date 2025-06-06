package main

import "fmt"

type User struct {
	Name  string
	age   int
	email string
}

func main() {

	v := User{"abhishek", 22, "sewdefw@edfsgdf.com"}

	fmt.Println(v)
	fmt.Println(v.age)

}