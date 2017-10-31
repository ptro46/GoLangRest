package main

import "net/http"
import (
	"fmt"
	"github.com/fromkeith/gorest" // go get github.com/fromkeith/gorest
)

func main() {

	magasinService := new(MagasinService)
	gorest.RegisterService(magasinService)

	http.Handle("/", gorest.Handle())
	http.ListenAndServe(":8080", nil)
	fmt.Print("Finished")
}
