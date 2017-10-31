package main
import (
	"github.com/fromkeith/gorest" // go get github.com/fromkeith/gorest
	"fmt"
)



//************************Define Service***************************
type RayonService struct{
	//Service level config
	gorest.RestService    `root:"/StoreWS/api/rayon/" consumes:"application/json" produces:"application/json" charset:"utf-8" allowGzip:1 `

	//End-Point level configs: Field names must be the same as the corresponding method names,
	// but not-exported (starts with lowercase)
	rayonDetails 		gorest.EndPoint `method:"GET" path:"/{Id:int64}" output:"Rayon"`
	rayonProduitList	gorest.EndPoint `method:"GET" path:"/{Id:int64}/produits" output:"[]Produit"`
	postRayon			gorest.EndPoint `method:"POST" path:"/" postdata:"Rayon" output:"Rayon"`
	putRayon			gorest.EndPoint `method:"PUT" path:"/" postdata:"Rayon" output:"Rayon"`
	deleteRayon			gorest.EndPoint `method:"DELETE" path:"/{Id:int64}" output:"Rayon"`
}

func(serv RayonService) init() {
}
//Handler Methods: Method names must be the same as in config, but exported (starts with uppercase)

func(serv RayonService) RayonDetails(Id int64) (m Rayon){
	serv.ResponseBuilder().AddHeader("Access-Control-Allow-Origin","*") // debug only : permet les requetes cross-origin

	fmt.Printf("RayonDetails\n")

	db, err := connectDB()
	if err == nil {
		defer db.Close()

		rayon:=loadRayon(db,Id)
		if nil!=rayon {
			serv.ResponseBuilder().SetResponseCode(200)
			return *rayon
		} else {
			serv.ResponseBuilder().SetResponseCode(500).WriteAndOveride([]byte("{ \"message\" : \"Le Rayon indique n'existe pas\" }"))
		}
	}
	serv.ResponseBuilder().SetResponseCode(404).Overide(true)
	return
}

func(serv RayonService) RayonProduitList(Id int64) []Produit {
	serv.ResponseBuilder().AddHeader("Access-Control-Allow-Origin","*") // debug only : permet les requetes cross-origin

	fmt.Printf("RayonProduitList\n")

	result:=make([]Produit,0,0)

	db, err := connectDB()
	if err == nil {
		defer db.Close()

		rayon:=loadRayon(db,Id)
		if nil!=rayon {
			serv.ResponseBuilder().SetResponseCode(200)
			return loadProduitsFromRayon(db,Id)
		} else {
			serv.ResponseBuilder().SetResponseCode(500).WriteAndOveride([]byte("{ \"message\" : \"Le Rayon indique n'existe pas\" }"))
			return result
		}
	}
	serv.ResponseBuilder().SetResponseCode(404).Overide(true)
	return result
}

func(serv RayonService) PostRayon(r Rayon) (ray Rayon) {
	serv.ResponseBuilder().AddHeader("Access-Control-Allow-Origin","*") // debug only : permet les requetes cross-origin

	fmt.Printf("PostRayon\n")

	db, err := connectDB()
	if err == nil {
		defer db.Close()

		if r.IdtMagasin == 0 || len(r.Nom) == 0 || len(r.NomImage) == 0 {
			serv.ResponseBuilder().SetResponseCode(400).WriteAndOveride([]byte("{ \"message\" : \"Il faut indiquer le idt_magasin,nom,nom_image du rayon dans le json en entree\" }"))
			return
		}

		existingMagasin := loadMagasin(db,r.IdtMagasin);
		if ( existingMagasin == nil ) {
			serv.ResponseBuilder().SetResponseCode(500).WriteAndOveride([]byte("{ \"message\" : \"Le Magasin indique n'existe pas\" }"))  
			return 
		}


		ray:=createRayon(db,r.IdtMagasin,r.Nom,r.NomImage)
		if nil!=ray {
			serv.ResponseBuilder().SetResponseCode(200)
			return *ray
		}
	}
	serv.ResponseBuilder().SetResponseCode(404).Overide(true)
	return
}

func(serv RayonService) PutRayon(r Rayon) (ray Rayon) {
	serv.ResponseBuilder().AddHeader("Access-Control-Allow-Origin","*") // debug only : permet les requetes cross-origin

	fmt.Printf("PutRayon\n")

	db, err := connectDB()
	if err == nil {
		defer db.Close()

		if r.Idt == 0 {
			serv.ResponseBuilder().SetResponseCode(400).WriteAndOveride([]byte("{ \"message\" : \"Il faut indiquer idt du rayon dans le json en entree\" }"))  
			return
		}

		existingRayon := loadRayon(db,r.Idt);
		if ( existingRayon == nil ) {
			serv.ResponseBuilder().SetResponseCode(500).WriteAndOveride([]byte("{ \"message\" : \"Le Rayon indique n'existe pas\" }"))  
			return
		}

		if ( len(r.NomImage) == 0 ) {
			r.NomImage = existingRayon.NomImage;
		}

		if ( len(r.Nom) == 0 ) {
			r.Nom = existingRayon.Nom ;
		}

		ray:=updateRayon(db,r.Idt,r.Nom,r.NomImage)
		if nil!=ray {
			serv.ResponseBuilder().SetResponseCode(200)
			return *ray
		}
	}
	serv.ResponseBuilder().SetResponseCode(404).Overide(true)
	return
}

func(serv RayonService) DeleteRayon(Id int64) (r Rayon){
	serv.ResponseBuilder().AddHeader("Access-Control-Allow-Origin","*") // debug only : permet les requetes cross-origin

	fmt.Printf("DeleteRayon %d\n",Id)

	db, err := connectDB()
	if err == nil {
		defer db.Close()

		existingRayon := loadRayon(db,Id);
		if ( existingRayon == nil ) {
			serv.ResponseBuilder().SetResponseCode(500).WriteAndOveride([]byte("{ \"message\" : \"Le Rayon indique n'existe pas\" }"))  
		}

		ray:=deleteRayon(db,Id)
		if nil!=ray {
			serv.ResponseBuilder().SetResponseCode(200)
			return *ray
		}
	}
	serv.ResponseBuilder().SetResponseCode(404).Overide(true)
	return
}

