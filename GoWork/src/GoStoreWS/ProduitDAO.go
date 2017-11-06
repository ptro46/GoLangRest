package main

import (
	"database/sql"        // package SQL
	_ "github.com/lib/pq" // driver Postgres
)

func loadProduitsFromRayon(db *sql.DB, idtRayon int64) []Produit {
	result := make([]Produit, 0, 0)

	rows, err := db.Query("select idt,idt_rayon,nom,nom_image,prix from produit where idt_rayon=$1 order by idt", idtRayon)
	if err != nil {
		return result
	}
	defer rows.Close()

	for rows.Next() {
		var idt int64
		var idtRayon int64
		var nom string
		var nomImage string
		var prix float64
		err = rows.Scan(&idt, &idtRayon, &nom, &nomImage, &prix)
		if err != nil {
			return make([]Produit, 0, 0)
		}
		newProduit := NewProduit(idt, idtRayon, nom, nomImage, prix)
		result = append(result, *newProduit)
	}
	return result
}

func rowResultSetToProduit(row *sql.Row) (*Produit, error) {
	var err error
	var idt int64
	var idtRayon int64
	var nom string
	var nomImage string
	var prix float64
	err = row.Scan(&idt, &idtRayon, &nom, &nomImage, &prix)
	if err != nil {
		return nil, err
	}
	return NewProduit(idt, idtRayon, nom, nomImage, prix), nil
}

func rowsResultSetToProduit(rows *sql.Rows) (*Produit, error) {
	var err error
	if rows.Next() {
		var idt int64
		var idtRayon int64
		var nom string
		var nomImage string
		var prix float64
		err = rows.Scan(&idt, &idtRayon, &nom, &nomImage, &prix)
		if err != nil {
			return nil, err
		}
		return NewProduit(idt, idtRayon, nom, nomImage, prix), nil
	}
	return nil, err
}

func loadProduit(db *sql.DB, idt_ int64) *Produit {
	rows, err := db.Query("select idt,idt_rayon,nom,nom_image,prix from produit where idt=$1", idt_)
	if err != nil {
		return nil
	}
	defer rows.Close()

	produit, err := rowsResultSetToProduit(rows)
	return produit
}

func createProduit(db *sql.DB, idtRayon_ int64, nom_ string, nomImage_ string, prix_ float64) *Produit {
	rows := db.QueryRow("insert into produit(idt_rayon,nom,nom_image,prix) values($1,$2,$3,$4) returning idt,idt_rayon,nom,nom_image,prix", idtRayon_, nom_, nomImage_, prix_)

	produit, err := rowResultSetToProduit(rows)
	if err != nil {
		return nil
	}
	return produit
}

func updateProduit(db *sql.DB, idt_ int64, nom_ string, nomImage_ string, prix_ float64) *Produit {
	rows := db.QueryRow("update produit set nom=$1,nom_image=$2,prix=$3 where idt=$4 returning idt,idt_rayon,nom,nom_image,prix", nom_, nomImage_, prix_, idt_)

	produit, err := rowResultSetToProduit(rows)
	if err != nil {
		return nil
	}
	return produit
}

func deleteProduit(db *sql.DB, idt_ int64) *Produit {
	rows := db.QueryRow("delete from produit where idt=$1 returning idt,idt_rayon,nom,nom_image,prix", idt_)

	produit, err := rowResultSetToProduit(rows)
	if err != nil {
		return nil
	}
	return produit
}
