package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Index function
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// CellIndex - list all cells
func CellIndex(w http.ResponseWriter, r *http.Request) {
	cells := Cells{
		Cell{Radio: "UMTS"},
		Cell{Radio: "GSM"},
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cells); err != nil {
		panic(err)
	}
}

//CellShow - show specific cell
func CellShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cellID := vars["cellID"]
	fmt.Fprintln(w, "Cell show:", cellID)
}
