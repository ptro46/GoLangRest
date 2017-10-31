package main

import (
	"github.com/fromkeith/gorest" // go get github.com/fromkeith/gorest
)

//************************Define Service***************************
type HelloService struct {
	gorest.RestService `root:"/first/"`

	helloWorld gorest.EndPoint `method:"GET" path:"/hello/" output:"string"`
}

func (serv HelloService) HelloWorld() string {
	serv.ResponseBuilder().SetResponseCode(200)
	return "GoLang Hello World"
}
