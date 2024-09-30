package main

import (
	"fmt"
	"os"
)

func main() {
	port, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		port = "5000"
	}
	fmt.Println("Servidor está rodando na porta: ", port)

}
