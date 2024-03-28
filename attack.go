package toyloadtestingtool

import "time"

// Result represents the metrics we want out of an http.Response
type Result struct {
	Code      uint16
	Timestamp time.Time
	Timing    time.Duration
	BytesOut  uint64
	BytesIn   uint64
	Error     error
}

// Results is a slice of Result defined only to be sortable with sort.Interface
type Results []Result

func (r Results) Len() int           { return len(r) }
func (r Results) Less(i, j int) bool { return r[i].Timestamp.Before(r[j].Timestamp) }
func (r Results) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
