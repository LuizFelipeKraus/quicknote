package handlers

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"text/template"

	"github.com/LuizFelipeKraus/quicknotes/internal/apperror"
)

type noteHandler struct {
}

func NewNotehandler() *noteHandler {
	return &noteHandler{}
}

func (nh *noteHandler) NoteCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Método não permitido! %d", http.StatusMethodNotAllowed)
		slog.Error(fmt.Sprintf("Aconteceu um erro ao executar! %d", http.StatusInternalServerError))
		return
	}
	fmt.Fprint(w, "Criar Anotação!")
}

func (nh *noteHandler) NoteView(w http.ResponseWriter, r *http.Request) error {
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.RawQuery)

	file := []string{
		"views/templates/base.html",
		"views/templates/pages/note-view.html",
	}

	id := r.URL.Query().Get("id")

	if id == "0" {
		return apperror.WithStatus(errors.New("Não existe esse id "), http.StatusNotFound)
	}

	if id == "" {

		slog.Error(fmt.Sprintf("Aconteceu um erro ao executar! %d", http.StatusInternalServerError))
		return apperror.WithStatus(errors.New("Anotação é obrigatória"), http.StatusBadRequest)
	}

	t, err := template.ParseFiles(file...)
	if err != nil {
		fmt.Println(err)
		//http.Error(w, "Aconteceu um erro ao executar!", http.StatusInternalServerError)
		slog.Error(fmt.Sprintf("Aconteceu um erro ao executar! %d", http.StatusInternalServerError))
		return apperror.WithStatus(errors.New("Não exite template"), http.StatusNotFound)
	}

	return t.ExecuteTemplate(w, "base", id)
}

func (nh *noteHandler) NoteList(w http.ResponseWriter, r *http.Request) {
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

func (nh *noteHandler) NoteNew(w http.ResponseWriter, r *http.Request) {
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
