package timesheet

import (
	"fmt"
	"strings"
	"text/scanner"
)

func Example() {
	report := make(chan part)
	var s scanner.Scanner
	s.Init(strings.NewReader("2018 August"))
	l := &lexer{
		report: report,
		stream: s,
	}
	// print all parts as they come in
	go func() {
		for {
			part, more := <-report
			if !more {
				return
			}
			fmt.Println(part)
		}
	}()
	l.lex()
	// output:
	// Year: 2018
	// August
}