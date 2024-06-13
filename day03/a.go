package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

/*
. . . . 4 3 2 & . . .
5 7 . . * . . 6 9 . .
. . . $ 7 8 9 . . . .

Go over array row-by-row, when you hit a number you can assume it's the leftmost number of total and following is created:
* numberBuffer, which stores the entire sequential number
* specialHit, wboolean which indicates for number that specialCharacter exists in adjacent dimensions

As soon as we fully pass a number these values are reset & if specialHit is true then numberBuffer is stored in resultsList.

This list is then reduced at the end

*/

func isSymbol(c rune) bool {
	return string(c) != "." && !unicode.IsDigit(c)
}

func isAdjacentToSymbol(schematic [][]rune, row, col int) bool {
	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	rows := len(schematic)
	cols := len(schematic[0])

	for _, dir := range directions {
		ni, nj := row+dir[0], col+dir[1]
		if ni >= 0 && ni < rows && nj >= 0 && nj < cols {
			if isSymbol(schematic[ni][nj]) {
				return true
			}
		}
	}
	return false
}

func getPartNumberTotal(grid [][]rune) int {
	var results [][]string
	var noPartsResults [][]string

	for i := 0; i < len(grid); i++ {
		stringBuffer := ""
		specialHit := false
		results = append(results, []string{})
		noPartsResults = append(noPartsResults, []string{})

		for j := 0; j < len(grid[i]); j++ {
			r := grid[i][j]
			if string(r) == "." && stringBuffer != "" && specialHit {
				results[i] = append(results[i], stringBuffer)
				specialHit = false
				stringBuffer = ""
				continue
			}
			if isSymbol(r) && stringBuffer != "" && specialHit {
				results[i] = append(results[i], stringBuffer)
				specialHit = false
				stringBuffer = ""
				continue
			}
			if unicode.IsDigit(r) {
				stringBuffer += string(r)
				if isAdjacentToSymbol(grid, i, j) {
					specialHit = true
				}
				continue
			}
			noPartsResults[i] = append(noPartsResults[i], stringBuffer)
			specialHit = false
			stringBuffer = ""
			// fmt.Printf("dimension = {%d,%d}, rune = %s, buffer = %s, specialHit = %v\n", i, j, string(r), stringBuffer, specialHit)
		}
	}

	fmt.Printf("%v", noPartsResults)
	sum := 0
	for _, line := range results {
		for _, val := range line {
			s, _ := strconv.Atoi(val)
			sum += s
		}
	}
	return sum
}

func readFileInto2dArray(filePath string) [][]rune {
	// Initialize a 2D slice to hold the characters
	var grid [][]rune

	file, err := os.Open("a.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %s", err)
	}

	// Print the 2D slice
	// for _, row := range grid {
	// 	fmt.Println(row)
	// }

	return grid
}

func main() {
	grid := readFileInto2dArray("a.txt")
	sum := getPartNumberTotal(grid)
	fmt.Println(sum)
}
