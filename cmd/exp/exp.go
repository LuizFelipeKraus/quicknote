package main

import (
	"fmt"
	"html/template"
	"os"
)

type TemplateDate struct {
	Nome  string
	Idade int
}

func main() {
	t, err := template.ParseFiles("b.html", "a.html")

	fmt.Println(t.Name())

	if err != nil {
		panic(err)
	}

	//err = t.Execute(os.Stdout, nil)
	err = t.ExecuteTemplate(os.Stdout, "a.html", nil)
	if err != nil {
		panic(err)
	}
}
