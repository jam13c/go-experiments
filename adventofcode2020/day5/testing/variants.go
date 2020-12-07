package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println(seatID(traverse("BFFFBBFRRR")))
	fmt.Println(seatIDRegex(regexp.MustCompile("B|F|R|L"), "BFFFBBFRRR"))
}

func seatIDRegex(re *regexp.Regexp, input string) int64 {
	result := re.ReplaceAllStringFunc(input, func(match string) string {
		if match == "B" || match == "R" {
			return "1"
		}
		return "0"
	})

	val, _ := strconv.ParseInt(result, 2, 32)
	return val
}

func seatID(row, col int) int {
	return (row * 8) + col
}

func traverse(input string) (int, int) {
	rowMin, rowMax := 0, 127
	for i := 0; i < 7; i++ {
		bisect(&rowMin, &rowMax, input[i] == 'F')
	}
	if rowMax != rowMin {
		panic("Row not found correctly")
	}
	colMin, colMax := 0, 7
	for i := 7; i < 10; i++ {
		bisect(&colMin, &colMax, input[i] == 'L')
	}
	if colMin != colMax {
		panic("Column not found correctly")
	}
	return rowMin, colMin

}

func bisect(min, max *int, takeLower bool) {
	half := (*min) + ((*max - *min) / 2)
	if takeLower {
		*max = half
	} else {
		*min = half + 1
	}
}
