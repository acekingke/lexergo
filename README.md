# Introduction
Jamie Zawinski:
> Some people, when confronted with a problem, think "I know, I'll use regular expressions." Now they have two problems.

It's so many regular expression forms are difficult to understand, such as `perl, python, grep awk`

[Rob Pike's talk](
https://www.youtube.com/watch?v=HxaD_trXwRE)
also point out some disadvantage about regular expressions, He said it more easy write it in go for pragrammer than use regular expression.

Regular expression in Math definition:

* a single charater is regular expression
* **concat**: if R1 ,R2 are regular expression, R1 concat R2 are also regular expression
* **select**: if R1 ,R2 are regular expression, R1 | R2 are also regular expression
* **repeat :** R1 repeat concat R1 ,the result is also regular expression
This project rewrite  [regular-expression-in-hy](
https://github.com/acekingke/regular-expression-in-hy)  in go . It just offer three type basic function, `Concat_fn, Select_fn, Repeate_fn` and return tuple like 
`matched, readString, restString` ,matched is show whether string matched the pattern, readString show that string has matched, rest part is restString.

## Samples
```go
var MatchSpace_fn Bool_string_string_fn = func(x string) (bool, string, string) {
	return Match_fn(x, func(in string) bool {
		return in == " "
	})
} // MatchSpace_fn

var MatchTab_fn = func(x string) (bool, string, string) {
	return Match_fn(x, func(in string) bool {
		return in == "\t"
	})
} // match_tab

var Match_alpha_fn = func(x string) (bool, string, string) {
	return Match_fn(x, func(in string) bool {
		return in >= "a" && in <= "z" || in >= "A" && in <= "Z"
	})
}
var Match_digit_fn = func(x string) (bool, string, string) {
	return Match_fn(x, func(in string) bool {
		return in >= "0" && in <= "9"
	})
}

var blank_fn Bool_string_string_fn = Repeat_fn(Select_fn(MatchSpace_fn, MatchTab_fn), 1, MaxInt)

// alhpabet apha_digit repeat 0~MAX
var identifer_fn Bool_string_string_fn = Concat_fn(
	Match_alpha_fn,
	Repeat_fn(Select_fn(
		Match_alpha_fn, Match_digit_fn), 0, MaxInt),
)
```
Useage:
```
flag, read, rest := identifer_fn("a12q123_")	if flag != true {
		t.Errorf("identifer failed")
} else {
	fmt.Println("match:", read)
	fmt.Println("rest:", rest)
}
```

## LICENSE
MIT
