package main

import (
	"html/template"
	"os"
)

type TemplateDate struct {
	Nome  string
	Idade int
}

func main() {
	t, err := template.ParseFiles("hello.html")
	if err != nil {
		panic(err)
	}
	data := TemplateDate{Nome: "Robson"}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
