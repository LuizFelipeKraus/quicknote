package main

import (
	"fmt"
	"net/http"
)

func noteCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Método não permitido!", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "Criar Anotação!")
}

func noteView(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.RawQuery)

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Nota não encontrada!", http.StatusNotFound)
		return
	}
	note := `
		<div> 
			<h3> Esta é a nota %s</h3>
			<p> Este é o conteudo da anotação</p>
		</div>
	`
	fmt.Fprintf(w, note, id)
}

func noteList(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-type", "text/html")
	fmt.Fprintf(w, "<h1> Lista de anotações e lembretes </h1>")
}

func main() {
	fmt.Println("Servidor rodando na porta 5000!")
	mux := http.NewServeMux()

	mux.HandleFunc("/", noteList)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/create", noteCreate)

	http.ListenAndServe(":5000", mux)
}
