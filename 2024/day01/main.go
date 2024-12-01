package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var inputFile = "input.txt"

func main() {
	args := os.Args[1:]

	if len(args) == 2 && args[1] == "base" {
		inputFile = "input.base.txt"
	}

	if args[0] == "part1" {
		part1()
		return
	}

	if args[0] == "part2" {
		part2()
		return
	}
}

func part2() {
	lines, err := readFile()
	if err != nil {
		fmt.Println("Error reading file")
	}

	firstLocationIds := []int{}
	secondLocationIncidenceMap := map[int]int{}

	for _, line := range lines {
		ids := strings.Split(line, "   ")

		firstId, _ := strconv.Atoi(ids[0])
		firstLocationIds = append(firstLocationIds, firstId)

		secondId, _ := strconv.Atoi(ids[1])

		_, ok := secondLocationIncidenceMap[secondId]
		if ok {
			secondLocationIncidenceMap[secondId] += 1
		} else {
			secondLocationIncidenceMap[secondId] = 1
		}
	}

	similarity := 0
	for _, id := range firstLocationIds {
		incidence := secondLocationIncidenceMap[id]
		similarity += incidence * id
	}

	fmt.Println(similarity)
}

func part1() {
	lines, err := readFile()
	if err != nil {
		fmt.Println("Error reading file")
	}

	distance := 0

	firstLocationIds := []int{}
	secondLocationIds := []int{}

	for _, line := range lines {
		ids := strings.Split(line, "   ")

		firstId, _ := strconv.Atoi(ids[0])
		firstLocationIds = append(firstLocationIds, firstId)

		secondId, _ := strconv.Atoi(ids[1])
		secondLocationIds = append(secondLocationIds, secondId)
	}

	sort.Ints(firstLocationIds)
	sort.Ints(secondLocationIds)

	for i := range len(firstLocationIds) {
		distance += absolute(firstLocationIds[i] - secondLocationIds[i])
	}

	fmt.Println(distance)
}

func readFile() ([]string, error) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
		return []string{}, err
	}

	lines := strings.Split(string(data), "\n")

	if len(lines) > 0 {
		lines = lines[:len(lines)-1]
	}

	return lines, nil
}

func absolute(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
