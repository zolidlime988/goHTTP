package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type FullName struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("Server is listening on port 3000")
	http.ListenAndServe(":3000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received response type :", r.Method)
	var name FullName
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error : " + err.Error()))
	}
	json.Unmarshal(body, &name)
	w.Header().Add("Content-Type", "application/json")
	res, err := json.Marshal(name)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error : " + err.Error()))
	}
	w.Write([]byte(res))
}
