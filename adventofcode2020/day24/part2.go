package main

import (
	"fmt"
	"strings"
	"time"
)

type direction int

var directions = []hex{
	newHex(1, 0),
	newHex(1, -1),
	newHex(0, -1),
	newHex(-1, 0),
	newHex(-1, +1),
	newHex(0, +1),
}

type hex struct {
	q int // x axis
	r int // y axis
	s int // z axis
}

const inpu = `sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`

func main() {
	start := time.Now()
	directionsList := readInput(inpu)
	tiles := initiliseTiles(directionsList)

	for i := 0; i < 100; i++ {
		tiles = gameOfLife(tiles)
		fmt.Printf("day %v: %v\n", i+1, countBlackTiles(tiles))
	}

	fmt.Println(countBlackTiles(tiles))
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Print("Execution time: ")
	fmt.Println(elapsed)
}

func gameOfLife(tiles map[hex]bool) map[hex]bool {
	newTiles := make(map[hex]bool)
	for tile, isBlack := range tiles {
		for d := 0; d < 6; d++ {
			if _, ok := tiles[hexNeighbor(tile, direction(d))]; !ok && isBlack {
				tiles[hexNeighbor(tile, direction(d))] = false
			}
		}
	}
	for tile, isBlack := range tiles {
		blackNeighbours := 0
		for d := 0; d < 6; d++ {
			if blackNeighbour, ok := tiles[hexNeighbor(tile, direction(d))]; ok && blackNeighbour {
				blackNeighbours++
			}
		}
		if isBlack {
			if blackNeighbours == 0 || blackNeighbours > 2 {
				newTiles[tile] = false
			} else {
				newTiles[tile] = true
			}
		} else {
			if blackNeighbours == 2 {
				newTiles[tile] = true
			} else {
				newTiles[tile] = false
			}
		}
	}

	return newTiles
}

func initiliseTiles(directionsList [][]direction) map[hex]bool {
	tiles := make(map[hex]bool)
	refTile := newHex(0, 0)
	tiles[refTile] = false

	for _, directions := range directionsList {
		currTile := refTile
		for _, direction := range directions {
			currTile = hexNeighbor(currTile, direction)
			if _, ok := tiles[currTile]; !ok {
				tiles[currTile] = false
			}
		}
		if _, ok := tiles[currTile]; ok {
			tiles[currTile] = !tiles[currTile]
		} else {
			tiles[currTile] = true
		}
	}
	return tiles
}

func countBlackTiles(tiles map[hex]bool) int {
	blackCount := 0
	for _, isBlack := range tiles {
		if isBlack {
			blackCount++
		}
	}
	return blackCount
}

func newHex(q, r int) hex {
	h := hex{q: q, r: r, s: -q - r}
	return h
}

func hexNeighbor(h hex, direction direction) hex {
	directionOffset := directions[direction]
	return newHex(h.q+directionOffset.q, h.r+directionOffset.r)
}

func readInput(text string) [][]direction {

	lines := strings.Split(text, "\n")

	directions := make([][]direction, len(lines))
	for x, line := range lines {
		for i := 0; i < len(line); i++ {
			if string(line[i]) == "e" {
				directions[x] = append(directions[x], 0)
			} else if string(line[i]) == "w" {
				directions[x] = append(directions[x], 3)
			} else if string(line[i]) == "n" {
				if string(line[i+1]) == "e" {
					directions[x] = append(directions[x], 5)
				} else if string(line[i+1]) == "w" {
					directions[x] = append(directions[x], 4)
				}
				i++
			} else if string(line[i]) == "s" {
				if string(line[i+1]) == "e" {
					directions[x] = append(directions[x], 1)
				} else if string(line[i+1]) == "w" {
					directions[x] = append(directions[x], 2)
				}
				i++
			}
		}
	}
	return directions
}
