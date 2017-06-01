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

	if err := json.NewEncoder(w).Encode(cells); err != nil {
		panic(err)
	}
}

//CellShow - show specific cell
func CellShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cellID := vars["cellId"]
	fmt.Fprintln(w, "Cell show:", cellID)
}
