package main

import (
	"net/http"
	"math"
)

func soma( ) string {

	inicio := 0.001

	for j := 0; j <= 1000000; j++ {

		inicio = inicio + math.Sqrt(inicio)
	}
	
	return "Code.education Rocks!"
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(soma()))
	})
	http.ListenAndServe(":8000", nil)
}