package main

import "strings"

func findIndex(target string, arr []string) int {
	for i, val := range arr {
		if val == target {
			return i
		}
	}
	return -1 // If the string is not found in the array
}

func mapKeyCodeToInteger(keyCode string) int {
	keyMap := map[string]int{
		"":   0,
		"49": 1,
		"50": 2,
		"51": 3,
		"52": 4,
		"53": 5,
		"54": 6,
		"55": 7,
		"56": 8,
		"57": 9,
	}

	// Split the input string by the "+" delimiter
	keys := strings.Split(keyCode, "+")

	// Initialize the result
	result := 0

	// Map the key codes to integer keys and exclude modifier keys
	for _, key := range keys {
		if key != "2" && key != "4" {
			result += keyMap[key]
		}
	}

	return result
}
