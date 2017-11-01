package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func testUnmashallOneMagasin() {
	magasinJSON := `{"idt":3,"nom":"Toulouse"}`
	fmt.Printf("testMarshallOneMagasin %s\n", magasinJSON)
	magasin := new(Magasin)
	errUnmarshal := json.Unmarshal([]byte(magasinJSON), magasin)
	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	} else {
		fmt.Printf("testUnmashallOneMagasin//Unmarshal %+v\n", magasin)
	}
}

func testUnmarshallArrayOfMagasins() {
	magasinsJSON := `[{"idt":3,"nom":"Toulouse"},{"idt":2,"nom":"Gourdon"}]`
	var magasins []Magasin
	errUnmarshal := json.Unmarshal([]byte(magasinsJSON), &magasins)
	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	} else {
		for _, mag := range magasins {
			fmt.Printf("testUnmarshallArrayOfMagasins//Unmarshal %+v\n", mag)
		}
	}
}

func testUnmarshallComposite() {
	unknowJSON := `{"idt":1,"nom":"Toulouse","rayons":[{"idt":1,"nom":"Eaux"},{"idt":2,"nom":"Viandes"}]}`
	var parsed map[string]interface{}
	errUnmarshal := json.Unmarshal([]byte(unknowJSON), &parsed)
	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	} else {
		idtMagasin := parsed["idt"].(float64)
		nomMagasin := parsed["nom"].(string)
		magasin := NewMagasin(int64(idtMagasin), nomMagasin)
		fmt.Printf("testUnmarshallComposite//Unmarshal %+v\n", magasin)

		arrayOfRayons := parsed["rayons"].([]interface{})
		for _, oneRayon := range arrayOfRayons {
			rayonMap := oneRayon.(map[string]interface{})
			idtRayon := rayonMap["idt"].(float64)
			nomRayon := rayonMap["nom"].(string)
			fmt.Printf("    Rayon :: %d %s\n", int64(idtRayon), nomRayon)

		}
	}
}

func main() {

	testUnmashallOneMagasin()
	testUnmarshallArrayOfMagasins()
	testUnmarshallComposite()

	fmt.Println("Finished")
}
