package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "strconv"
  "time"

  "github.com/marjamis/UtilityBelt/kubernetes"
  "github.com/marjamis/UtilityBelt/redis"
  "github.com/prometheus/client_golang/prometheus/promhttp"
)

func handler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/"):]
  var filename string

  if len(title) == 0 {
    filename = "./static/templates/index.html"
  } else {
    filename = "./" + title
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

func delay(w http.ResponseWriter, r *http.Request) {
  di, _ := strconv.Atoi(r.URL.Path[len("/delay/"):])
  time.Sleep(time.Duration(di) * time.Second)
  output := "Delayed " + r.URL.Path[len("/delay/"):] + " seconds and continued."
  fmt.Println(output)
  fmt.Fprintf(w, "%s", output)
  return
}

func main() {
  http.HandleFunc("/", handler)
  http.HandleFunc("/delay/", delay)
  http.Handle("/metrics", promhttp.Handler())
  http.HandleFunc("/redis", redis.RedisHandler)
  http.HandleFunc("/kubernetes", kubernetes.Handler)
  fmt.Println("Listening on port 8080...")
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    fmt.Println("Error: ", err)
  }
}
