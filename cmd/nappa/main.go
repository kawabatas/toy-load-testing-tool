package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"time"

	nappa "github.com/kawabatas/toy-load-testing-tool"
)

var (
	rate     uint64
	targetsf string
	duration time.Duration
	output   string
	cpus     int
)

func init() {
	flag.Uint64Var(&rate, "rate", 1, "Requests per second")
	flag.StringVar(&targetsf, "targets", "testdata/targets-example-dot-com.txt", "Targets file")
	flag.DurationVar(&duration, "duration", 5*time.Second, "Duration of the test")
	flag.StringVar(&output, "output", "stdout", "Reporter output file")
	flag.IntVar(&cpus, "cpus", runtime.NumCPU(), "Number of CPUs to use")
}

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}

	runtime.GOMAXPROCS(cpus)

	if err := run(rate, duration, targetsf, output); err != nil {
		log.Fatal(err)
	}
}

// run is an utility function that validates the attack arguments, sets up the
// required resources, launches the attack and reports the results
func run(rate uint64, duration time.Duration, targetsf, output string) error {
	if rate == 0 {
		return errors.New("error rate can't be zero")
	}

	if duration == 0 {
		return errors.New("error duration can't be zero")
	}

	var out io.Writer
	switch output {
	case "stdout":
		out = os.Stdout
	default:
		file, err := os.Create(output)
		if err != nil {
			return fmt.Errorf("error output file (%s): %s", output, err)
		}
		defer file.Close()
		out = file
	}

	targets, err := nappa.NewTargetsFromFile(targetsf)
	if err != nil {
		return fmt.Errorf("error targets file (%s): %s", targetsf, err)
	}

	log.Printf("Nappa is attacking %d targets for %s...\n", len(targets), duration)
	results := nappa.Attack(targets, rate, duration)
	log.Println("Done!")
	log.Printf("Writing report to '%s'...", output)
	if err = nappa.ReportText(results, out); err != nil {
		return fmt.Errorf("error reporting %s", err)
	}
	return nil

}
