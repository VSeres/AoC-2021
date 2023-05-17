package day24

import (
	"strconv"
)

var div = [14]int{1, 1, 1, 26, 1, 26, 1, 26, 1, 1, 26, 26, 26, 26}
var cehck = [14]int{13, 12, 11, 0, 15, -13, 10, -9, 11, 13, -14, -3, -2, -14}
var offset = [14]int{14, 8, 5, 4, 10, 13, 16, 5, 6, 13, 6, 7, 13, 3}

type Key struct {
	depth int
	z     int
}

var badStates = make(map[Key]bool)

func Solve(smallest bool) string {
	if smallest {
		return run(0, 0, "", []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	}
	return run(0, 0, "", []int{9, 8, 7, 6, 5, 4, 3, 2, 1})
}

func run(depth int, z int, modelNumber string, numbers []int) string {
	if badStates[Key{depth: depth, z: z}] || depth == 14 {
		return ""
	}
	startZ := z
	for _, w := range numbers {
		z = startZ
		x := startZ
		y := 25

		x %= 26
		z /= div[depth]
		x += cehck[depth]
		if x == w {
			x = 1
		} else {
			x = 0
		}
		if x == 0 {
			x = 1
		} else {
			x = 0
		}
		y *= x
		y += 1
		z *= y
		y = 0
		y += w
		y += offset[depth]
		y *= x
		z += y

		if depth == 13 && z == 0 {
			return modelNumber + strconv.Itoa(w)
		}

		num := run(depth+1, z, modelNumber+strconv.Itoa(w), numbers)
		if num != "" {
			return num
		}
	}
	badStates[Key{depth: depth, z: startZ}] = true
	return ""
}
