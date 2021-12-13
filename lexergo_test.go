package lexergo

import (
	"fmt"
	"testing"
)

var blank_fn Bool_string_string_fn = Repeat_fn(Select_fn(MatchSpace_fn, MatchTab_fn), 1, MaxInt)

// alhpabet apha_digit repeat 0~MAX
var identifer_fn Bool_string_string_fn = Concat_fn(
	Match_alpha_fn,
	Repeat_fn(Select_fn(
		Match_alpha_fn, Match_digit_fn), 0, MaxInt),
)

func TestLexer(t *testing.T) {
	flag, read, rest := blank_fn("\t \t \t")
	if flag != true {
		t.Errorf("blank failed")
	} else {
		fmt.Println("match:", read)
		fmt.Println("rest:", rest)
	}
}

func TestIdentifer(t *testing.T) {
	flag, read, rest := identifer_fn("a12q123_")
	if flag != true {
		t.Errorf("identifer failed")
	} else {
		fmt.Println("match:", read)
		fmt.Println("rest:", rest)
	}
}
