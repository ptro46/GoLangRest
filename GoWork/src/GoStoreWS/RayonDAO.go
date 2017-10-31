package main

import (
	"database/sql"			// package SQL
	_ "github.com/lib/pq"	// driver Postgres
)

func loadRayons(db *sql.DB) []Rayon {
	result:=make([]Rayon,0,0)

	rows, err := db.Query("select idt,idt_magasin,nom,nom_image from rayon order by idt")
	if err != nil {
		return result
	}


	for rows.Next() {
		var idt int64
		var idtRayon int64
		var nom string
		var nomImage string
		err = rows.Scan(&idt, &idtRayon, &nom, &nomImage)
		if err != nil {
			return make([]Rayon,0,0)
		}
		newRayon:=NewRayon(idt,idtRayon,nom,nomImage)
		result=append(result,*newRayon)
	}
	return result
}

func loadRayonsFromMagasin(db *sql.DB, idtMagasin int64) []Rayon {
	result:=make([]Rayon,0,0)

	rows, err := db.Query("select idt,idt_magasin,nom,nom_image from rayon where idt_magasin=$1 order by idt",idtMagasin)
	if err != nil {
		return result
	}


	for rows.Next() {
		var idt int64
		var idtRayon int64
		var nom string
		var nomImage string
		err = rows.Scan(&idt, &idtRayon, &nom, &nomImage)
		if err != nil {
			return make([]Rayon,0,0)
		}
		newRayon:=NewRayon(idt,idtRayon,nom,nomImage)
		result=append(result,*newRayon)
	}
	return result
}

func rowResultSetToRayon(row *sql.Row) (*Rayon,error) {
	var err error
	var idt int64
	var idtRayon int64
	var nom string
	var nomImage string
	err = row.Scan(&idt, &idtRayon, &nom, &nomImage)
	if err != nil {
		return nil,err
	}
	return NewRayon(idt,idtRayon,nom,nomImage),nil
}

func rowsResultSetToRayon(rows *sql.Rows) (*Rayon,error) {
	var err error
	if rows.Next() {
		var idt int64
		var idtRayon int64
		var nom string
		var nomImage string
		err = rows.Scan(&idt, &idtRayon, &nom, &nomImage)
		if err != nil {
			return nil,err
		}
		return NewRayon(idt,idtRayon,nom,nomImage),nil
	}
	return nil,err
}

func loadRayon(db *sql.DB, idt_ int64) *Rayon {
	rows, err := db.Query("select idt,idt_magasin,nom,nom_image from rayon where idt=$1",idt_)
	if err != nil {
		return nil
	}

	rayon,err := rowsResultSetToRayon(rows)
	return rayon
}

func createRayon(db *sql.DB, idtMagasin_ int64, nom_ string, nomImage_ string) *Rayon {
	rows := db.QueryRow("insert into rayon(idt_magasin,nom,nom_image) values($1,$2,$3) returning idt,idt_magasin,nom,nom_image",idtMagasin_,nom_,nomImage_)

	rayon,err := rowResultSetToRayon(rows)
	if ( err != nil ) {
		return nil
	}
	return rayon
}

func updateRayon(db *sql.DB, idt_ int64, nom_ string, nomImage_ string) *Rayon {
	rows := db.QueryRow("update rayon set nom=$1,nom_image=$2 where idt=$3 returning idt,idt_magasin,nom,nom_image",nom_,nomImage_,idt_)

	rayon,err := rowResultSetToRayon(rows)
	if ( err != nil ) {
		return nil
	}
	return rayon
}

func deleteRayon(db *sql.DB, idt_ int64) *Rayon {
	rows := db.QueryRow("delete from rayon where idt=$1 returning idt,idt_magasin,nom,nom_image",idt_)

	rayon,err := rowResultSetToRayon(rows)
	if ( err != nil ) {
		return nil
	}
	return rayon
}

