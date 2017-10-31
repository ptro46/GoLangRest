package main

var bouchon []Magasin

func loadMagasins() []Magasin {
	bouchon := make([]Magasin, 2)

	bouchon[0] = *NewMagasin(0, "Toulouse")
	bouchon[1] = *NewMagasin(1, "Gourdon")

	return bouchon
}

func loadMagasin(idt_ int64) *Magasin {
	if idt_ == 0 {
		return &bouchon[0]
	} else if idt_ == 1 {
		return &bouchon[1]
	}
	return nil
}
