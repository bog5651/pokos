package utils

import (
	"fmt"
	"strconv"
)

func IntJoin(arr []int64, split string) string {
	var str string
	for _, value := range arr {
		str += fmt.Sprintf("%s%s", strconv.Itoa(int(value)), split)
	}
	return str
}
