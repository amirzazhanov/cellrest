package main

import (
	"gopkg.in/mgo.v2"
	//	"gopkg.in/mgo.v2/bson"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"net/http"
	"time"
)

type tomlConfig struct {
	BindHostname    string `toml:"hostname"`
	Port            string `toml:"port"`
	MongoHostname   string `toml:"mongo_hostname"`
	MongoDatabase   string `toml:"mongo_database"`
	MongoCollection string `toml:"mongo_collection"`
}

var config tomlConfig

func main() {
	var session *mgo.Session
	var err error
	if _, err = toml.DecodeFile("cellrest_config.toml", &config); err != nil {
		log.Print("[CRITICAL] ", "Problem parsing configuration file", err)
	}
	for {
		session, err = mgo.Dial(config.MongoHostname)
		if err != nil {
			log.Println("[CRITICAL] ", "Problem conecting database", err)

			log.Println("[CRITICAL] ", "Sleep 10 seconds", err)
			time.Sleep(10 * time.Second)
			continue
		} else {
			break
		}
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	ensureIndex(session)
	routes = append(routes,
		Route{"CellByType", "GET", "/cellstype/{cellType}", CellByType(session)},
		Route{"CellByID", "GET", "/cells/{cellMCC}/{cellNet}/{cellArea}/{cellID}", CellByID(session)},
	)
	r := NewRouter()

	log.Fatal(http.ListenAndServe(config.BindHostname+":"+config.Port, r))
}

func ensureIndex(s *mgo.Session) {
	session := s.Copy()
	defer session.Close()

	//	c := session.DB(config.MongoDatabase).C(config.MongoCollection)
	/*
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
	*/
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
