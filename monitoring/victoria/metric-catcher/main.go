package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Alert struct {
	Status string `json:"status"`
	Alerts []struct {
		Labels      map[string]string `json:"labels"`
		Annotations map[string]string `json:"annotations"`
	} `json:"alerts"`
}

func alertHandler(w http.ResponseWriter, r *http.Request) {
	var alert Alert
	if err := json.NewDecoder(r.Body).Decode(&alert); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Received alert: %+v\n", alert)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/on_alert", alertHandler)
	fmt.Println("Starting server on port 9090...")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
