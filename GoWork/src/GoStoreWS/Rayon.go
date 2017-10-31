package main

type Rayon struct {
	Idt int64 `json:"idt,omitempty"`
	IdtMagasin int64 `json:"idt_magasin,omitempty"`
	Nom string `json:"nom,omitempty"`
	NomImage string `json:"nom_image,omitempty"`
}

func NewRayon(idt_ int64,idtMagasin_ int64,nom_ string,nomImage_ string) *Rayon {
	return &Rayon{Idt:idt_,IdtMagasin:idtMagasin_,Nom:nom_,NomImage:nomImage_}
}