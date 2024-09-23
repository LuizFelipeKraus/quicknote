package main

import (
	"fmt"
	"net/http"
)

func noteCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(405)
		fmt.Fprint(w, "Método não permitido!")
		return
	}
	fmt.Fprint(w, "Criar Anotação!")
}

func noteView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Visualizar Anotação!")
}

func noteList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Visualizar Lista Anotações!")
}

func main() {
	fmt.Println("Servidor rodando na porta 5000!")
	mux := http.NewServeMux()

	mux.HandleFunc("/", noteList)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/create", noteCreate)

	http.ListenAndServe(":5000", mux)
}
