package main

var nbMagasin int64
var bouchon []Magasin

func loadMagasins() []Magasin {
	bouchon := make([]Magasin, 2)

	bouchon[0] = *NewMagasin(0, "Toulouse")
	bouchon[1] = *NewMagasin(1, "Gourdon")
	nbMagasin = 2
	return bouchon
}

func loadMagasin(idt_ int64) *Magasin {
	bouchon := make([]Magasin, 2)

	bouchon[0] = *NewMagasin(0, "Toulouse")
	bouchon[1] = *NewMagasin(1, "Gourdon")
	nbMagasin = 2

	if idt_ == 0 {
		return &bouchon[0]
	} else if idt_ == 1 {
		return &bouchon[1]
	}
	return nil
}

func createMagasin(nom_ string) *Magasin {
	bouchon := make([]Magasin, 3)

	bouchon[0] = *NewMagasin(0, "Toulouse")
	bouchon[1] = *NewMagasin(1, "Gourdon")
	nbMagasin = 2

	magasin := NewMagasin(nbMagasin, nom_)
	bouchon[nbMagasin] = *magasin
	nbMagasin++
	return magasin
}
