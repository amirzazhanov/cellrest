package main

import (
//	"time"
)

/* Cell cell description
radio,mcc,net,area,cell,unit,lon,lat,range,samples,changeable,created,updated,averageSignal
UMTS,262,2,801,86355,,13.28527,52.521711,37,7,1,1282569574,1300175362,-91
GSM,262,2,801,1795,,13.276605,52.525348,5714,9,1,1282569574,1300175362,-87
GSM,262,2,801,1794,,13.284916,52.523771,6278,13,1,1282569574,1300816026,-91
UMTS,262,2,801,211250,,13.285081,52.521622,33,3,1,1282569574,1299486792,-94
*/
type Cell struct {
	Radio         string  `json:"radio"`
	MCC           uint32  `json:"mcc"`
	Net           uint32  `json:"net"`
	Area          uint32  `json:"area"`
	Cell          uint32  `json:"cell"`
	Unit          uint32  `json:"unit"`
	Lon           float32 `json:"lon"`
	Lat           float32 `json:"lat"`
	Range         uint32  `json:"range"`
	Samples       uint32  `json:"samples"`
	Changeable    bool    `json:"changeable"`
	Created       uint32  `json:"created"`
	Updated       uint32  `json:"updated"`
	AverageSignal uint32  `json:"averageSygnal"`
}

// Cells defines slice of Cell
type Cells []Cell
