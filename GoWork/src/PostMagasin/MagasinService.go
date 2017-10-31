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
	postMagasin    gorest.EndPoint `method:"POST" path:"/" postdata:"Magasin" output:"Magasin"`
}

func (serv MagasinService) PostMagasin(m Magasin) (mag Magasin) {
	serv.ResponseBuilder().AddHeader("Access-Control-Allow-Origin", "*")

	fmt.Printf("AddMagasin\n")

	if len(m.Nom) == 0 {
		serv.ResponseBuilder().
			SetResponseCode(400).
			WriteAndOveride([]byte("{ \"message\" : \"Il faut indiquer le nom du magasin dans le json en entree\" }"))
	} else {
		mag := createMagasin(m.Nom)
		if nil != mag {
			serv.ResponseBuilder().SetResponseCode(200)
			return *mag
		}
	}
	return
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
