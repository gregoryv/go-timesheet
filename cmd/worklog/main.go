package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	timesheet "github.com/gregoryv/go-timesheet"
)

func main() {
	flag.Parse()

	filePaths := flag.Args()
	if len(filePaths) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	p := timesheet.NewParser()
	for _, path := range filePaths {
		body, err := ioutil.ReadFile(path)
		fatal(err)

		reported, err := p.SumReported(body)
		fatal(err)
		fmt.Println(reported)
	}
}

func fatal(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
