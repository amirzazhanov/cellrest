package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Page structure simple
type Page struct {
	Title string
	Body  []byte
}

// Index function
func Index(w http.ResponseWriter, r *http.Request) {
	pTmp := &Page{Title: "TestPage", Body: []byte("This is a sample simple Page.")}
	t, _ := template.ParseFiles("index-tpl.html")
	t.Execute(w, pTmp)
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

		var cells []Cell
		err := c.Find(bson.M{"radio": cellType}).All(&cells)
		if err != nil {
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed find cell: ", err)
			return
		}
		for _, everycell := range cells {
			if everycell.Radio == "" {
				ErrorWithJSON(w, "CellType not found", http.StatusNotFound)
				return
			}
		}

		respBody, err := json.MarshalIndent(cells, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		ResponseWithJSON(w, respBody, http.StatusOK)
	}
}
func CellByID(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()
		vars := mux.Vars(r)
		cellMCC, _ := strconv.ParseUint(vars["cellMCC"], 10, 32)
		cellNet, _ := strconv.ParseUint(vars["cellNet"], 10, 32)
		cellArea, _ := strconv.ParseUint(vars["cellArea"], 10, 32)
		cellID, _ := strconv.ParseUint(vars["cellID"], 10, 32)

		log.Println("CellMCC: ", cellMCC, "CellNet: ", cellNet, "CellArea: ", cellArea, "CellID: ", cellID)
		c := session.DB("cells").C("cells")

		var cells []Cell
		err := c.Find(bson.M{"mcc": cellMCC, "net": cellNet, "area": cellArea, "cell": cellID}).All(&cells)
		if err != nil {
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed find cell: ", err)
			return
		}
		for _, everycell := range cells {
			if everycell.MCC != uint32(cellMCC) {
				ErrorWithJSON(w, "Cell not found", http.StatusNotFound)
				return
			}
		}

		respBody, err := json.MarshalIndent(cells, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		ResponseWithJSON(w, respBody, http.StatusOK)
	}
}
