package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/TanglingTreats/mugen-api/api"
	"github.com/TanglingTreats/mugen-api/dotenv"
)

func main() {
	// Set flags
	env := flag.String("env", ".env", "env file to use")
	listenAddr := flag.String("listenaddr", ":8080", "server address")
	flag.Parse()

	dotenv.InitEnv(*env)

	fmt.Println("Starting RESTful service")

	server := api.NewServer(*listenAddr)

	fmt.Printf("Listening at %s\n", *listenAddr)

	log.Fatal(server.Start())
}
