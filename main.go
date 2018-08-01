package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func main() {
    username := "username"
    file, err := os.Open("config.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      rawURL := scanner.Text()
      updatedURL := strings.Replace(rawURL, "{{ USERNAME }}", username, -1)
      fmt.Println(updatedURL)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
