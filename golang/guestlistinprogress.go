package kata

import (
	"fmt"
	"sort"
	"strings"
)

func Meeting(s string) string {
	var err error

	guests := strings.Split(s, ";")

	for _, v := range guests {
		names := strings.Split(v, ":")

		splitnames := append(splitnames, names)
		
		if err != nil {
			return err
		}

	for _, n := range splitnames {
		firstname := strings.ToUpper(n[0])
		lastname := strings.ToUpper(n][1])

		formattedname := append(formattedname, lastname, firstname)
		formattednames := append(formattednames, formattedname)

		if err != nil {
			return err
		}
	}

	}

}