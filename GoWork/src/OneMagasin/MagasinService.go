package main

import (
	"fmt"
	"github.com/fromkeith/gorest" // go get github.com/fromkeith/gorest
)

//************************Define Service***************************
type MagasinService struct {
	gorest.RestService `root:"/StoreWS/api/magasin/" consumes:"application/json" produces:"application/json"`

	magasinList    gorest.EndPoint `method:"GET" path:"/" output:"[]Magasin"`
	magasinDetails gorest.EndPoint `method:"GET" path:"/{Id:int64}" output:"Magasin"`
}

func (serv MagasinService) MagasinList() []Magasin {
	serv.ResponseBuilder().AddHeader("Access-Control-Allow-Origin", "*")

	fmt.Printf("MagasinList\n")

	serv.ResponseBuilder().SetResponseCode(200)
	return loadMagasins()
}

func (serv MagasinService) MagasinDetails(Id int64) (m Magasin) {
	serv.ResponseBuilder().AddHeader("Access-Control-Allow-Origin", "*")

	fmt.Printf("MagasinDetails parameter %d\n", Id)
	magasin := loadMagasin(Id)
	if nil != magasin {
		serv.ResponseBuilder().SetResponseCode(200)
		return *magasin
	} else {
		serv.ResponseBuilder().
			SetResponseCode(500).
			WriteAndOveride([]byte("{ \"message\" : \"Le Magasin indique n'existe pas\" }"))
	}

	return
}
