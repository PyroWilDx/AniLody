package utils

import (
	"strconv"
)

func IntSliceToStrSlice(intSlice []int) []string {
	strSlice := make([]string, len(intSlice))
	for i, v := range intSlice {
		strSlice[i] = strconv.Itoa(v)
	}
	return strSlice
}
