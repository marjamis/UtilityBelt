package redis

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/redis.v3"
)

type RedisItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type RedisCollection struct {
	RedisItems []RedisItem
}

func RedisConnection() (connection *redis.Client, err error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println("Redis connection test: ", pong, err)
	return client, err
}

func RedisHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Redis entry")
	action := r.URL.Query().Get("action")
	redisClient, err := RedisConnection()
	//Appears no err is generated on failure to connect need to test elsewise
	if err != nil {
		http.Error(w, "Error with connecting to Redis", http.StatusInternalServerError)
		return
	}

	var text RedisItem
	items := []RedisItem{}
	collection := RedisCollection{items}
	if action == "display" {
		data := redisClient.Keys("*")
		for i := 0; i < len(data.Val()); i++ {
			value, err := redisClient.Get((data.Val()[i])).Result()
			if err != nil {
				http.Error(w, "Error with getting key", http.StatusInternalServerError)
				return
			}
			text = RedisItem{data.Val()[i], value}
			collection.RedisItems = append(collection.RedisItems, text)
		}
		js, err := json.Marshal(collection)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", js)
		return
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
	fmt.Fprintf(w, "")

	fmt.Printf("redisAction: %s - IP: %s\n", action, r.RemoteAddr)
	redisClient.Close()
	return
}
