package day18

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Solve() string {
	file, err := os.Open("day18/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	eq := ""
	for scanner.Scan() {
		eq = add(eq, scanner.Text())
	}
	return magnitude(eq)
}

func add(a string, b string) string {
	if a == "" {
		return b
	}
	eq := fmt.Sprintf("[%s,%s]", a, b)
	exploded := true
	for exploded {
		eq, exploded = explode(eq)
		if !exploded {
			eq = split(eq)
		}
	}
	return eq
}

func split(eq string) string {
	exploded := false
	reg, _ := regexp.Compile("[0-9]{2}")
	for {
		str := reg.FindString(eq)
		if str == "" {
			break
		}
		num, _ := strconv.Atoi(str)
		replace := fmt.Sprintf("[%d,%d]", num/2, int(math.Ceil(float64(num)/2)))
		eq = strings.Replace(eq, str, replace, 1)
		eq, exploded = explode(eq)
		for exploded {
			eq, exploded = explode(eq)
		}
	}
	return eq
}

func explode(eq string) (string, bool) {
	digitBuffer := ""
	arr := make([]string, 0, len(eq))
	for _, c := range eq {
		if unicode.IsDigit(c) {
			digitBuffer += string(c)
		} else {
			if digitBuffer != "" {
				arr = append(arr, digitBuffer)
				digitBuffer = ""
			}
			arr = append(arr, string(c))
		}
	}
	depth := 0
	lastNum := -1
	sliceStart := -1
	sliceEnd := -1
	replace := ""
	for i, v := range arr {
		if v == "[" {
			depth += 1
		} else if v == "]" {
			depth -= 1
		} else if num, err := strconv.Atoi(v); err == nil {
			if depth < 5 {
				lastNum = i
				continue
			}
			sliceStart = i - 1
			sliceEnd = i + 3
			if lastNum == -1 {
				replace += "0"
			} else {
				previusNum, _ := strconv.Atoi(arr[lastNum])
				arr[lastNum] = strconv.Itoa(previusNum + num)
			}
			for j, char := range arr[i+3:] {
				if nextNum, err := strconv.Atoi(char); err == nil {
					numRight, _ := strconv.Atoi(arr[i+2])
					arr[i+3+j] = strconv.Itoa(nextNum + numRight)
					break
				}
			}
			tmp := arr[:sliceStart]
			tmp = append(tmp, "0")
			tmp = append(tmp, arr[sliceEnd+1:]...)
			return strings.Join(tmp, ""), true
		}
	}
	return eq, false
}

func magnitude(eq string) string {
	regex := regexp.MustCompile(`\[\d+,\d+\]`)
	max := 0
	for {
		str := regex.FindString(eq)
		if str == "" {
			break
		}
		tmp := strings.Split(str[1:len(str)-1], ",")
		a, _ := strconv.Atoi(tmp[0])
		b, _ := strconv.Atoi(tmp[1])
		value := 3*a + 2*b
		if value > max {
			max = value
		}
		eq = strings.Replace(eq, str, strconv.Itoa(value), -1)
	}
	return fmt.Sprintf("%s Max: %d", eq, max)
}
