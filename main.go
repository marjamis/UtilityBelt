package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "servicetesting/redis"
)

func handler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/"):]
  var filename string

  if len(title) == 0 {
    filename = "../templates/index.html"
  } else {
    filename = "../" + title
    w.Header().Set("Cache-Control", "max-age=86400")
  }

  body, err := ioutil.ReadFile(filename)
  if err != nil {
    http.Error(w, "404 Not Found", http.StatusNotFound)
  } else {
    fmt.Fprintf(w, "%s", body)
  }
  fmt.Printf("page: %s - IP: %s\n", filename, r.RemoteAddr)
  return
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/redis", redis.RedisHandler)
    fmt.Println("Listening on port 8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
      fmt.Println("Error: ", err)
    }
}
