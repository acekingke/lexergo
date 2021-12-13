package lexergo

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
