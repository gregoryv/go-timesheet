package timesheet

import (
	"fmt"
)

type Part struct {
	Tok Token
	Val string
	Pos Position
}

func (p *Part) Errorf(format string, args ...interface{}) error {
	p.Val = fmt.Sprintf(format, args...)
	p.Tok = Error
	return fmt.Errorf(p.Val)
}

func (p *Part) String() string {
	return fmt.Sprintf("%s[%s]: %q", p.Tok, p.Pos.String(), p.Val)
}

func NewPart() *Part {
	return &Part{}
}
