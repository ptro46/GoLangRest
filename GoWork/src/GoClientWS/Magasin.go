package main

type Magasin struct {
	Idt int64 `json:"idt,omitempty"`
	Nom string `json:"nom,omitempty"`
}

func NewMagasin(idt_ int64,nom_ string) *Magasin {
	return &Magasin{Idt:idt_,Nom:nom_}
}