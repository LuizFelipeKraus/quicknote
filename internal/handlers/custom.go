package handlers

import (
	"errors"
	"net/http"
	"text/template"

	"github.com/LuizFelipeKraus/quicknotes/internal/apperror"
)

type HandlerWithError func(w http.ResponseWriter, r *http.Request) error

func (f HandlerWithError) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := f(w, r); err != nil {
		var statusErr apperror.StatusError
		if errors.As(err, &statusErr) {
			if statusErr.StatusCode() == http.StatusNotFound {
				file := []string{
					"views/templates/base.html",
					"views/templates/pages/404.html",
				}
				t, err := template.ParseFiles(file...)
				if err != nil {
					http.Error(w, err.Error(), statusErr.StatusCode())
				}
				t.ExecuteTemplate(w, "base", statusErr.StatusCode())

			}
			http.Error(w, err.Error(), statusErr.StatusCode())
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
