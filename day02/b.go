package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	gameId int
	grabs  []Grab
}

type Grab struct {
	redCubes   int
	greenCubes int
	blueCubes  int
}

func (g Game) String() string {
	return fmt.Sprintf("Game %d: %v", g.gameId, g.grabs)
}

func (g Grab) String() string {
	return fmt.Sprintf("%d red, %d, blue, %d green", g.redCubes, g.blueCubes, g.greenCubes)
}

// Calculate total sum of all cubes per row per color
// Should be less than or equal to  12 red, 13 green and 14 blue
// Sum all game ID's that are possible
func main() {
	lines := readFileList()
	total := 0
	for _, line := range lines {
		game := parse_row(line)
		idValue := validate_cubes(game)
		fmt.Printf("%v - %d\n", game, idValue)
		total += idValue
	}
	fmt.Println(total)
}

func validate_cubes(game Game) int {
	r, gr, b := 0, 0, 0
	for _, g := range game.grabs {
		if g.redCubes > r {
			r = g.redCubes
		}
		if g.greenCubes > gr {
			gr = g.greenCubes
		}
		if g.blueCubes > b {
			b = g.blueCubes
		}
	}
	return r * gr * b
}

func parse_row(line string) Game {
	subs := strings.Split(line, ":")
	gameId, err := strconv.Atoi(strings.ReplaceAll(subs[0], "Game ", ""))
	if err != nil {
		fmt.Println(err)
	}
	grabList := parse_entries(subs[1])
	return Game{gameId: gameId, grabs: grabList}
}

func parse_entries(entries string) []Grab {
	entryList := strings.Split(entries, ";")
	var grabList []Grab
	for _, entry := range entryList {
		r, g, b := parse_entry(entry)
		grab := Grab{redCubes: r, greenCubes: g, blueCubes: b}
		grabList = append(grabList, grab)
	}
	return grabList
}

func parse_entry(entry string) (int, int, int) {
	red, green, blue := 0, 0, 0
	parts := strings.Split(entry, ",")
	for _, p := range parts {
		p := strings.TrimSpace(p)
		e := strings.Split(p, " ")
		val, _ := strconv.Atoi(e[0])
		if e[1] == "green" {
			green += val
		}
		if e[1] == "red" {
			red += val
		}
		if e[1] == "blue" {
			blue += val
		}
	}
	return red, green, blue
}

func readFileList() []string {
	file, err := os.Open("a.txt")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var fileLines []string
	for scanner.Scan() {
		if scanner.Text() != "" {
			fileLines = append(fileLines, scanner.Text())
		}
	}

	file.Close()

	return fileLines
}
