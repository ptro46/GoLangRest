package main

import (
	"fmt"
	"github.com/fromkeith/gorest" // go get github.com/fromkeith/gorest
)

//************************Define Service***************************
type MagasinService struct {
	//Service level config
	gorest.RestService `root:"/StoreWS/api/magasin/" consumes:"application/json" produces:"application/json" charset:"utf-8" allowGzip:1 `

	//End-Point level configs: Field names must be the same as the corresponding method names,
	// but not-exported (starts with lowercase)
	magasinList      gorest.EndPoint `method:"GET" path:"/" output:"[]Magasin"`
	magasinDetails   gorest.EndPoint `method:"GET" path:"/{Id:int64}" output:"Magasin"`
	magasinRayonList gorest.EndPoint `method:"GET" path:"/{Id:int64}/rayons" output:"[]Rayon"`
	postMagasin      gorest.EndPoint `method:"POST" path:"/" postdata:"Magasin" output:"Magasin"`
	putMagasin       gorest.EndPoint `method:"PUT" path:"/" postdata:"Magasin" output:"Magasin"`
	deleteMagasin    gorest.EndPoint `method:"DELETE" path:"/{Id:int64}" output:"Magasin"`
}

func (serv MagasinService) init() {
}

//Handler Methods: Method names must be the same as in config, but exported (starts with uppercase)

func (serv MagasinService) MagasinList() []Magasin {
	serv.ResponseBuilder().AddHeader("Access-Control-Allow-Origin", "*") // debug only : permet les requetes cross-origin

	fmt.Printf("MagasinList\n")

	serv.ResponseBuilder().SetResponseCode(200)
	return loadMagasins(DatabaseSingleton.db)
}

func (serv MagasinService) MagasinRayonList(Id int64) []Rayon {
	serv.ResponseBuilder().AddHeader("Access-Control-Allow-Origin", "*") // debug only : permet les requetes cross-origin

	fmt.Printf("MagasinRayonList\n")

	result := make([]Rayon, 0, 0)

	magasin := loadMagasin(DatabaseSingleton.db, Id)
	if nil != magasin {
		serv.ResponseBuilder().SetResponseCode(200)
		return loadRayonsFromMagasin(DatabaseSingleton.db, Id)
	} else {
		serv.ResponseBuilder().SetResponseCode(500).WriteAndOveride([]byte("{ \"message\" : \"Le Magasin indique n'existe pas\" }"))
	}
	return result
}

func (serv MagasinService) MagasinDetails(Id int64) (m Magasin) {
	serv.ResponseBuilder().AddHeader("Access-Control-Allow-Origin", "*") // debug only : permet les requetes cross-origin

	fmt.Printf("MagasinDetails\n")

	magasin := loadMagasin(DatabaseSingleton.db, Id)
	if nil != magasin {
		serv.ResponseBuilder().SetResponseCode(200)
		return *magasin
	} else {
		serv.ResponseBuilder().SetResponseCode(500).WriteAndOveride([]byte("{ \"message\" : \"Le Magasin indique n'existe pas\" }"))
	}
	return
}

// receive json[{}] --> idt:[0] nom:[]
func (serv MagasinService) PostMagasin(m Magasin) (mag Magasin) {
	serv.ResponseBuilder().AddHeader("Access-Control-Allow-Origin", "*") // debug only : permet les requetes cross-origin

	fmt.Printf("AddMagasin\n")

	if len(m.Nom) == 0 {
		serv.ResponseBuilder().SetResponseCode(400).WriteAndOveride([]byte("{ \"message\" : \"Il faut indiquer le nom du magasin dans le json en entree\" }"))
	} else {
		mag := createMagasin(DatabaseSingleton.db, m.Nom)
		if nil != mag {
			serv.ResponseBuilder().SetResponseCode(200)
			return *mag
		}
	}
	return
}

func (serv MagasinService) PutMagasin(m Magasin) (mag Magasin) {
	serv.ResponseBuilder().AddHeader("Access-Control-Allow-Origin", "*") // debug only : permet les requetes cross-origin

	fmt.Printf("PutMagasin\n")

	if m.Idt == 0 || len(m.Nom) == 0 {
		serv.ResponseBuilder().SetResponseCode(400).WriteAndOveride([]byte("{ \"message\" : \"Il faut indiquer idt,nom du magasin dans le json en entree\" }"))
		return
	}

	existingMagasin := loadMagasin(DatabaseSingleton.db, m.Idt)
	if existingMagasin == nil {
		serv.ResponseBuilder().SetResponseCode(500).WriteAndOveride([]byte("{ \"message\" : \"Le Magasin indique n'existe pas\" }"))
		return
	}

	magasin := updateMagasin(DatabaseSingleton.db, m.Idt, m.Nom)
	if nil != magasin {
		serv.ResponseBuilder().SetResponseCode(200)
		return *magasin
	}
	return
}

func (serv MagasinService) DeleteMagasin(Id int64) (m Magasin) {
	serv.ResponseBuilder().AddHeader("Access-Control-Allow-Origin", "*") // debug only : permet les requetes cross-origin

	fmt.Printf("DeleteMagasin %d\n", Id)

	existingMagasin := loadMagasin(DatabaseSingleton.db, Id)
	if existingMagasin == nil {
		serv.ResponseBuilder().SetResponseCode(500).WriteAndOveride([]byte("{ \"message\" : \"Le Magasin indique n'existe pas\" }"))
	}

	mag := deleteMagasin(DatabaseSingleton.db, Id)
	if nil != mag {
		serv.ResponseBuilder().SetResponseCode(200)
		return *mag
	}
	return
}
