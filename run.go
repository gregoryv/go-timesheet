package timesheet

import (
	"strings"
)

type lexFn func(s *Scanner, out chan Part) lexFn

const validMonths = "JanuaryFebruaryMarchAprilMayJune" +
	"JulyAugustSeptemberOctoberNovemberDecember"

func lexMonth(s *Scanner, out chan Part) lexFn {
	p := Part{Tok: Month, Pos: s.Pos()}
	val, ok := s.Scan("JFMASOND")
	if !ok {
		p.Errorf("invalid month")
	} else {
		rest, _ := s.ScanAll("abcdefghijklmnopqrstuvxyz")
		p.Val = val + rest
		if !strings.Contains(validMonths, p.Val) {
			p.Errorf("invalid month")
		}
	}
	out <- p
	s.Scan(" \n")
	return nil
}

func lexYear(s *Scanner, out chan Part) lexFn {
	pos := s.Pos()
	val, ok := s.ScanAll("0123456789")
	p := Part{Tok: Number, Val: val, Pos: pos}
	if !ok {
		p.Errorf("invalid year")
	}
	out <- p
	s.Scan(" ")
	return nil
}

func (l *Lexer) run(s *Scanner, out chan Part) {
	// We expect to start the file with a year
	for fn := lexYear; fn != nil; fn = fn(s, out) {
	}
	close(out)
}
