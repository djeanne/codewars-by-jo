package kata

import (
	"sort"
	"strings"
)

func Meeting(s string) string {

	var err error

	guests := strings.Split(s, ";")

	splitnames := make([][]string, 0)

	for _, v := range guests {
		names := strings.Split(v, ":")

		splitnames = append(splitnames, names)

		if err != nil {
			panic(err)
		}
	}

	new_list := make([]string, 0)

	for _, n := range splitnames {
		n[0], n[1] = strings.ToUpper(n[1]), strings.ToUpper(n[0])

		guest := strings.Join(n, ", ")
		new_guest := "(" + guest + ")"
		new_list = append(new_list, new_guest)

		if err != nil {
			panic(err)
		}
	}

	sort.Strings(new_list)

	guest_list := strings.Join(new_list, "")

	return guest_list
}
