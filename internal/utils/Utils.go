package utils

import (
	"fmt"
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

func ParseInt(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Sprintf("Error Parsing Int %s", value))
	}
	return i
}

func ParseFloat32(value string) float32 {
	f64, err := strconv.ParseFloat(value, 32)
	if err != nil {
		panic(fmt.Sprintf("Error Parsing Float %s", value))
	}
	return float32(f64)
}
