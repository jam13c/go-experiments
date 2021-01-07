package main

import (
	"fmt"
	"strings"
)

func main() {
	//part1(strings.Split(RawInput, "\n"))
	part2(strings.Split(DemoInput, "\n"))
}

func part1(input []string) {
	floor := newFloor()

	for _, item := range input {
		floor.flipTile(item)
	}

	fmt.Println(floor.countBlack())
}

func part2(input []string) {
	floor := newFloor()
	for _, item := range input {
		floor.flipTile(item)
	}

	for i := 1; i <= 10; i++ {
		floor.runCycle()
		fmt.Printf("Day %v: %v\n", i, floor.countBlack())
	}
}

type floor map[hexVector]bool // true = white

func newFloor() floor {
	return floor(map[hexVector]bool{})
}

func (f floor) flipTile(input string) hexVector {
	t := newTokenizer(input)
	token, end := t.nextToken()
	tile := newHexVector(0, 0, 0)
	for !end {
		tile = tile.add(token.toHexVector())
		token, end = t.nextToken()
	}
	if _, found := f[tile]; found {
		//fmt.Printf("Tile %v flipping to %v\n", tile, !v)
		f[tile] = !f[tile]
	} else {
		//fmt.Printf("Tile %v flipping to false\n", tile)
		f[tile] = false
	}
	return tile
}

func (f *floor) runCycle() {

	flr := *f
	countBlack := func(vectors []hexVector) int {
		count := 0
		for _, v := range vectors {
			if isWhite, found := flr[v]; found && !isWhite {
				count++
			}
		}
		return count
	}

	for tile, isWhite := range flr {
		for _, neighbour := range tile.nearby() {
			if _, found := flr[neighbour]; !found && !isWhite {
				flr[neighbour] = false
			}
		}
	}

	newFlr := newFloor()
	for tile, isWhite := range flr {
		blackNeighbours := countBlack(tile.nearby())
		if isWhite {
			if blackNeighbours == 2 {
				newFlr[tile] = false
			} else {
				newFlr[tile] = true
			}
		} else {

			if blackNeighbours == 0 || blackNeighbours > 2 {
				newFlr[tile] = true
			} else {
				newFlr[tile] = false
			}
		}
	}
	f = &newFlr
}

func (f floor) countBlack() int {
	count := 0

	for _, isWhite := range f {
		if !isWhite {
			count++
		}
	}
	return count
}

type tokenizer struct {
	data string
	idx  int
}

func newTokenizer(input string) *tokenizer {
	return &tokenizer{input, 0}
}

func (t *tokenizer) nextToken() (token, bool) {
	if t.idx < len(t.data) {
		curr := t.data[t.idx]
		switch curr {
		case 'w', 'e':
			t.idx++
			return token(curr), false
		case 'n', 's':
			val := t.data[t.idx : t.idx+2]
			t.idx += 2
			return token(val), false
		}
	}
	return token('\000'), true
}

type token string

func (t token) toHexVector() hexVector {
	switch t {
	case "ne":
		return newHexVector(1, -1, 0)
	case "sw":
		return newHexVector(-1, 1, 0)
	case "se":
		return newHexVector(0, 1, -1)
	case "nw":
		return newHexVector(0, -1, 1)
	case "e":
		return newHexVector(1, 0, -1)
	case "w":
		return newHexVector(-1, 0, 1)
	}
	panic("No hex vector")
}

type hexVector struct {
	a, b, c int
}

func newHexVector(a, b, c int) hexVector {
	return hexVector{a, b, c}
}

func (v *hexVector) add(h hexVector) hexVector {
	return hexVector{v.a + h.a, v.b + h.b, v.c + h.c}
}

func (v *hexVector) nearby() []hexVector {
	return []hexVector{
		v.add(newHexVector(1, -1, 0)),
		v.add(newHexVector(-1, 1, 0)),
		v.add(newHexVector(0, 1, -1)),
		v.add(newHexVector(0, -1, 1)),
		v.add(newHexVector(1, 0, -1)),
		v.add(newHexVector(-1, 0, 1)),
	}
}

const BasicInput = `esew
nwwswee`

const DemoInput = `sesenwnenenewseeswwswswwnenewsewsw
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

const RawInput = `seseswnwseneweseswseswnwnese
neneeneseneeeesenewwne
swswswneswsewswseswesewswseswneeswswnwsw
eneeswnwneneenwneneneeneneswnwneswne
seeseseeeeseenenenwswseweese
seswweseseseswseswneswwnwswswseneswnw
wswseswswneeswswsw
wnewnwnewwwwwwwwwsewwsw
wseswnenwneswwnwseeenwnwnwewesewse
seweseseneeeseeswsenwseeenwseeeew
seswswnwswswswnesesesewseeneseswsesesesw
wsenewsenwwnewswwewnwnwsewwwsw
swwwnewnwwseswswnwswenwswseeswesw
senenwswwnwnewwseseseswswwweneswwne
nwseseeeesweee
swwswswswwswswsweswnwwswsw
sewnwswneneneenenewneeneeneneswwee
swenenwswnwnenwnwnwenwnenesenwnenwnwnwwnw
eweeneswnesenesewneneeeneseneneew
enesesenwsweeeeeenwnwnwsweswew
nweeeswneswenwseswnweeeneenenesw
neeeswnwnenwnewnwnwswenwwwnwsesenw
neswnwenwneswnwnwnwneswnesenwnw
neeseeweneeswweeswneneneeneenenw
nenwnenwnenwewnewsenwnwnwnwnene
nwewnwnwwnenwnwwnwnwwnwwenwnwsenw
seseswsewneswseseseseseswwseseswe
eseswnewenwwnenene
wswneneswswsewnwseswswswseswswwwswswne
neeswneneewnenenenenwnewneeeneneswse
eneeneneneneeeneneenewsene
eneweewwsenesene
nwenwenwnwswnenwsenwnesenenenenwnesw
wesenesenewweewwwswwnwsew
eeswneeewsewneneneweneeneene
eeswswswnwweswnenwsw
swesewnwneswseesesesesesenwsenwesee
enwsenwnwswnwwnwseseseseeswewnesenw
eeseeeseenwseenweeeenwesese
nenwnwswneneswneenenwneneswnenwnwnwnenee
ewswwswwswewwnwwenww
swnenenwnenenenesenenesewnenwnenwesenesw
nwseeenwnwwnwnwenwswnwswnwnwneenwnew
ewwnwwwwnwswnwnenwsenwnwewnwww
eswwwseewsewweeswwnwnwwenwwwne
wswseswswswnenwswwswswnewswswse
nwnenwseseswsesesenwseseseweswew
swesweenenwnwnenwnenwwneneseenenwsw
wswwwswneswwwswwewnenwswwwew
wsewwwnwwwwnwsenenewnwewwsw
wswswwesweswnwswwwswswsweswswswsw
eswswwnwwswsesweswswswneswswnwnwsesw
neenwsewwswneseswwseswnwwwwswww
enwseweseseseeswseneneseswseswneee
swnwswswneseswswswsesewseseneneseswsee
wsenwwswwwnwswwwwneswewwneww
nweswsweeewnenesee
swsweswnwswseneswswswseswswswsw
swsewseswseeseneeswswwswswnenenwnesesw
seeeenenewenenenenesew
nwnesewsesesenwsesenwsenwnwnenwnenwnww
enwswsewnesesesenwe
nwnwnesenwwwwnwnw
nenesenewneenenenenwswneenesenenwnewne
wnwnesenenwesewsenenenenwneeeneswse
wwswswswswswswwswwswesww
enwsesesewseeesweneeeswesesweenw
wwnwseneneneeenewswnwnwneswsenewnwsenw
swswswswswswewswswwswsww
nenwwnenwsenwnwswseswnwneneseseesenww
seeseeseeeesenewseswweseseswnenese
wwsewwwewneeswnwwwewwwenesw
wenewseeeneneenenenenenwswswnenwnee
neswnwnwesesenwswewneewsenwesweswswse
nesweewswsewswnesewnewswwwwwsw
nwneeeneseneneeneswenenwswnewsenwe
swswswswnwswwswswesweeesenwswwswse
senenwewnenenenenwnwneneswneenenenewne
senwseneesweeesenenwseseseneewwe
ewwwwwnesenesenwnwsewsesenenewswnw
sewswswswseneswseseswseswneswswse
eenwneneeeeeewswneseneeew
swwwwseswswnweswwnwenenweenwseeene
eeneeeeswnwneseenwesweeenweee
enwwwwnwnwwwnwnwneseswswnwnwnwwenwnw
wwwswswswswwswewswsw
swwwwnesewswewwswneswseenw
seeneeeeeweeeseee
swneswwnwsenweswnweswswnwseeseseseswnw
swswseswnwswswsweseswnwswswswneswsesew
neneswswwwwwswwnwswswwswsesweswsesw
sweneswnesenenwnenwesenewnwnewnwwene
nwnwnenenwwnesweeswnwswseseneseewnw
nenwnenenwneneneswneneenenwswnenw
nwseswseswwwswnwneswswwswwwwswnese
esenwenwnwwnewnwnwneswnwnwwnwnwnwnw
wwwswsenwwswwwwwwnwnwwsewwe
wesenweneeeseswseeneswswnenwwesw
swswnwseseswsenwnenwnwseewweswswseene
swswnewswswneswswswsenwweswwswswswsww
eswwnenwnwneewenweswswnewweswese
seneenwsewwnwnee
swswswswsenwsweswswswenwswsenwswsesesw
nwnweenwnwnenwnwneswswnwnwsenwnwnwnwnwne
wnwwnwwwwwwsew
nwswnwnwnenenwnwnwswnwseneswnwnwnwne
newwneeeswseneeneneneneweeeenene
eeenwnwneeeeeeesweenwswwseew
esenewwwwnwswswswswwwnwswswswww
eeweeeseseeesee
wewwnenewswwneseswwnwwwseswswww
wweenenewswswnwwewwswwwswswsww
seseeswneesewnwsenwsesesesenwwswnwsee
nwnwnenwnenwnwnwsesenwnwnwwnwnenenwnwse
eswswseseswwswseseswswnweseswenwsesw
eeseneneeesenenwswwweewneenene
swsenwwswnwnwnwswnwenweneswneeneswnw
nenwenenewseneweewswneenesweene
swwwwwwwnwswew
seesenwesesesewnwsenweseseswseswsee
nenenwnwnenewnwnwnenwseswnwnwnwnenwseswse
neneeeenwwweseseswewseesenwswnw
enewesenenesenwwsesweneeneeneenenee
seeeeeneewseewseseeeeesesewnw
wwswswswswwwnwwewswsw
wswnwnenweswnwnwnwnwnwnwsenwewse
nwnwnwnwswnenenwenwenwnwneseswnwnwnenwne
ewswswswwnwswswwswnwswswswswwwsesenwne
wnwsenesenwsenwesewewseswnwnenwneswne
nesweenenwneeneseeeneneseenewneenw
nwswswswesenwsenwesweseswsewswesenwsw
sweswneneeneswnwesenwwnwseeseswwnw
enwnwnwnwnenwnwnwnenwwnwesewnwnwsenwnw
eneseweeeneneneswwseneswweswswsw
swswwwseneeseeewnwwseswnwswenwneswnw
senewwswsesenewnwseseswswswneseneswsesw
swnwnwnenwsenwsenwnwsenwnwwnwnwew
seseseseswneesenweseeseswsesenwsewsese
eewweeneneswswenenweneneeeesw
ewneneeeseeeneneesweswenweswene
wseeneseeswnwwesesesenwnwwsesee
sesesesewseseneseseswnesesesenwsesenwse
seseswseneneewseseseseseseesesee
swswsesenwneeseswsesesesesesesew
nenwseswswnesenewwnwenwseeeeneneenene
neswneswswnewnwswnwseneeneneenwenwne
swnwnwnwnwnwseenewnenenwnwnwnweewnwnw
swnwnewnwwnwnwwnewswwsewwnwswnewew
nwnenwswnenwnwnwnenesenwsenenwnenenesenwnw
nenwewseswneweesweseeenweswswnw
swsweswseeswwswswswneswwwnwnwswnwse
swwwswwsewenewsewswswwnewnwswnw
swenenwnweswnwswwswwwnwweswnewnw
nwnenwnwnwnwwenenwnwsenenewenwnwswnesw
neesewswnwswseswnwenenwwswswswswewse
seseseseeeseseseswnw
swswenwswnenwnwnenwnwnwsenenwnwnenenwnwnew
wseseeeseseneseswesesesesese
seswseseseseswneswnwwseseseswesw
nwnwnwnwenweswsesenwswnwnwnwnwnwewswnwne
swneneenenwnenenewneneeweeeeswnese
wwnwnewwwwwswwww
eseeeseeeeenwewseenwesenwwse
seswseswneswseswswswwseswsw
wneswseswswnwswneseswsweswneswwswswsesw
wswwnwnwnwnwnesweswseswseswenwwswew
neeeeswneeswnweeeeeeesweee
neswswnwswesenwsenesenwsewseewswswne
newseesesesenwsenwewsewsese
sewseenwseseseswseseeswseseseneenwse
weseneeeeeesewneseneeswewsew
wswnwnwsewnwnwenwewwnwsesenwnenwnwnw
senenenenenewnenwnwse
seseseseseswsesenesesesesewsenwseneese
swwneewwneswsewseseeneenwsenenese
swnwnwneswseswsweswewneswnewnweseswsw
swswswswseeswswnwseswswnw
nenwseneenwneneenwneneneswseseeneswenw
swneswswwswswseswswswswnweswewswswe
swnenenwneswenenewnwnwnenenenese
nwswswnewseswnwswnwwwesweeswesese
nwenwnwnwnenwnwnesenwnwnwswnwnwsesenese
sweewwnwswesewnenwswww
eweneeswwswneeneeeseesenenwnee
nwnwnwnwsewnwenwnwnweswwwnenwnwee
seswswswwswnwwseneswswswswneswwswnene
newnenenwnesenenesewneswnenwnwnwnenwne
sewnenenenwneeswwenwneswnese
swnwnwwnwnwnenwenenesenewneeswnenene
swnenenesenewneneswsewwwwneesesee
wewneneeneneswwwsenenesene
wenwswswswnwwenwwwwenewnwnwwww
wwewseswwwswswwwnewnewnwnwwwne
swswwswswnwswnwswseswswswnwswswswseeswne
wewswnweneeseseeweeeewwsene
swsewwsenwseswewneswsesee
seseswnwneneseesesesesesesenweewsesese
seseseseswnesesesenwswsesew
swnwseswewseswswswswswse
wwesenwnwnweswswnwneewswnwneewse
nwnwnwnwnwnesesenwnwsenewnwnwnwnwnwnwnw
wswwwnenwswseswewwswswwnwwswswsw
senenenewnenesenenenenenwenweese
swneneeneswenewswnenwnweneesee
neenenewneeewneneswseneseneneneene
swnewswswsenwwswseswnweswswnwsweswsw
ewwnwenwwwsewnwnwnwsesw
seswseswswswseswsesenwsese
neneswnweeeeeneeswwnwswseee
swnenwnenwnwneswenenenwnenenweneneswnee
nenenwnwneneneneswnenenwnwsenwnewesenwwe
wwnwnwnenwwwwewseeswwwnwnwswnwnw
neneeneeeeneneeeeewe
nwneswweswswnesweesenwswswswswsesewse
seewseneseneseseseswswewseeseenwsee
nwswseswswseseswseswseswse
swseewnwseswwwseeswswnesewneneese
nwneeseeeneeeeeeee
wwswenwwnwnwwwsenwwwwwnwsenew
weneneneeneneswne
nenwswneswswwswwwwwswenewwsesewsw
sesesenenwneswwenesenewesweswseeenw
sweewswwnwenwseseseenwsesw
enwneenwewsenwnwwsewwwsww
nwnweseneswneeneneneweswswwneseswnese
nweseswneeseenwnenenesesenwenenenwwe
nwwnwnwsenwnwnwenenwnwnwsenwswnwsenwwnw
wewwwwseswnwsewwwnenwwwwsenesw
seswseswswswswseseswnwswswesw
sweswwnewneneneneneseneneneenwenene
swwnwswswwwewswswenwwse
senwenwnwwnwesesenwwnwnwnenwsenwnwnwne
seseeeweseswwnenwsenesewnweseeese
nwwewwnwnenwswnwsewswwnewenwnwnwnw
eesweesenweewneeeeenwsweene
newswesenwnewswneswswswseswswnwswnesw
sesesewwsesewseseneeswsenesesenesese
nwnwwswnwnenwwnwwnwnwneseswesenwswnwse
wwnesewnwnesesewsewnwwneneswse
eeweneeneweneeseeewneeeene
neneswnenwnenenwnenenwnwnwene
wwwseswswneenewseswseswwswnenesese
nwnwnwnwseswwnwswnwnwsenwnwsenwsenenwnee
newswswwwwseseneswwswswnewswwww
seseseseneseseswswswseseesenesesewswnwse
eeseswweseeenwsenesesesesesewsesee
neenenenenenwnweesenwesenwenesenesee
nwnwnwnenwswnwnwsenwnwnwnenewnwnw
seswnewswswswswnwe
eswnewneeneneneenenenwnenewnwewswse
wnewnenenwnwseneneseswsewenwsesenwwwse
eeseseweneseseewnesewnesesesenw
seeneeeewseeeneewsesweeseee
seseseswswsenwneeswnwswseseseswseswswnese
sewwswswseneswnwwwwswwne
nwnwnwnenwnwnwwnwnweswnwnwswenw
seswnwseseswnwnwseseenenw
newsewswnwwwneswsewneeeseswnwwenenw
neneeswnenenewseesenewe
seswneswsesewseeswse
sesenwnewneneeeneneneenenenesewnenenenw
wwneneswewswswsew
sewwwnwnwnwwsenwnenwsesenww
nwwnwnenenwwnwwswswswewseneesenewwnw
eeswswneneneneeneeeeseeewneee
wsewswwwnenwswnwswwsenenwenenw
nwwnwwwwnwewnewwwsenwswwswenwne
swsenesenwsesesewseswneeseseseseswsenw
wswswswseeswswswswnwswneweseswswswswnw
wnwnwwnwnwwnwnwwnwwswnwwenwwese
swseesenwswnwsenwswsweeweenwneeenene
neswnenewneneneneneneneneswnesenenesew
wnwneenenewnenenesenwne
senwnwseenesewseswwseswsweswsenenwsee
nwnwnwwnwenwsweswwnwseenwwnwnwnwnw
neseneneneswnwswsenewnewneenwsewnenenese
nwseseswwswseesenesenese
swsenwsewwneesewweswnenwneenwwnwse
eswswwswnenewseseswnwseswswnewswswsesw
wewneseeswwwnwneewnwswwswswwsesw
swseeewseswwswswneswswnwswsw
seewswwswwswsenwnwswwswnewswswswsww
seswswswsesesenewswswsenwswswswneswswsw
enwwwwswswwwswsewswnwnwseeswnwwsw
seneseneswseswseseswwwsweseswnwesesese
eswnesweswswswswnwswswswswwwsw
eeeeweeseseneewe
swnenwenwnwnenesewwswnwnwnwnwne
seswneswswwwsewenewwswwswwwww
seswseseswswswswswsenwsesenwneeewseswswse
wwnwnwwwenwwnwnww
seswseeseeseeneesesesewnese
wnwnwwnesewwwswewsesewnwewwe
swnwseswswswnweswneweswswswswswswnesw
swseseseseseswswseswne
nwswwwwnwswnwsesewenewwsesenwnenw
swwnwsweeewnwnwwsweseswsw
nwnwnwnweswnwsenwnwnwnwnwnwnwnww
eswnenwswnwneeneee
enweneeeeeeesenenee
nwwnwweneeseenenenwnwnwwswswnenenw
eeeneseesweeeewenwesweeee
swswnewswwsesenwwswswwwwewneswswnesw
seseseseseseseesesewesesewe
neswwwwwneswwswwwwsewneneww
wsewwwwwwnwnwnwwnwnww
nenwneneneneneseneneenenenewneswnwesene
swnwseseswswswseweswsweswswnwnesewseesw
weswswneswseewwnwnwneseesese
eewwnwnweswww
swwwwwwwswsewwewswnewwneww
senesesesenwnwnwnenewsenwnwnewnw
nwnenwnwnenenwewseenwnwnenwnwwnewnw
wwwweswwseneneseenwseswnwwnwwwnw
nwnwwswnenwnwnwwwseewswnwnenweww
neswnwnewneneeneneneneneneneneseswnesene
eeenwewseeenwsewnwsewe
nwnwwnwwnesenwnwnwnwnwnwsenwnwswnenwnw
nwseswswswnewnenwnwneneenenwwneneeesw
nwnwnwnwwswsenweenwnweenwnwnwswnwnww
nwesweeeeenweeseweeeeewee
swseseswseswswswswnwnwswswsweneswswsese
swesesweeswswswnwwswswswsewse
wswswweseswswswenwswnwswswnwwwwsw
wsesesesenwswswnwseseseeseseesesesese
wsenewwwwwwwwnewsew
ewswwwwswwwwwnwweswww
enwwnwnenwesewne
wswswseswsesesenesesesese
swseseseswseseseswswnewe
swswnwnenenwswneneneneenenenenenwswnenw
eneeseeseenewseeseewe
nwneseswswswswswnwswswswswwneeseswswne
swswsesesesenwseseesenew
newnwsenwsenenwnwnenwnwnwnesenwwnwnw
eneenenewwneneneneneneseneenesene
wnwseweeenwswsewwneswwneenewewe
nwnwnwwwnwwneenwsenenwseswnwswwnwnw
nwnwswneneswnwnwnenwwnweswswwnwnenwnw
neneenweseswseeewneweneenweseee
eeneweseeeeenwsenweneeeeswswe
ewsewswnwwwwnwneewseswneenewswwe
eeweeesweeeneeeeswswseenwnwee
nwnweswswsewwsenesweswwnwnwenwswe
wwewwsesenwswnwwsenwnenwwnwsewwnw
eewsewwnewwwwneswwwswswswww
enwewswesenweenwneewseswseeee
sesenewswwwweneneseeswnwnenwneee
nwewnwsewnwnenwwnewwse
eseeseseeesesesenwseneswneneswesew
wswwseenewswenwwweswswnenww
neseswseseeswswswwswswnwsenwsweseswswse
eeswwenwnenwneeweseweswneswene
nenwnenwnwnwneneswnenenenwneenenwswseswnw
swneeeeseenweneseeweswneesenwenwsw
nwnwnwsenwwnwsenwnwne
newwnwsenwnewsewswnenwswwenenwnesese
senwsenwnesenwnewnenwne`
