package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"krysopath.it/deemarker/report"
)

func init() {
	flag.Parse()
}

func main() {
	r, err := report.ReadReports(flag.Arg(0))
	if err != nil {
		panic(err)
	}
	jsonBytes, err := json.Marshal(r)
	fmt.Fprintf(os.Stdout, "%s\n", string(jsonBytes))
}
