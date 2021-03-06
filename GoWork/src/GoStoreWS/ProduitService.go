package main

import (
	"fmt"
	"github.com/fromkeith/gorest" // go get github.com/fromkeith/gorest
)

//************************Define Service***************************
type ProduitService struct {
	//Service level config
	gorest.RestService `root:"/StoreWS/api/produit/" consumes:"application/json" produces:"application/json" charset:"utf-8" allowGzip:1 `

	//End-Point level configs: Field names must be the same as the corresponding method names,
	// but not-exported (starts with lowercase)
	produitDetails gorest.EndPoint `method:"GET" path:"/{Id:int64}" output:"Produit"`
	postProduit    gorest.EndPoint `method:"POST" path:"/" postdata:"Produit" output:"Produit"`
	putProduit     gorest.EndPoint `method:"PUT" path:"/" postdata:"Produit" output:"Produit"`
	deleteProduit  gorest.EndPoint `method:"DELETE" path:"/{Id:int64}" output:"Produit"`
}

func (serv ProduitService) init() {
}

//Handler Methods: Method names must be the same as in config, but exported (starts with uppercase)

func (serv ProduitService) ProduitDetails(Id int64) (p Produit) {
	serv.ResponseBuilder().AddHeader("Access-Control-Allow-Origin", "*") // debug only : permet les requetes cross-origin

	fmt.Printf("ProduitDetails\n")

	produit := loadProduit(DatabaseSingleton.db, Id)
	if nil != produit {
		serv.ResponseBuilder().SetResponseCode(200)
		return *produit
	} else {
		serv.ResponseBuilder().SetResponseCode(500).WriteAndOveride([]byte("{ \"message\" : \"Le Produit indique n'existe pas\" }"))
	}
	return
}

func (serv ProduitService) PostProduit(p Produit) (pro Produit) {
	serv.ResponseBuilder().AddHeader("Access-Control-Allow-Origin", "*") // debug only : permet les requetes cross-origin

	fmt.Printf("PostProduit\n")

	if p.IdtRayon == 0 || len(p.Nom) == 0 || len(p.NomImage) == 0 || p.Prix == 0 {
		serv.ResponseBuilder().SetResponseCode(400).WriteAndOveride([]byte("{ \"message\" : \"Il faut indiquer le idt_rayon,nom,nom_image,prix du produit dans le json en entree\" }"))
		return
	}

	rayon := loadRayon(DatabaseSingleton.db, p.IdtRayon)
	if rayon == nil {
		serv.ResponseBuilder().SetResponseCode(500).WriteAndOveride([]byte("{ \"message\" : \"Le Rayon indique n'existe pas\" }"))
		return
	}

	produit := createProduit(DatabaseSingleton.db, p.IdtRayon, p.Nom, p.NomImage, p.Prix)
	if nil != produit {
		serv.ResponseBuilder().SetResponseCode(200)
		return *produit
	}
	return
}

func (serv ProduitService) PutProduit(p Produit) (pro Produit) {
	serv.ResponseBuilder().AddHeader("Access-Control-Allow-Origin", "*") // debug only : permet les requetes cross-origin

	fmt.Printf("PutProduit\n")

	if p.Idt == 0 {
		serv.ResponseBuilder().SetResponseCode(400).WriteAndOveride([]byte("{ \"message\" : \"Il faut indiquer idt du produit dans le json en entree\" }"))
		return
	}

	existingProduit := loadProduit(DatabaseSingleton.db, p.Idt)
	if existingProduit == nil {
		serv.ResponseBuilder().SetResponseCode(500).WriteAndOveride([]byte("{ \"message\" : \"Le Produit indique n'existe pas\" }"))
		return
	}

	if len(p.NomImage) == 0 {
		p.NomImage = existingProduit.NomImage
	}

	if len(p.Nom) == 0 {
		p.Nom = existingProduit.Nom
	}

	if p.Prix == 0 {
		p.Prix = existingProduit.Prix
	}

	fmt.Printf("Nom %s, NomImage %s\n", p.Nom, p.NomImage)

	produit := updateProduit(DatabaseSingleton.db, p.Idt, p.Nom, p.NomImage, p.Prix)
	if nil != produit {
		serv.ResponseBuilder().SetResponseCode(200)
		return *produit
	}
	return
}

func (serv ProduitService) DeleteProduit(Id int64) (p Produit) {
	serv.ResponseBuilder().AddHeader("Access-Control-Allow-Origin", "*") // debug only : permet les requetes cross-origin

	fmt.Printf("DeleteProduit %d\n", Id)

	existingProduit := loadProduit(DatabaseSingleton.db, Id)
	if existingProduit == nil {
		serv.ResponseBuilder().SetResponseCode(500).WriteAndOveride([]byte("{ \"message\" : \"Le Produit indique n'existe pas\" }"))
		return
	}

	produit := deleteProduit(DatabaseSingleton.db, Id)
	if nil != produit {
		serv.ResponseBuilder().SetResponseCode(200)
		return *produit
	}
	return
}
