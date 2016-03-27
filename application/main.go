package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "encoding/json"
)


type RedisItem struct {
  Key	string	`json:"key"`
  Value	string	`json:"value"`
}

type RedisCollection struct {
  RedisItems []RedisItem
}

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
    fmt.Fprintf(w, "%s", "No such page")
  } else {
    fmt.Fprintf(w, "%s", body)
  }
  fmt.Printf("page: %s - IP: %s\n", filename, r.RemoteAddr)
  return
}

func redisHandler(w http.ResponseWriter, r *http.Request) {
  value := r.URL.Query().Get("action")

  var text RedisItem
  items := []RedisItem{}
  collection := RedisCollection{items}
  if value == "display" {
    text = RedisItem{"remoteKey", "remoteData"}
    collection.RedisItems = append(collection.RedisItems, text)
    text = RedisItem{"remoteKey2", "remoteData2"}
    collection.RedisItems = append(collection.RedisItems, text)
  } else if value == "add" {
    text = RedisItem{"remoteKey1", "remoteData"}
  } else if value == "del" {
    text = RedisItem{"remoteKey2", "remoteData"}
  } else {
    http.Error(w, "Unknown Error", http.StatusInternalServerError)
    return
  }

  js,err := json.Marshal(collection)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  fmt.Fprintf(w, "%s", js)
  fmt.Printf("redisAction: %s - IP: %s\n", value, r.RemoteAddr)
  return
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/redis", redisHandler)
    fmt.Println("Listening on port 8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
      fmt.Println("Error: ", err)
    }
}
