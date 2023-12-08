package mylib

import "strings"
import "strconv"

type String struct {
	Str string
}

func (s String) ToInt() int {
	i, _ := strconv.Atoi(s.Str)
	return i
}

func (s String) ToFloat() float64 {
	f, _ := strconv.ParseFloat(s.Str, 64)
	return f
}

func (s String) ToBool() bool {
	b, _ := strconv.ParseBool(s.Str)
	return b
}

func (s String) ReplaceAll(old, new string) string {
	return strings.ReplaceAll(s.Str, old, new)
}
