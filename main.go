package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
    "flag"
)

func main() {
    // Flags
    // TODO: Add flag to define which file to read from
    flag.Parse()

    // Get username
    username := flag.Arg(0)
    if username == "" {
      log.Fatal("App must be provided a username")
    }

	file, err := os.Open("config.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		context := strings.Fields(scanner.Text())
		checkAvailability(username, context[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func checkAvailability (username string, url string){
    updatedURL := strings.Replace(url, "{{USERNAME}}", username, -1)
    resp, _ := http.Get(updatedURL)

    availability := "unavailable"
    if resp.StatusCode == 404 {
        availability = "available"
    }
    fmt.Println(username, availability, "at", updatedURL)
}
