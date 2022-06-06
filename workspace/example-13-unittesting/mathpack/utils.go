package mathpack

import (
	"strconv"
	"strings"
)

func Sum(input string) (total int, err error) {

	strNums := strings.Split(input, ",")
	for _, sn := range strNums {
		var n int
		n, err = strconv.Atoi(strings.TrimSpace(sn))
		total += n
	}

	return
}
