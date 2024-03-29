package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type User struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

var user User
var PORT = ":1234"
var DATA = make(map[string]string)

func servingLog(r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	servingLog(r)
	w.WriteHeader(http.StatusNotFound)
	Body := "Thanks for visiting!\n"
	fmt.Fprintf(w, "%s", Body)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	servingLog(r)
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is: " + t + "\n"
	fmt.Fprintf(w, "%s", Body)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	servingLog(r)
	if r.Method != http.MethodPost {
		http.Error(w, "Error:", http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "%s\n", "Method not allowed!")
		return
	}
	d, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}
	log.Println(user)
	if user.Username != "" {
		DATA[user.Username] = user.Password
		log.Println(DATA)
		w.WriteHeader(http.StatusOK)
	} else {
		log.Println("user.Username is empty")
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	servingLog(r)
	if r.Method != http.MethodGet {
        http.Error(w, "Error:", http.StatusMethodNotAllowed)
        fmt.Fprintf(w, "%s\n", "Method not allowed!")
        return
    }
	d, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "ReadAll - Error", http.StatusBadRequest)
        return
    }
	err = json.Unmarshal(d, &user)
    if err != nil {
        log.Println(err)
        http.Error(w, "Unmarshal - Error", http.StatusBadRequest)
        return
    }
    fmt.Println(user)
	_, ok := DATA[user.Username]
    if ok && user.Username != "" {
        log.Println("Found!")
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "%s\n", d)
	} else {
		log.Println("Not found!")
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "Map - Resource not found!", http.StatusNotFound)
	}
	return
}

func main() {
	arguments := os.Args
    if len(arguments) != 1 {
        PORT = ":" + arguments[1]
    }
	mux := http.NewServeMux()
	s := &http.Server{
		Addr: PORT,
		Handler: mux,
		IdleTimeout: 10 * time.Second,
		ReadTimeout: time.Second,
		WriteTimeout: time.Second,
	}
	mux.Handle("/time", http.HandlerFunc(timeHandler))
	mux.Handle("/add", http.HandlerFunc(addHandler))
	mux.Handle("/get", http.HandlerFunc(getHandler))
	mux.Handle("/", http.HandlerFunc(defaultHandler))
	
	fmt.Println("Ready to serve at", PORT)
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}