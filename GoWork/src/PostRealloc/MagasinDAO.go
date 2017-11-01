package main

import (
	"fmt"
	"runtime"
)

type BouchonWrapper struct {
	magasins []Magasin
}

var bouchonWrapper *BouchonWrapper

func bouchonFinalizer(bouchonWrapper *BouchonWrapper) {
	fmt.Println("finalizer has run for BouchonWrapper")
}

func buildBouchon() {
	if bouchonWrapper == nil {
		fmt.Println("Create bouchonWrapper")
		bouchonWrapper = &BouchonWrapper{}

		runtime.SetFinalizer(bouchonWrapper, bouchonFinalizer)

		bouchonWrapper.magasins = make([]Magasin, 2)

		bouchonWrapper.magasins[0] = *NewMagasin(0, "Toulouse")
		bouchonWrapper.magasins[1] = *NewMagasin(1, "Gourdon")
	} else {
		fmt.Println("Existing bouchonWrapper")
	}
}

func loadMagasins() []Magasin {
	buildBouchon()

	return bouchonWrapper.magasins
}

func loadMagasin(idt_ int64) *Magasin {
	buildBouchon()

	if idt_ == 0 {
		return &bouchonWrapper.magasins[0]
	} else if idt_ == 1 {
		return &bouchonWrapper.magasins[1]
	}
	return nil
}

func createMagasin(nom_ string) *Magasin {
	buildBouchon()
	nbMagasin := len(bouchonWrapper.magasins)
	fmt.Println("nbMagasin %d ", nbMagasin)
	magasin := NewMagasin(int64(nbMagasin), nom_)
	newSlice := make([]Magasin, nbMagasin+1)
	copy(newSlice, bouchonWrapper.magasins)
	newSlice[nbMagasin] = *magasin
	bouchonWrapper.magasins = newSlice
	return magasin
}

func resetMagasinsBouchon() []Magasin {
	bouchonWrapper = nil
	buildBouchon()
	runtime.GC()
	return bouchonWrapper.magasins
}
