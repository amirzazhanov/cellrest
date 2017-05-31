package main

import (
	"fmt"
	"github.com/gorilla/mux"
	//	"html"
	"log"
	"net/http"
)

/*
radio,mcc,net,area,cell,unit,lon,lat,range,samples,changeable,created,updated,averageSignal
UMTS,262,2,801,86355,,13.28527,52.521711,37,7,1,1282569574,1300175362,-91
GSM,262,2,801,1795,,13.276605,52.525348,5714,9,1,1282569574,1300175362,-87
GSM,262,2,801,1794,,13.284916,52.523771,6278,13,1,1282569574,1300816026,-91
UMTS,262,2,801,211250,,13.285081,52.521622,33,3,1,1282569574,1299486792,-94
*/
type Cell struct {
	Radio         string
	MCC           uint32
	Net           uint32
	Area          uint32
	Cell          uint32
	Unit          uint32
	Lon           float32
	Lat           float32
	Range         uint32
	Samples       uint32
	Changeable    bool
	Created       uint32
	Updated       uint32
	averageSignal uint32
}

type Cells []Cell

func main() {
	r := mux.NewRouter().StrictSlash(true)
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
