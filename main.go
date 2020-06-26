package main

import (
	"log"
	"net/http"
	"os"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"gangsters": [
		{
			"name": "Mruczek",
			"description": "Szef kociej mafii, poszukiwany listem gończym. Skazany za handel niecertyfikowaną karmą na czarnym rynku.",
			"price": "$1000",
			"photo": "https://i.ibb.co/rbgVDQ0/mruczek.jpg"
		},
		{
			"name": "Puszek",
			"description": "Lokalny raper, od czasu ataku na funkcjonariusza policji, w trakcie przeszukania, przebywa w bliżej nieokreślonym podziemiu.",
			"price": "$400",
			"photo": "https://i.ibb.co/mHs5wFW/puszek.png"
		},
		{
			"name": "Bambam i Kicio",
			"description": "Bracia. Należą do żylety lokalnego klubu sportowego FC Meow Meow. Obaj poszukiwani za przeprowadzenie akcji \"strącenie\" polegającej na zrzucaniu przemdiotów z półek.",
			"price": "$700",
			"photo": "https://i.ibb.co/1zXX2S8/bambamikicio.jp://i.ibb.co/1zXX2S8/bambamikicio.jpg"
		}
	]}`))
}

func main() {
	http.HandleFunc("/", home)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5001"
	}
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
