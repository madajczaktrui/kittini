package main

import (
	"log"
	"net/http"
	"os"
	"encoding/json"
	"math/rand"
	"time"
	"io/ioutil"
)

type Gangsters struct {
	Gangsters []Gangster `json:"gangsters"`
}

type Gangster struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Price string `json:"price"`
	Photo string `json:"photo"`
}

func randomGangster() Gangster {
	data, _ := ioutil.ReadFile("gangsters.json")
	var gangsters Gangsters
	json.Unmarshal([]byte(data), &gangsters)
	rand.Seed(time.Now().Unix())
	return gangsters.Gangsters[rand.Intn(len(gangsters.Gangsters))]
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(randomGangster())
	if err == nil {
		w.Write(data)
	}
}

func main() {
	http.HandleFunc("/", home)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5001"
	}
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
