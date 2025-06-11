package main

import (
	"CRUD/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("starting server at port 8000 ")

	log.Fatal(http.ListenAndServe(":8000", routes.CreateRoutes()))

}
