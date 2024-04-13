package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/TanglingTreats/mugen-api/api"
	"github.com/TanglingTreats/mugen-api/dotenv"
)

func main() {
	dotenv.InitEnv()

	fmt.Println("Starting RESTful service")
	listenAddr := flag.String("listenaddr", ":8080", "server address")

	server := api.NewServer(*listenAddr)

	fmt.Printf("Listening at %s\n", *listenAddr)

	log.Fatal(server.Start())
}
