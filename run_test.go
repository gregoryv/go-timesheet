package timesheet

import (
	. "github.com/gregoryv/qual"
	"testing"
)

func TestLexer_run(t *testing.T) {
	for _, c := range []struct {
		start   lexFn
		input   string
		exp     *Part
		hasNext bool
	}{ /*
			{1, lexOperator, " ", Error.Is("invalid Operator")},
			{1, lexOperator, "+", Operator.Is("+")},
			{1, lexOperator, "-", Operator.Is("-")},
			{2, lexLeftParenthesis, "(-", Operator.Is("-", Position{1, 2})},
			{1, lexLeftParenthesis, "  (",
				Error.Is("invalid LeftParenthesis", Position{1, 1}),
			},
			{1, lexLeftParenthesis, "(", LeftParenthesis.Is("(")},
			{2, lexReported, "6\n39", Number.Is("39", Position{2, 1})},
			{2, lexReported, "6 (", LeftParenthesis.Is("(", Position{1, 3})},
			{1, lexReported, "\n6", Number.Is("6", Position{2, 1})},
			{1, lexReported, "  \n   6", Number.Is("6", Position{2, 4})}, // date number
			{1, lexReported, "  \n5", Number.Is("5", Position{2, 1})},    // week number
			{2, lexDay, "Mon 8", Number.Is("8", Position{1, 5})},
			{1, lexDay, "Mo", Error.Is("invalid Day")},
			{1, lexDay, "mon", Error.Is("invalid Day")},
			{1, lexDay, "Mon", Day.Is("Mon")},
			{1, lexDate, " 4", Error.Is("invalid Number")},
			{1, lexDate, "4", Number.Is("4")},
			{2, lexWeek, "26   1", Number.Is("1", Position{1, 6})},
			{1, lexWeek, "     2", Number.Is("2", Position{1, 6})},
			{1, lexWeek, "jkl", Error.Is("invalid Number")},
			{1, lexWeek, "26", Number.Is("26")},
			{1, lexYear, "2018", Number.Is("2018")},
			{1, lexYear, "not a year", Error.Is("invalid Number")},*/
		{lexSep, "-----", Separator.Is("-----"), true},
		{lexMonth, "April  \n---\n11", Month.Is("April"), true},
		{lexMonth, "August\n", Month.Is("August"), true},
		{lexMonth, "not a month", Error.Is("invalid month"), true},
		{lexMonth, "August something more",
			Error.Is("expect newline", Position{1, 7}), true,
		},
		{lexMonth, "Augusty", Error.Is("invalid month"), true},
		{lexMonth, "august", Error.Is("invalid month"), true},
		{lexMonth, " August", Error.Is("invalid month"), true},
	} {
		input, exp := c.input, c.exp
		got, next := c.start(NewLexer(c.input).scanner)
		Assert(t, Vars{input, exp, got},
			got.Equals(exp),
			c.hasNext == (next != nil),
		)
	}
}

func skipParts(i int, out chan Part) (p Part) {
	for j := 0; j < i; j++ {
		p = <-out
	}
	return
}

func TestScanPart(t *testing.T) {
	cases := []struct {
		msg, txt string
		tok      Token
	}{
		{"", "1234", Number},
		{"", "as1234", Error},
	}
	for _, c := range cases {
		s := NewScanner(c.txt)
		got := ScanPart(s, Number)
		Assert(t, Vars{c, got}, c.tok == got.Tok)
	}

}
