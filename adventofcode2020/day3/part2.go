package main

import (
	"fmt"
	"strings"
	"sync"
)

type testCase struct {
	slope [][]rune
	right int
	down  int
}

func main() {

	slope := makeSlope()
	testCases := []testCase{
		{slope, 1, 1},
		{slope, 3, 1},
		{slope, 5, 1},
		{slope, 7, 1},
		{slope, 1, 2},
	}

	testsChan := make(chan int)
	finalValue := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(testCases))

	for _, item := range testCases {
		go func(test testCase) {
			defer wg.Done()
			testsChan <- doMap(test)
		}(item)
	}

	go doReduce(testsChan, finalValue)

	wg.Wait()
	close(testsChan)

	fmt.Println(<-finalValue)

}

func doMap(test testCase) int {
	x, y, trees := 0, 0, 0
	for i := 0; i < len(test.slope)-1; i++ {
		y += test.right
		x += test.down
		if len(test.slope) > x && test.slope[x][y] == Tree {
			trees++
		}
	}
	return trees
}

func doReduce(mapList chan int, sendFinalValue chan int) {
	final := 0
	for result := range mapList {
		if final == 0 {
			final = result
		} else {
			final *= result
		}
	}
	sendFinalValue <- final
}

func makeSlope() [][]rune {
	input := strings.Split(RawInput, "\n")
	slope := make([][]rune, len(input))
	for x, row := range input {
		slope[x] = make([]rune, len(row)*500)
		for y, c := range repeat(row, 500) {
			slope[x][y] = c
		}
	}
	return slope
}

func repeat(input string, count int) string {
	var result string
	for i := 0; i < count; i++ {
		result += input
	}
	return result
}

const Tree = rune('#')
const Gap = rune('.')

const RawInput = `.#.#....##.......#..........#..
...#...........##...#..#.......
#.####......##.#...#......#.#..
##.....#.#.#..#.#............#.
##.....#....#.........#...##...
###..#.....#....#..............
..........#..#.#..#.#....#.....
##.....#....#.#...#.##.........
#...#......#....##....#..#.#...
.##.##...#....##..#.#.....#...#
.....#.#..........##.#........#
.##..................#..#..##.#
#.#..........##....#.####......
.#......#.#......#.........#...
#....#..##.##..##........#.#...
##..#.##..#...#..####.#..#.....
###....#.###.##...........##..#
.....#.##.....##.#..#####....##
....#.###....#..##....##...#...
..###.#...##.....#.##..#..#.#..
#...#..#..#.........#..#.......
##..#.#.....#.#.#.......#...#.#
...#...##.#........#...#.......
..#..#.#..#...#...#...........#
........#.....#......#...##....
#........##.##.#.#...#...#.....
####.......#.##.###.#....#.....
...#...........#...#......#...#
##...#...#............#.......#
....#...........##.......#.....
###......#.....#....#...#.#...#
.....##..........#.......#.#...
##.##.##...#......#....#.......
##..#.#..#......#...#..#.......
....#....##.##............####.
..#.###..#.##.###..#.##.......#
#.##..#.#.....#..#.....##......
..##..#.....##.#.##........#...
.#..#.#......#..#............#.
.....#..#.#...#....#.##.#......
.#...##.#..#.#...##...##..##...
###............#.#..#..#...#...
..#..##.####.#.....#.....##.###
#....#.##..##....#..#...#.##.#.
.....#.##.........##...##......
.........####.#....#.#......#.#
.........#.#..#...#.#..#.#....#
.#.....#..##.##..##....#.......
..........##......#.##.###....#
.##...###..##.#...#........##..
..............#.#....#.#.###.##
..##.##.......#.#......##...#..
.#.....#..##..#.###...#..#.##.#
#.....#.#..#...#........#...#..
.#......#....#.#.....###...#..#
..##.#....#..##......#.....#...
..#.#.##..#.....#.####..###....
.........#......#..#...........
..#........#.##.#.....##.##..#.
.......#.........#....#...#.#..
.##.....#.#....#.#.......#.....
..........#.##........##...##..
###..###.#.#..#..#####.##.#.##.
..##..##.#.#...#..#.#.#......#.
#..#..#..#..##..#.....#......#.
..#....#.##..#......##.........
..#.##......#...##.#......#....
.......#..#.##.#.....#.........
.......#.#.#.###...##......#...
.....#.#..........#..#...#.....
....##..........#..........##..
..#......#.....#.##.#..#...#.#.
....#.....#..#...#..#.#.##..###
.####....#........#...#........
...##.#.##.#..#...##...#.##....
....#...#...#.#.#.#...#..#.....
.....#...#.#.....#.#........##.
..#.#.......###.#.....##.......
......#.........##....#....#..#
.............##.....##.........
.........##...##.......#.....#.
##.........#..........#.###..##
...#.....#......#....#..##.....
##..#...#...##.#.....#.#......#
..#...##.#.......#.#......#.##.
......#.......#.#...........#..
..........#.....##............#
#........#...#..#.......###.##.
.##...........#.#........#.#.#.
...#..##...#.#....#####.#......
.....##...###...#..#.##...####.
...#....#.....#..#.......#.....
#....#....#...#..#..#.######..#
#.###...........#......#...#..#
.#.#.#.#..#....#....#...##.#...
.#..#.........#.#....###...#...
......#..##.##..........#....##
.....#......##....##.....#...#.
.#...#.#.#....##....#..#....#.#
..................#..###.#..##.
..#.........#......#....#..###.
#.#.....#..#..#....###..###....
..##..##.#..##........##...##..
##..#........##..###..#.....#.#
..#..###..#......#....#...#...#
#..#.#..............##.#..#.#..
.....####....#...####.....#.#..
.....#....##.#......###........
##.##...#.#.#.#.......#....##..
.#......#...#.#....#..##.#.##.#
#.#.##.#.#......#..##........##
...##.....#.....#...#..###...#.
........###.....#.....#...##..#
.....#.##.##......#.#....#...#.
.#....##.......#..#.####.......
.#..#....#..........#......#.#.
.#.##.##.....###.#.#...........
.........#......#..##..........
....#...##.#.#.#..#.#.........#
..#.....#.##...#..#..#.###....#
...#.##......#.....##....#.....
###............#.#....#...#....
.......#.....#..#.#.#....#..#.#
...#......#.#..##..#....#...#.#
............##........##..##...
..#..#.##..#......###..#.......
........#.........#............
..#...#.#########.#...##..###..
#....#......#.......#.#.....#..
#.#..#....###.###....#...#.#...
#...###.#.#.......#.##......#..
.................#...#.#.#.....
##....#...#........#....#.#..#.
......#.....#...#..........#.#.
##..........#...#..........#.##
..#.#.##.#....#.#......#...##..
.....#.......#..#.....#........
#.##.#..##..#.......##.........
....#......#..#..#.#...#.......
...#....#................###...
.##.....#.#....#.#..........##.
...#..#....#.##.##......#......
..#.#....#.......#.#..##.......
....#.....#..........##.#.#####
#.....................##..#..#.
.###..#.##.......##.#...#..#...
...###.......#..#...#......#..#
#..#...#.#..#.#..#..#.##.......
#...##.......#..#..#.##..###...
......#....#.#.#........#.##..#
..##..#....#....#..#.#..#......
..##.#...#.#######..#...#.....#
..#....#..#.........#..##......
...#....#.#......#..#..#.#.....
#..#....#........#.#..##....###
#....#..##......##.##.....#.###
...#.#..........#..#.#.#.#.##..
......##..#.#..#.#....#....#...
##....#....#..#..#.##......#...
....#.#..##.#.#...###....##.#..
...#.......##..#.......#...#...
......##.......#..##.....#...#.
...#.#...#...........#...#.....
.#....#...#......##.##..###..#.
.#..........#...#...#...##.##..
.....###..#.....#..##....#.####
..#.###..#..##..##.....#.#.....
.............#.###...##.#.....#
....###.......###.#.....#..#.#.
........##.#.........#.....###.
.....###.#..#.....#...#..#.....
.#....#..##.#..#.#....#.......#
........#......#.#..#.#..#...##
...#.##.##......#..............
.#.....##.#.....#..#......##...
#..#..#.....#.....#.....###....
.##...........#..#.##.....#....
..#.#......#.#...#.##.#..#...##
...#..........#.....#..........
#.#.#.#.#...#....#...#.....##..
#......##...#...#..........#.#.
....##........#.#..............
#..#.#.#..#........##......#.##
........####...##.#.....#......
....#........#.#..#..##..#.#...
.#.....#..###...#..#.....#..#..
#......###.#..#....#..#.#......
....#.....##.##..#...#.#..##.#.
..##..#...#.#......#....#...#.#
#..##...##..#...###...#..#.....
.......#.....#...........##....
#..##....#........#....##..#.#.
.#........#..##...###.#..#.....
.#.#....#..##...#...##.#..###..
#.........#.......#.....#.#....
#..#.....#.#.###.#..#......#...
....#..#.#....#..##..###....###
###.##.#.#..#...........#.#.#..
..##.#.......#......#..##....#.
.....#.#.#.......##.......#...#
...........#.##....##.##....#.#
...#.......#..#.##..#......#..#
#.#.#...#......##.#...........#
##........#...........###.#..#.
..........#.#.#....#.#..##.#.#.
...#.#.#....#..........#..#....
#.#....###.#.#..#.......###...#
.#....#......#.#.#..#..#.......
......##.............#....#.#.#
.#..........#.........#.##.....
##....#....##....#..#.......#..
#.##.##.#..#..#.....#..#.##.#..
.#..#.......##..#.....##.##....
.......#..........#.#.##..#.##.
....#.....#.#...##....##.......
.......#.........#...##....##.#
#.....#......#..........#...#..
...#.#.......#.#..#....###..#..
.....#.#.#.........#...........
.#..###.#.#........#.#.........
.........#..#......##...##....#
...###..#.....##.....#.###....#
.##...#...#........###.#..#....
.##........#..#.###.######.##.#
##.#...#.#....#..##.#....##....
.......##.....##.#..###.#......
..##...##........#.......#....#
#..##...#.####...###......#...#
.##.....#.##.#.#.....###.#..##.
..###....#.#.###.#....#........
....#..###..#...#....#..#..#.#.
#.#.##....##...##.......#......
.........#...#....#..#.........
.............#...#..##.#.......
...#.##.......#...#.#..##.##...
.####.#.##..#.#......#.##...#.#
.#..#.#.....#.................#
..#.##..###....#...#......####.
..##..##...........#....#...#..
....#...#...#...#.......#....#.
#.#...###...#...#.#...#....##.#
......#...#.#.......#.....#...#
....##...#.#.#....#....#.#....#
.....#.....#...##..#...#....##.
#.....#....#......##.##....#...
...#.#....#...#....#.#....##..#
...#.#..#...##....###..#.......
...##......###...###.#...#..#..
##.......#.......###.......#..#
..##.##..###.#............#...#
#.....##..#..##....##..#.......
......#.#...#......#.....#.....
#...........#....#..##.##.#....
.......#..#......#...#....#...#
.#...##...........#......#...#.
#........#....##...###.#....#..
.....#.......##.........#.##...
.#.###..#....#..##.#..#.#..#...
#.......#.##.#.#....#.#..#....#
###.....#.#.......#..#......#.#
#..#.#.......#.#..##..##.#.#...
#..#.#.#.###........#.....#...#
#.#.#..#..##.....#...........#.
..#.#..#.....#...#...#...##....
...#.##......#...##.#...#.#.#.#
#..#.#.#.#.......####..........
..#......#.#......##.###.....##
..#...##..#.........##....#.##.
##.##.##.#.#.....#..........##.
.#.....###.#..#....#..#.###...#
#...##.......###....#.#..#.....
..#....##.........##.........##
......#....#.##.......#........
..#.#.#..#...#...#...##.#...#..
......#..##.#.#.#...##...#.#.##
#..#...##.#.....#...#.##.......
..#..#.........##.#...#.##...##
##.##.#....#.......#.##..#.....
.....##...##.##...##.........##
#......#...#.......#...#...#...
...##...........#...#..#.......
.#.##.#..#........#....#.......
#.#...#..#......##...#.#.##....
##........####..#.#...#.#.##.##
#..#.#.##......##.#.#..#.......
.....#.........#..#.####....#..
......##..#....#...#.#....#....
#...##........#.........#.....#
.#.#...#.#.#..#............##.#
.#..#....#....#.....#...#.....#
..###...#..#.....#.##.###...#.#
.#.###..#..#...#.#...#.#......#
#...#####......###........##...
.....#.....#..#.#....#..##.....
....##...#.#.##.#####...#....#.
.#.#.........##.#.......#..##..
.#...#.#...#...#....#.#...##.#.
.##...#..#.#..#......#.#.#..##.
..#.....#..#.....##.....#......
..#........#..##...#.......###.
.#....#.......#....#....#..#...
....#......#.#.#.........#.....
..##...#.#.#...#.#........#....
.#.....####...##.#..#...##.....
...#.....#...#...#....#....#...
.........#..#.#.....#..#.#..#..
.........##...........#.......#
......#..#.....##...#.##.#.....
.#......##........##...#.#.##..
.....#.#..##...........#..#..#.
...#.......#...#.#..#.##..#.##.
...#.......#.....#.#...#.##.#..
#.....#.............##.#..####.
.#...#......#...##.#....#.#....
.##..##.##....#.#.....#.......#
...#...#....#....##.#..#....##.
..............##....#.......#.#
.#.#.#...##..#..#...###.#..#...
.#.#...#.#..#.#..#...######..#.
........#......#.#..#.#....#...
..###.....###.#.##....#...##...
.##.#.....#.......##.......#...
..#..##...#..........#.#....#.#`
