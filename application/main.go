package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "encoding/json"
  "gopkg.in/redis.v3"
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

func RedisConnection() (connection *redis.Client, err error){
  client := redis.NewClient(&redis.Options{
    Addr: "172.17.0.3:6379",
    Password: "",
    DB: 0,
  })

  pong, err := client.Ping().Result()
  fmt.Println("Redis connection test: ", pong, err)
  return client, err
}


func redisHandler(w http.ResponseWriter, r *http.Request) {
  action := r.URL.Query().Get("action")
  redisClient,_ := RedisConnection()
  //var text RedisItem
  items := []RedisItem{}
  collection := RedisCollection{items}
  if action == "display" {
    data := redisClient.Keys("*")
    fmt.Println(data)
  } else if action == "add" {
    key := r.URL.Query().Get("key")
    value := r.URL.Query().Get("value")
    err := redisClient.Set(key, value, 0).Err()
    if err != nil {
      http.Error(w, "Error with adding key", http.StatusInternalServerError)
      return
    }
  } else if action == "del" {
    key := r.URL.Query().Get("key")
    err := redisClient.Del(key).Err()
    if err != nil {
      http.Error(w, "Error with deleting key", http.StatusInternalServerError)
      return
    }
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
  fmt.Printf("redisAction: %s - IP: %s\n", action, r.RemoteAddr)
  redisClient.Close()
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
