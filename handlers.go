package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
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
	cellType := vars["cellType"]
	fmt.Fprintln(w, "Cell show:", cellType)
}

func CellByType(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()
		vars := mux.Vars(r)
		cellType := vars["cellType"]
		log.Println("CellType: ", cellType)
		c := session.DB("cells").C("cells")

		var cell Cell
		err := c.Find(bson.M{"radio": cellType}).One(&cell)
		if err != nil {
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed find cell: ", err)
			return
		}

		if cell.Radio == "" {
			ErrorWithJSON(w, "CellType not found", http.StatusNotFound)
			return
		}

		respBody, err := json.MarshalIndent(cell, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		ResponseWithJSON(w, respBody, http.StatusOK)
	}
}
