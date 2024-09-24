package main

import (
	"html/template"
	"os"
)

type TemplateDate struct {
	Nome string
}

func main() {
	t, err := template.New("teste").Parse("<h1>Hello {{ . }}!</h1>")

	if err != nil {
		panic(err)
	}
	data := TemplateDate{Nome: "Robson"}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
