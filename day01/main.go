package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	content, err := ioutil.ReadFile("calibration.txt")
	if err != nil {
		log.Fatal("Error while reading File")
	}
	raw_calibration := string(content)
	calibrations := strings.Split(raw_calibration, "\n")
	result := 0

	// Loop over the string list
	// For each of those find the numbers
	// Calculate the result & drop in results array

	for _, obs := range calibrations {
		obs_strings := []string{}
		i := 0
		for i < len(obs) {
			if r, ok := extractNumber(rune(obs[i])); ok {
				obs_strings = append(obs_strings, r)
				i++
			} else if digit, _, ok := extractWrittenNumbers(obs[i:]); ok {
				obs_strings = append(obs_strings, digit)
				i++
			} else {
				i++
			}
		}
		if len(obs_strings) == 0 {
			continue
		}
		obs_result, _ := strconv.Atoi(obs_strings[0] + obs_strings[len(obs_strings)-1])
		s := fmt.Sprintf("%v - %d", obs_strings, obs_result)
		fmt.Println(s)
		result = result + obs_result
	}
	fmt.Print(result)
}

func extractWrittenNumbers(s string) (string, int, bool) {
	writtenNumbers := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for word, digit := range writtenNumbers {
		if strings.HasPrefix(s, word) {
			return digit, len(word), true
		}
	}
	return "", 0, false

}

func extractNumber(ch rune) (string, bool) {
	if unicode.IsNumber(ch) {
		return string(ch), true
	}
	return "", false
}
