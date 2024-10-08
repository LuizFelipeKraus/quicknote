package main

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
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

	file := []string{
		"views/templates/base.html",
		"views/templates/pages/note-view.html",
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Nota não encontrada!", http.StatusNotFound)
		slog.Error(fmt.Sprintf("Aconteceu um erro ao executar!", http.StatusInternalServerError))
		return
	}

	t, err := template.ParseFiles(file...)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Aconteceu um erro ao executar!", http.StatusInternalServerError)
		slog.Error(fmt.Sprintf("Aconteceu um erro ao executar!", http.StatusInternalServerError))
		return
	}

	t.ExecuteTemplate(w, "base", id)
}

func noteList(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	file := []string{
		"views/templates/base.html",
		"views/templates/pages/home.html",
	}
	t, err := template.ParseFiles(file...)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Aconteceu um erro ao executar!", http.StatusInternalServerError)
		slog.Error(fmt.Sprintf("Aconteceu um erro ao executar! %d", http.StatusInternalServerError))
		return
	}

	t.ExecuteTemplate(w, "base", nil)

}

func noteNew(w http.ResponseWriter, r *http.Request) {
	file := []string{
		"views/templates/base.html",
		"views/templates/pages/note-new.html",
	}
	t, err := template.ParseFiles(file...)
	if err != nil {
		http.Error(w, "Aconteceu um erro ao executar!", http.StatusInternalServerError)
		slog.Error(fmt.Sprintf("Aconteceu um erro ao executar! %d", http.StatusInternalServerError))
		return
	}

	t.ExecuteTemplate(w, "base", nil)
}

func main() {
	config := loadConfig()
	logger := newLogger(os.Stderr, config.GetLevelLog())
	slog.SetDefault(logger)

	slog.Info(fmt.Sprintf("Servidor rodando na porta %s", config.ServerPort))
	mux := http.NewServeMux()

	staticHandler := http.FileServer(http.Dir("views/static"))

	mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))

	mux.HandleFunc("/", noteList)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/new", noteNew)
	mux.HandleFunc("/note/create", noteCreate)

	http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort), mux)
}
