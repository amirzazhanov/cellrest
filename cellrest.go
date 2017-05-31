package main

import (
	"fmt"
	"github.com/gorilla/mux"
	//	"html"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter() //.StrictSlash(true)
	r.HandleFunc("/", Index)
	r.HandleFunc("/cells", CellIndex)
	r.HandleFunc("/cells/{cellId}", CellShow)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

// Index function
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// CellIndex - list all cells
func CellIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo Index!")
}

//CellShow - show specific cell
func CellShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cellID := vars["cellId"]
	fmt.Fprintln(w, "Cell show:", cellID)
}
