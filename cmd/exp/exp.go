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
	//insertPost()
	//insertPostWithReturn()
	selectById()
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

func insertPostWithReturn() {
	title := "POST 2"
	content := "Conteudo do post 2"
	author := "Luiz"

	query := `
		INSERT INTO posts(title, content, author) VALUES ($1, $2, $3) RETURNING id;
	`

	row := conn.QueryRow(context.Background(), query, title, content, author)
	var id int
	if err := row.Scan(&id); err != nil {
		panic(err)
	}

	fmt.Println("post criado com sucesso id = ", id)
}

func selectById() {
	var title, content, author string
	id := 2
	query := "SELECT title, content, author FROM posts WHERE id = $1;"
	row := conn.QueryRow(context.Background(), query, id)
	if err := row.Scan(&title, &content, &author); err != nil {
		panic(err)
	}

	fmt.Printf("POST: title = %s, content = %s, author = %s \n", title, content, author)
}
