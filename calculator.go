package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type dataNumbers struct {
	Num1     int `json:Num1`
	Num2     int `json:Num2`
	SumTotal int `json:SumTotal`
}

func main() {
	fmt.Println("Calculadora")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/suma", sum)
	log.Fatal(http.ListenAndServe(":4300", router))
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido a la Calculadora API")
}

func sum(w http.ResponseWriter, r *http.Request) {
	var newNumbers dataNumbers

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "Inserte un dato valido")
	}

	json.Unmarshal(reqBody, &newNumbers)
	newNumbers.SumTotal = newNumbers.Num1 + newNumbers.Num2

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newNumbers)
}
