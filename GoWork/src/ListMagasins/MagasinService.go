package main

import (
	"fmt"
	"github.com/fromkeith/gorest" // go get github.com/fromkeith/gorest
)

//************************Define Service***************************
type MagasinService struct {
	gorest.RestService `root:"/StoreWS/api/magasin/" consumes:"application/json" produces:"application/json"`

	magasinList gorest.EndPoint `method:"GET" path:"/" output:"[]Magasin"`
}

func (serv MagasinService) MagasinList() []Magasin {
	serv.ResponseBuilder().AddHeader("Access-Control-Allow-Origin", "*")

	fmt.Printf("MagasinList\n")

	serv.ResponseBuilder().SetResponseCode(200)
	return loadMagasins()
}
