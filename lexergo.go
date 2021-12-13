package lexergo

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

// rewrite the
// https://github.com/acekingke/regular-expression-in-hy
// to the golang

// function first
func first_fn(in string) string {
	if len(in) == 0 {
		return ""
	}
	return string(in[0])
}

// function rest
func rest_fn(in string) string {
	if len(in) < 1 {
		return ""
	}
	return in[1:]
}

type bool_fn func(in string) bool
type Bool_string_string_fn func(in string) (bool, string, string)

// function match fuctions
// return matched match_string, rest_string
func Match_fn(in string, fn bool_fn) (bool, string, string) {
	if fn(first_fn(in)) {
		return true, first_fn(in), rest_fn(in)
	} else {
		return false, "", in
	}
}

// defeine concat
func Concat_fn(args ...Bool_string_string_fn) Bool_string_string_fn {
	return func(in string) (bool, string, string) {
		count := 0
		length := len(args)
		rest := in
		matched := false
		s := ""
		readStr := ""
		for _, fn := range args {
			if rest == "" {
				return false, "", ""
			} else {
				matched, s, rest = fn(rest)
				if matched {
					count++
					readStr += s
				}
				if count == length {
					return true, readStr, rest
				}
			}
		}
		return false, "", in
	}

}

// function select
func Select_fn(args ...Bool_string_string_fn) Bool_string_string_fn {
	return func(in string) (bool, string, string) {
		rest := in
		matched := false
		s := ""
		readStr := ""
		ret := false
		for _, fn := range args {
			matched, s, rest = fn(rest)
			if matched {
				ret = true
				readStr += s
				break
			}
		}
		return ret, readStr, rest
	}
}

// Repeat
func Repeat_fn(arg Bool_string_string_fn,
	min int, max int) Bool_string_string_fn {

	return func(in string) (bool, string, string) {
		rest := in
		count := 0
		matched := true
		s := ""
		readStr := ""
		for {
			if rest == "" {
				return false, "", ""
			} else {
				matched, s, rest = arg(rest)
				if matched {
					count++
					readStr += s
				}
				// if not flag or not rest:
				// break
				if !matched || rest == "" {
					break
				}
			}
		}
		if count >= min && count <= max {
			return true, readStr, rest
		} else {
			return false, readStr, rest
		}
	}
}
