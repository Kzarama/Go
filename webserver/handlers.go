package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func HandleRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello world!!")
}

func HandleHome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the API endpoint")
}

func PostRequest(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var metadata MetaData
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	fmt.Fprintf(w, "Payload %v\n", metadata)
}

func UserPostRequest(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}