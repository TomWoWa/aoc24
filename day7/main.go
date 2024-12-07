package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"git.hupojo.de/tom/aoc24/utils"
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	utils.Start()

	lines := strings.Split(strings.Trim(string(data), "\n"), "\n")

	result := int64(0)

	for _, line := range lines {
		split := strings.Split(line, ":")
		test, _ := strconv.ParseInt(split[0], 10, 64)
		values := make([]int64, 0)
		for _, s := range strings.Split(strings.Trim(split[1], " "), " ") {
			value, _ := strconv.ParseInt(s, 10, 64)
			values = append(values, value)
		}

		result += check(test, values)
	}

	utils.Submit(fmt.Sprint(result))
}

func check(test int64, values []int64) int64 {
	if len(values) == 1 && values[0] == test {
		return test
	}

	if len(values) <= 1 {
		return 0
	}

	if mult := check(test, append([]int64{values[0] * values[1]}, values[2:]...)); mult > 0 {
		return mult
	}
	if add := check(test, append([]int64{values[0] + values[1]}, values[2:]...)); add > 0 {
		return add
	}

	if concat := check(test, append([]int64{merge(values[0], values[1])}, values[2:]...)); concat > 0 {
		return concat
	}

	return 0
}

func merge(a, b int64) int64 {
	res := (a * int64(math.Pow10(len(fmt.Sprint(b))))) + b
	return res
}
