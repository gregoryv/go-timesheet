// Copyright (c) 2019 Gregory Vinčić. All rights reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package parser

import (
	"testing"
)

func TestLexer_Run(t *testing.T) {
	l := NewLexer("2018")
	out := l.Run()
	cases := []struct {
		part Part
		exp  string
	}{
		{<-out, "Year[1,1]: \"2018\""},
	}
	for _, c := range cases {
		got := c.part.String()
		exp := c.exp
		if got != exp {
			t.Errorf("%q, expected\n%q", got, exp)
		}
	}
}

func TestNewLexer(t *testing.T) {
	if l := NewLexer(""); l == nil {
		t.Fail()
	}
}
