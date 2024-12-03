package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var inputFile = "input.txt"

func main() {
	args := os.Args[1:]

	if len(args) == 2 && args[1] == "base" {
		inputFile = fmt.Sprint("input.base.", args[0], ".txt")
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

func part1() {
	text, err := readFile()
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	validMul := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	allMatches := validMul.FindAllString(text, -1)

	multSum := 0

	for _, match := range allMatches {
		allNumbers := regexp.MustCompile(`[0-9]{1,3}`).FindAllString(match, -1)

		multi := 1
		for _, number := range allNumbers {
			multi *= toInt(number)
		}

		multSum += multi
	}

	fmt.Println(multSum)
}

func part2() {
	text, err := readFile()
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	validMul := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don\'t\(\)`)
	allMatches := validMul.FindAllString(text, -1)

	multSum := 0
	canMult := true

	for _, match := range allMatches {
		if regexp.MustCompile(`do\(\)`).MatchString(match) {
			canMult = true
			continue
		}

		if regexp.MustCompile(`don\'t\(\)`).MatchString(match) {
			canMult = false
			continue
		}

		allNumbers := regexp.MustCompile(`[0-9]{1,3}`).FindAllString(match, -1)

		multi := 1
		if !canMult {
			multi = 0
		}

		for _, number := range allNumbers {
			multi *= toInt(number)
		}

		multSum += multi
	}

	fmt.Println(multSum)
}

func readFile() (string, error) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(data), nil
}

func toInt(n string) int {
	value, _ := strconv.Atoi(n)
	return value
}
