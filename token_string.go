// generated by stringer -type Token token.go; DO NOT EDIT

package timesheet

import "fmt"

const _Token_name = "UndefinedErrorYearHoursNoteMonthSeparatorDayDateHourLeftParenthesisRightParenthesisOperatorColonMinutesTagWeek"

var _Token_index = [...]uint8{0, 9, 14, 18, 23, 27, 32, 41, 44, 48, 52, 67, 83, 91, 96, 103, 106, 110}

func (i Token) String() string {
	if i < 0 || i >= Token(len(_Token_index)-1) {
		return fmt.Sprintf("Token(%d)", i)
	}
	return _Token_name[_Token_index[i]:_Token_index[i+1]]
}
