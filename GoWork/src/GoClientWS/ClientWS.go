package main

import (
	//	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func testGetAllMagasins() {
	res, err := http.Get("http://localhost:8080/StoreWS/api/magasin")
	if err != nil {
		log.Fatal(err)
	}
	magasins, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("testGetAllMagasins %s\n", magasins)
}

func testPostOneMagasin(name string) (magasin *Magasin) {
	strBody := "{\"nom\":\"" + name + "\"}"
	//	body := bytes.NewBuffer([]byte(strBody))
	res, err := http.Post("http://localhost:8080/StoreWS/api/magasin", "application/json", strings.NewReader(strBody))
	if err != nil {
		log.Fatal(err)
	} else {
		defer res.Body.Close()
		magasinJSON, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("testPostOneMagasin %s\n", magasinJSON)
		magasin := new(Magasin)
		errUnmarshal := json.Unmarshal([]byte(magasinJSON), magasin)
		if errUnmarshal != nil {
			log.Fatal(errUnmarshal)
		}
		fmt.Printf("testPostOneMagasin//Unmarshal %+v\n", magasin)
		return magasin
	}
	return nil
}

func testPutMagasin(idt int, name string) {
	strBody := "{\"idt\":" + fmt.Sprintf("%d", idt) + ",\"nom\":\"" + name + "\"}"
	//	body := bytes.NewBuffer([]byte(strBody))

	client := &http.Client{}
	request, err := http.NewRequest("PUT", "http://localhost:8080/StoreWS/api/magasin", strings.NewReader(strBody))
	request.ContentLength = int64(len(strBody))
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("testPutMagasin %s\n", contents)
	}

}

func testDeleteMagasin(idt int) {
	strUrl := "http://localhost:8080/StoreWS/api/magasin/" + fmt.Sprintf("%d", idt)
	client := &http.Client{}
	request, err := http.NewRequest("DELETE", strUrl, strings.NewReader(""))
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("testDeleteMagasin %s\n", contents)
	}
}

func main() {
	testGetAllMagasins()

	newMagasin := testPostOneMagasin("Toulouse")
	testGetAllMagasins()

	testPutMagasin(int(newMagasin.Idt), "Toulouse Blagnac")
	testGetAllMagasins()

	testDeleteMagasin(int(newMagasin.Idt))
	testGetAllMagasins()

	fmt.Println("Finished")
}
