package main

import "net/http"
import (
	"fmt"
	"github.com/fromkeith/gorest" // go get github.com/fromkeith/gorest
)

var DatabaseSingleton *Database

func main() {

	db, err := connectDB()
	if err == nil {
		DatabaseSingleton = &Database{db: db}

		magasinService := new(MagasinService)
		magasinService.init()
		gorest.RegisterService(magasinService)

		rayonService := new(RayonService)
		rayonService.init()
		gorest.RegisterService(rayonService)

		produitService := new(ProduitService)
		produitService.init()
		gorest.RegisterService(produitService)

		http.Handle("/", gorest.Handle())
		http.ListenAndServe(":8080", nil)
		fmt.Print("Finished")
	}
}
