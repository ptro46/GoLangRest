package main

import "net/http"
import (
	"fmt"
	"github.com/fromkeith/gorest"
)

func main() {

	helloService := new(HelloService)
	gorest.RegisterService(helloService)

	http.Handle("/", gorest.Handle())
	http.ListenAndServe(":8181", nil)
	fmt.Print("Finished")
}
