package kata

import "strings"

func AbbrevName(name string) string{
	abbr := strings.Split(name, " ")
	fname, sname := abbr[0], abbr[1]
	return strings.ToUpper(string(fname[0])) + "." + strings.ToUpper(string(sname[0]))
}