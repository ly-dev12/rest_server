package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	ARRAY_USER []User
)

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"Email"`
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Add("Content-Type", "application/json")

	for _, value := range ARRAY_USER {
		if value.Id == string(vars["id"]) {
			buf, _ := json.Marshal(value)
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(buf))
			return
		}
	}

	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, "Not exist user with this ID")
}

func usuarioHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	list := "dfEWTsnm9qw9kfFGds343097RE69DGDF59gjdfpglsdfk456so632GWE2aprop767upoifghohdfJYHTHGJloDpwewGRqGfLOUdSGDGWQog56hfoGFo3ws"
	code := ""

	go func() {
		for i := 0; i < 13; i++ {
			random := rand.Intn(len(list) - 1)
			code += string(list[random])
		}
		fmt.Println("Codigo: ", code)
	}()

	time.Sleep(1 * time.Millisecond)
	newUser := &User{Id: code, Name: "Anonimus", Age: 12, Email: "bartsdfg@hotmail.com"}

	err := json.NewDecoder(r.Body).Decode(newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error in get body")
		return
	}
	ARRAY_USER = append(ARRAY_USER, *newUser)
	buf, _ := json.Marshal(ARRAY_USER)
	//fmt.Println(ARRAY_USER, string(buf))
	fmt.Fprintf(w, "Person: %+v", string(buf))
}

func main() {

	mux := mux.NewRouter()

	mux.HandleFunc("/users/{id}", usersHandler).Methods("GET")
	mux.HandleFunc("/users/add", usuarioHandle).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", mux))
}
