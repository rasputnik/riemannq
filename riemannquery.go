package main

import (
	"encoding/json"
	"flag"
	"github.com/amir/raidman"
	"log"
	"os"
)

func main() {

	var server, query string
	flag.StringVar(&server, "s", "riemann:5555", "riemann server")
	flag.StringVar(&query, "q", "true", "query string")
	flag.Parse()

	conn, err := raidman.Dial("tcp", server)

	if err != nil {
		log.Fatal(err)
	}

	events, err := conn.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	output, err := json.Marshal(events)
	if err != nil {
		log.Fatal("json encoding issue:", err)
	}
	os.Stdout.Write(output)

	conn.Close()
}
