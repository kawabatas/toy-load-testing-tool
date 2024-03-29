# toy-load-testing-tool
(A toy) HTTP load testing tool

This is a repository to learn how a HTTP load testing tool works.
I heavily referenced a simple Go-based HTTP load testing tool, [Vegeta](https://github.com/tsenart/vegeta).

# How to use

Build the binary
```
make
```

Print usage
```
./nappa

Usage of ./nappa:
  -cpus int
    	Number of CPUs to use (default 4)
  -duration duration
    	Duration of the test (default 5s)
  -output string
    	Reporter output file (default "stdout")
  -rate uint
    	Requests per second (default 1)
  -targets string
    	Targets file (default "testdata/targets-example-dot-com.txt")
```

Run a load test
```
./nappa -rate 2 -duration 3s -targets testdata/targets-example-dot-com.txt
```

# References
- [tsenart/vegeta](https://github.com/tsenart/vegeta)
