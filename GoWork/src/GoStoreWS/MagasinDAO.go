package main

import (
	"database/sql"        // package SQL
	_ "github.com/lib/pq" // driver Postgres
)

func loadMagasins(db *sql.DB) []Magasin {
	result := make([]Magasin, 0, 0)

	rows, err := db.Query("select idt,nom from magasin order by idt")
	if err != nil {
		return result
	}
	defer rows.Close()

	for rows.Next() {
		var idt int64
		var nom string
		err = rows.Scan(&idt, &nom)
		if err != nil {
			return make([]Magasin, 0, 0)
		}
		newMagasin := NewMagasin(idt, nom)
		result = append(result, *newMagasin)
	}
	return result
}

func rowResultSetToMagasin(row *sql.Row) (*Magasin, error) {
	var err error
	var idt int64
	var nom string
	err = row.Scan(&idt, &nom)
	if err != nil {
		return nil, err
	}
	return NewMagasin(idt, nom), nil
}

func rowsResultSetToMagasin(rows *sql.Rows) (*Magasin, error) {
	var err error
	if rows.Next() {
		var idt int64
		var nom string
		err = rows.Scan(&idt, &nom)
		if err != nil {
			return nil, err
		}
		return NewMagasin(idt, nom), nil
	}
	return nil, err
}

func loadMagasin(db *sql.DB, idt_ int64) *Magasin {
	rows, err := db.Query("select idt,nom from magasin where idt=$1", idt_)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var magasin *Magasin
	magasin, err = rowsResultSetToMagasin(rows)
	return magasin
}

func createMagasin(db *sql.DB, nom_ string) *Magasin {
	rows := db.QueryRow("insert into magasin(nom) values($1) returning idt,nom", nom_)

	var magasin *Magasin
	var err error
	magasin, err = rowResultSetToMagasin(rows)
	if err != nil {
		return nil
	}
	return magasin
}

func updateMagasin(db *sql.DB, idt_ int64, nom_ string) *Magasin {
	rows := db.QueryRow("update magasin set nom=$1 where idt=$2 returning idt,nom", nom_, idt_)

	var magasin *Magasin
	var err error
	magasin, err = rowResultSetToMagasin(rows)
	if err != nil {
		return nil
	}
	return magasin
}

func deleteMagasin(db *sql.DB, idt_ int64) *Magasin {
	rows := db.QueryRow("delete from magasin where idt=$1 returning idt,nom", idt_)

	var magasin *Magasin
	var err error
	magasin, err = rowResultSetToMagasin(rows)
	if err != nil {
		return nil
	}
	return magasin
}
