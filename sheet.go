package timesheet

import (
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Sheet struct {
	Period   string // Year Month
	Reported Tagged
	Tags     []Tagged
}

func NewSheet() *Sheet {
	return &Sheet{Reported: Tagged{0, ""}}
}

func Load(filepath string) (sheet *Sheet, err error) {
	p := NewParser()
	body, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}
	return p.Parse(body)
}

func Render(w io.Writer, year int, month time.Month, hours int) {
	period := fmt.Sprintf("%v %v", year, time.Month(month))
	fmt.Fprintf(w, "%s\n", period)
	fmt.Fprint(w, strings.Repeat("-", len(period)), "\n")

	tmp := time.Date(year, time.Month(month), 1, 23, 0, 0, 0, time.UTC)
	var lastWeek int
	for month := tmp.Month(); month == tmp.Month(); tmp = tmp.Add(24 * time.Hour) {
		_, week := tmp.ISOWeek()
		if lastWeek != week {
			fmt.Fprintf(w, "%2v ", week)
			lastWeek = week
		} else {
			fmt.Fprint(w, "   ")
		}

		fmt.Fprintf(w, "%+2v %3s", tmp.Day(), tmp.Weekday().String()[:3])
		switch tmp.Weekday() {
		case 0, 6:
		default:
			fmt.Fprint(w, " ", hours)
		}
		fmt.Fprint(w, "\n")
	}
}

func (par *Parser) Parse(body []byte) (sheet *Sheet, err error) {
	sheet = NewSheet()
	lex := NewLexer(string(body))
	out := lex.Run()
	tagDur := make(map[string]time.Duration, 0)
	var dur time.Duration // for tags
	operator := 1         // +1 or -1
	tagged := make([]Tagged, 0)
	inTag := false
	for {
		p, more := <-out
		switch p.Tok {
		case LeftParenthesis, RightParenthesis:
			inTag = !inTag
		case Year:
			sheet.Period += p.Val
		case Month:
			sheet.Period += " " + p.Val
		case Operator:
			if p.Val == "-" {
				operator = -1
			}
		case Tag:
			if _, exists := tagDur[p.Val]; !exists {
				tagDur[p.Val] = 0
			}
			tagDur[p.Val] += dur
			dur = 0
			operator = 1
		case Hours:
			h, _ := strconv.Atoi(p.Val)
			hh := time.Duration(h*operator) * time.Hour
			if inTag {
				dur += hh
			} else {
				sheet.Reported.Duration += hh
			}
		case Minutes:
			m, _ := strconv.Atoi(p.Val)
			mm := time.Duration(m*operator) * time.Minute
			if inTag {
				dur += mm
			} else {
				sheet.Reported.Duration += mm
			}
		case Error:
			err = fmt.Errorf("%s", p)
		}
		if !more || err != nil {
			break
		}
	}
	for tag, dur := range tagDur {
		tagged = append(tagged, Tagged{dur, tag})
	}
	sort.Sort(byTag(tagged))
	sheet.Tags = tagged
	return
}

func (sheet *Sheet) String() string {
	return fmt.Sprintf("%-14s %7s %s", sheet.Period, sheet.Reported,
		strings.Join(inParenthesis(sheet.Tags), " "))
}

func inParenthesis(tags []Tagged) []string {
	res := make([]string, 0)
	for _, tag := range tags {
		res = append(res, fmt.Sprintf("(%s)", tag))
	}
	return res
}
