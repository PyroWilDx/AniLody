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

func IsLetter(c byte) bool {
	return IsLowerCaseLetter(c) || IsUpperCaseLetter(c)
}

func IsLowerCaseLetter(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func IsUpperCaseLetter(c byte) bool {
	return c >= 'A' && c <= 'Z'
}
