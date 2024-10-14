package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

func main() {
	var err error
	url := "postgres://postgres:secret@localhost:5432/postgres"
	conn, err = pgx.Connect(context.Background(), url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Conexão bem sucedida")
	defer conn.Close(context.Background())
	createTable()
	insertPost()
}

func createTable() {
	query := `
		CREATE TABLE IF NOT EXISTS posts (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			content TEXT, 
			author TEXT NOT NULL
		)
	`
	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Table posts criada")
}

func insertPost() {
	title := "POST 1"
	content := "Conteudo do post 1"
	author := "Luiz"

	query := `
		INSERT INTO posts(title, content, author) VALUES ($1, $2, $3);
	`

	_, err := conn.Exec(context.Background(), query, title, content, author)
	if err != nil {
		panic(err)
	}
	fmt.Println("post criado com sucesso")
}
