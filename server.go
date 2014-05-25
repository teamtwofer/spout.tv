package main

import (
  "fmt"
  "net/http"
  "spouttv"
)

func main() {
  fmt.Println("Hello, world")
  http.HandleFunc("/", handleIndex)
  user := &spouttv.User{}
  user.Name = "potato"
  fmt.Println(user.Name)
}

func handleIndex(writer http.ResponseWriter, request *http.Request) {
  writer.Header().Set("Content-Type", "text/html")
}