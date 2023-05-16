package day2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve(partTwo bool) string {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	x := 0
	y := 0
	aim := 0
	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")
		num, _ := strconv.Atoi(command[1])
		switch command[0] {
		case "up":
			if partTwo {
				aim -= num
			} else {
				y -= num
			}
		case "down":
			if partTwo {
				aim += num
			} else {
				y += num
			}
		case "forward":
			x += num
			if partTwo {
				y += num * aim
			}
		}
	}

	return strconv.Itoa(x * y)
}
