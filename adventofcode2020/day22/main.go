package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	part2(RawInput)
}

func part1(input string) {
	p1, p2 := parseInput(input)
	winningHand := getWinningHand(newGameContext(), p1, p2)
	fmt.Println(getScore(winningHand))
}
func part2(input string) {
	p1, p2 := parseInput(input)
	winner, winningHand := getWinningHand2(1, newGameContext(), p1, p2)
	fmt.Printf("The winner of game 1 is player %d!\n", winner)
	fmt.Println(getScore(winningHand))
}

func getScore(h hand) int {
	m := len(h)
	sum := 0
	for i, card := range h {
		sum += (m - i) * card
	}
	return sum
}

func getWinningHand2(gameId int, ctx gameContext, p1 hand, p2 hand) (int, hand) {
	if ctx.round == 1 {
		fmt.Printf("== Game %d ==\n\n", gameId)
	}
	if ctx.checkRecursion(p1, p2) {
		return 1, p1
	}
	fmt.Printf("-- Round %d (Game %d) -\n-", ctx.round, gameId)
	fmt.Printf("Player 1's deck: %v\n", p1)
	fmt.Printf("Player 2's deck: %v\n", p2)
	c1, c2 := p1.pop(), p2.pop()
	fmt.Printf("Player 1 plays: %d\n", c1)
	fmt.Printf("Player 2 plays: %d\n", c1)

	winner := 0
	if len(p1) >= c1 && len(p2) >= c2 {
		fmt.Println("Playing a subgame to determine the winner...")
		winner, _ = getWinningHand2(gameId+1, newGameContext(), p1.copy(c1), p2.copy(c2))
	} else if c1 > c2 {
		winner = 1
	} else {
		winner = 2
	}

	if winner == 1 {
		p1.push(c1)
		p1.push(c2)
	} else {
		p2.push(c2)
		p2.push(c1)
	}
	fmt.Println(p1, p2)
	if len(p1) == 0 {
		fmt.Printf("Player 2 wins round %d of game %d!\n", ctx.round, gameId)
		if gameId > 1 {
			fmt.Printf("...anyway back to game %d!\n", gameId-1)
		}
		return 2, p2
	} else if len(p2) == 0 {
		fmt.Printf("Player 1 wins round %d of game %d!\n", ctx.round, gameId)
		if gameId > 1 {
			fmt.Printf("...anyway back to game %d!\n", gameId-1)
		}
		return 1, p1
	}
	ctx.round++
	return getWinningHand2(gameId, ctx, p1, p2)
}

func getWinningHand(ctx gameContext, p1 hand, p2 hand) hand {
	fmt.Printf("-- Round %d --\n", ctx.round)
	fmt.Printf("Player 1's deck: %v\n", p1)
	fmt.Printf("Player 2's deck: %v\n", p2)
	c1, c2 := p1.pop(), p2.pop()
	fmt.Printf("Player 1 plays: %d\n", c1)
	fmt.Printf("Player 2 plays: %d\n", c2)

	if c1 > c2 {
		p1.push(c1)
		p1.push(c2)
	} else {
		p2.push(c2)
		p2.push(c1)
	}
	if len(p1) == 0 {
		fmt.Printf("Player 2 wins the round!\n\n")
		return p2
	} else if len(p2) == 0 {
		fmt.Printf("Player 1 wins the round!\n\n")
		return p1
	}
	ctx.round++
	return getWinningHand(ctx, p1, p2)
}

func parseInput(input string) (p1 hand, p2 hand) {

	parsePlayer := func(rows []string) hand {
		h := hand{}
		for _, card := range rows[1:] {
			val, _ := strconv.Atoi(card)
			h = append(h, val)
		}
		return h
	}
	players := strings.Split(input, "\n\n")

	p1 = parsePlayer(strings.Split(players[0], "\n"))
	p2 = parsePlayer(strings.Split(players[1], "\n"))
	return
}

type hand []int

func (h *hand) push(card int) {
	*h = append(*h, card)
}

func (h *hand) pop() int {
	ref := *h
	card := ref[0]
	ref = ref[1:]
	*h = ref
	return card
}

func (h *hand) copy(num int) hand {
	newHand := make(hand, num)
	for i := 0; i < num && i < len(*h); i++ {
		newHand[i] = (*h)[i]
	}
	return newHand
}

type gameContext struct {
	round int
	p1    map[string]bool
	p2    map[string]bool
}

func (ctx *gameContext) checkRecursion(p1 hand, p2 hand) bool {
	s1, s2 := fmt.Sprintf("%v", p1), fmt.Sprintf("%v", p2)
	_, f1 := ctx.p1[s1]
	_, f2 := ctx.p2[s2]
	ctx.p1[s1] = true
	ctx.p2[s2] = true
	return f1 && f2
}

func newGameContext() gameContext {
	return gameContext{1, make(map[string]bool), make(map[string]bool)}
}

const DemoInput = `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`

const InfiniteGame = `Player 1:
43
19

Player 2:
2
29
14`

const RawInput = `Player 1:
19
5
35
6
12
22
45
39
14
42
47
38
2
26
13
30
4
34
43
40
16
8
23
50
36

Player 2:
1
21
29
41
32
28
9
37
49
20
17
27
24
3
33
44
48
31
15
25
18
46
7
10
11`
