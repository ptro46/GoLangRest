package main

type Produit struct {
	Idt int64 `json:"idt,omitempty"`
	IdtRayon int64 `json:"idt_rayon,omitempty"`
	Nom string `json:"nom,omitempty"`
	NomImage string `json:"nom_image,omitempty"`
	Prix float64 `json:"prix,omitempty"`
}

func NewProduit(idt_ int64,idtRayon_ int64,nom_ string,nomImage_ string,prix_ float64) *Produit {
	return &Produit{Idt:idt_,IdtRayon:idtRayon_,Nom:nom_,NomImage:nomImage_,Prix:prix_}
}