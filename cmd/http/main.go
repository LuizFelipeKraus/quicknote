package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/LuizFelipeKraus/quicknotes/internal/handlers"
)

func main() {
	config := loadConfig()
	logger := newLogger(os.Stderr, config.GetLevelLog())
	slog.SetDefault(logger)

	slog.Info(fmt.Sprintf("Servidor rodando na porta %s", config.ServerPort))
	mux := http.NewServeMux()

	staticHandler := http.FileServer(http.Dir("views/static"))

	mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))

	notehandler := handlers.NewNotehandler()

	mux.HandleFunc("/", notehandler.NoteList)
	mux.HandleFunc("/note/view", notehandler.NoteView)
	mux.HandleFunc("/note/new", notehandler.NoteNew)
	mux.HandleFunc("/note/create", notehandler.NoteCreate)

	http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort), mux)
}
