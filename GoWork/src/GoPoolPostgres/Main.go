package main

import (
	"bytes"
	"fmt"
	"log"
	"sync"
	"time"
)

var DatabaseSingleton *Database
var mutexWorkerInc sync.Mutex
var nbWorkersDone int
var nbWorkersWaiting int

func insertWithLongTx(idThread int) {
	var sStmt string = "insert into magasin(nom) values($1)"

	var buffer bytes.Buffer
	for pad := 0; pad <= idThread; pad++ {
		buffer.WriteString("  ")
	}
	padding := buffer.String()

	fmt.Printf("%s --> %d ask for DB Cnx\n", padding, idThread)
	transaction, err := DatabaseSingleton.db.Begin()
	fmt.Printf("%s --> %d Done (ask for DB Cnx)\n", padding, idThread)
	if err == nil {
		stmt, err := transaction.Prepare(sStmt)
		if err != nil {
			log.Fatal(err)
		} else {
			for i := 0; i < 10; i++ {
				nom := fmt.Sprintf("Magasin_%d_%d", idThread, i)
				fmt.Printf("%s --> %s\n", padding, nom)
				_, err := stmt.Exec(nom)
				if err != nil {
					log.Fatal(err)
				}
				time.Sleep(time.Duration(1) * time.Second)

			}
			stmt.Close()
			fmt.Printf("%s --> %d Commit\n", padding, idThread)
			transaction.Commit()
			fmt.Printf("%s --> %d Done (Commit)\n", padding, idThread)
		}
	} else {
		log.Fatal(err)
	}

	fmt.Printf("%s --> %d Pre-Dead\n", padding, idThread)
	mutexWorkerInc.Lock()
	nbWorkersDone += 1
	mutexWorkerInc.Unlock()
	fmt.Printf("%s --> %d Dead\n", padding, idThread)
}

func main() {

	db, err := connectDB()
	if err == nil {
		DatabaseSingleton = &Database{db: db}

		nbWorkersDone = 0
		nbWorkersWaiting = 20
		for i := 0; i < nbWorkersWaiting; i++ {
			fmt.Printf("create worker %d\n", i)
			go insertWithLongTx(i)
			time.Sleep(time.Duration(2) * time.Second)
		}

		for nbWorkersDone < nbWorkersWaiting {
			fmt.Printf("waiting for %d\n", (nbWorkersWaiting - nbWorkersDone))
			time.Sleep(time.Duration(1) * time.Second)
		}
	}
}
