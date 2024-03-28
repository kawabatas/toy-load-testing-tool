package toyloadtestingtool

import (
	"bytes"
	"testing"
)

func TestReadTargets(t *testing.T) {
	lines := bytes.NewBufferString("GET http://localhost:3000/\n\n      // HEAD http://localhost:3000 this is a comment \nHEAD http://localhost:3000/\n")
	targets, err := readTargets(lines)
	if err != nil {
		t.Fatalf("Couldn't parse valid source: %s", err)
	}
	for i, method := range []string{"GET", "HEAD"} {
		if targets[i].Method != method ||
			targets[i].URL.String() != "http://localhost:3000/" {
			t.Fatalf("Request was parsed incorrectly. Got: %s %s",
				targets[i].Method, targets[i].URL.String())
		}
	}
}

func TestNewTargets(t *testing.T) {
	lines := []string{"GET http://localhost:3000/", "HEAD http://localhost:3000/"}
	targets, err := NewTargets(lines)
	if err != nil {
		t.Fatalf("Couldn't parse valid source: %s", err)
	}
	for i, method := range []string{"GET", "HEAD"} {
		if targets[i].Method != method ||
			targets[i].URL.String() != "http://localhost:3000/" {
			t.Fatalf("Request was parsed incorrectly. Got: %s %s",
				targets[i].Method, targets[i].URL.String())
		}
	}
}
