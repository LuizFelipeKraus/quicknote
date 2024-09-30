package main

import (
	"flag"
	"fmt"
)

func main() {
	port := flag.String("port", "7000", "Server Port")
	flag.Parse()
	fmt.Println("Server is running on port", *port)

	var port2 string
	flag.StringVar(&port2, "port2", "6000", "Server Port")
	fmt.Println("Server is running on port", port2)
}
