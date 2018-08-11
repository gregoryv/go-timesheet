package timesheet

import (
	"fmt"
	"testing"
)

type TestCase struct {
	msg       string
	pos       Position
	line, col int
}

func p(line, col int) Position {
	p := NewPosition()
	p.line = line
	p.col = col
	return *p
}

func TestPosition_Val(t *testing.T) {
	p := NewPosition()
	line, col := p.Val()
	assert(t, "Should match", compareLineCol(line, 1, col, 1))
	p.Next()
	line, col = p.Val()
	assert(t, "Should match", compareLineCol(line, 1, col, 2))
}

func TestPosition_Back(t *testing.T) {
	// Case when a position is moved back over a new line
	special := p(3, 3)
	special.NextLine()

	cases := []TestCase{
		{"Stay on first, when already there", p(1, 1), 1, 1},
		{"Backup line", p(2, 1), 1, 1},
		{"Only backup column", p(1, 5), 1, 4},
		{"Last column should be remembered", special, 3, 3},
	}
	for _, c := range cases {
		c.msg += ", from " + c.pos.String()
		line, col := c.pos.Back()
		assert(t, c.msg, compareLineCol(c.line, line, c.col, col))
	}

	err := catchPanic(func() {
		pos := p(2, 1)
		pos.NextLine()
		pos.Back() // ok
		pos.Back() // not ok since we have no history left
	})
	if err == nil {
		t.Error("Expected 2 x Back to panic")
	}
}

func TestPosition_NextLine(t *testing.T) {
	cases := []struct {
		msg       string
		pos       Position
		line, col int
	}{
		{"", p(1, 1), 2, 1},
		{"Reset column when moving to next line", p(1, 5), 2, 1},
	}
	for _, c := range cases {
		line, col := c.pos.NextLine()
		assert(t, c.msg, compareLineCol(c.line, line, c.col, col))
	}
}

func TestPosition_Next(t *testing.T) {
	cases := []struct {
		msg       string
		pos       Position
		line, col int
	}{
		{"Advance column by 1", p(1, 1), 1, 2},
		{"", p(3, 5), 3, 6},
	}
	for _, c := range cases {
		line, col := c.pos.Next()
		assert(t, c.msg, compareLineCol(c.line, line, c.col, col))
	}
}

func assert(t *testing.T, msg string, errors ...error) {
	t.Helper()
	for _, err := range errors {
		if err != nil {
			t.Errorf("%s: %s", msg, err)
		}
	}
}

func compareLineCol(expLine, line, expCol, col int) (err error) {
	switch {
	case expLine != line:
		err = fmt.Errorf("expected line %v, got %v", expLine, line)
	case expCol != col:
		err = fmt.Errorf("expected col %v, got %v", expCol, col)
	}
	return
}

func ExamplePosition_String() {
	pos := NewPosition()
	fmt.Println(pos)
	//output:
	//1,1
}

func TestNewPosition(t *testing.T) {
	if pos := NewPosition(); pos == nil {
		t.Fail()
	}
}

func catchPanic(fn func()) (err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = fmt.Errorf("%s", err)
		}
	}()
	fn()
	return
}
