package main

import (
    "fmt"
    "log"
    "net/http"
    "sync"
)

var mu sync.Mutex

func get_users(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	fmt.Fprintf(w, GetUsers(w, r))
	mu.Unlock()
}

func get_user_by_id(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	fmt.Fprintf(w, GetUserByID(w, r))
	mu.Unlock()
}

func get_interests(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	fmt.Fprintf(w, GetInterests(w, r))
	mu.Unlock()
}

func get_users_interests(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	fmt.Fprintf(w, GetUsersInterests(w, r))
	mu.Unlock()
}

func main() {
	http.HandleFunc("/get_users/", get_users)
	http.HandleFunc("/get_user_by_id/", get_user_by_id)
	http.HandleFunc("/get_interests/", get_interests)
	http.HandleFunc("/get_users_interests/", get_users_interests)
	log.Fatal(http.ListenAndServe(":80", nil))
}
