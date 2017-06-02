package main

import (
	"gopkg.in/mgo.v2"
	//	"gopkg.in/mgo.v2/bson"
	"fmt"
	"log"
	"net/http"
)

func main() {
	session, err := mgo.Dial("192.168.10.2")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	ensureIndex(session)
	routes = append(routes, Route{
		"CellByType",
		"GET",
		"/cellstype/{cellType}",
		CellByType(session),
	})
	r := NewRouter()

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func ensureIndex(s *mgo.Session) {
	session := s.Copy()
	defer session.Close()

	c := session.DB("cells").C("cells")

	index := mgo.Index{
		Key: []string{"mcc"},
		//		Unique:     true,
		//		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}
func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}
