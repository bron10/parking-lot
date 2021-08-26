package packages

import (
	"fmt"
	"strings"
	// "time"
)

func ArrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
