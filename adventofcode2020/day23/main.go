package main

import (
	"fmt"
	"strings"
)

func main() {
	//demoInput := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	puzzleInput := []int{3, 6, 4, 2, 9, 7, 5, 8, 1}
	//part1(demoInput)
	part2(puzzleInput)
}

func part1(input []int) {
	crabCups := newCupList(input[0])
	crabCups.initialise(input[1:])
	play(crabCups, 100, true)

	current := crabCups.find(1).next
	var sb strings.Builder
	for i := 0; i < crabCups.size-1; i++ {
		fmt.Fprintf(&sb, "%d", current.value)
		current = current.next
	}
	fmt.Printf("Final: %s", sb.String())
}

func part2(input []int) {

	max := 0
	for _, v := range input {
		if v > max {
			max = v
		}
	}
	for i := max + 1; i <= 1000000; i++ {
		input = append(input, i)
	}

	crabCups := newCupList(input[0])
	crabCups.initialise(input[1:])
	play(crabCups, 10000000, false)
	current := crabCups.find(1)
	cup1 := current.next
	cup2 := cup1.next
	fmt.Println(cup1.value, cup2.value)
	fmt.Println(cup1.value * cup2.value)

}

func play(crabCups *cupList, turns int, doPrint bool) {
	current := crabCups.head
	turn := 1

	for turn <= turns {
		if doPrint {
			fmt.Printf("-- Move %d --\n", turn)
			fmt.Printf("cups : %s\n", crabCups.print(turn-1))
		} else if (turn % 1000000) == 0 {
			fmt.Printf("-- Move %d --\n", turn)
		}
		removedCups := make([]int, 0)
		for i := 0; i < 3; i++ {
			removedCups = append(removedCups, crabCups.delete(current.next))
		}
		if doPrint {
			fmt.Printf("Pick up: %v\n", removedCups)
		}
		foundDest := false
		objective := current.value
		var destCup *cup
		for !foundDest {
			objective--
			if objective <= 0 {
				objective = crabCups.maxCup + 1
			} else if !includes(removedCups, objective) {
				destCup, _ = crabCups.mapping[objective]
				foundDest = true
			}
		}
		if doPrint {
			fmt.Printf("destination: %d\n\n", destCup.value)
		}
		for _, cup := range removedCups {
			destCup = crabCups.add(destCup, cup)
		}
		current = current.next
		turn++
	}
}

func includes(removedCups []int, cup int) bool {
	for _, c := range removedCups {
		if c == cup {
			return true
		}
	}
	return false
}

type cup struct {
	value int
	next  *cup
	prev  *cup
}

func newCup(value int) *cup {
	return &cup{value, nil, nil}
}

type cupList struct {
	head    *cup
	size    int
	maxCup  int
	mapping map[int]*cup
}

func newCupList(headCup int) *cupList {
	head := newCup(headCup)
	size := 1
	mapping := map[int]*cup{
		headCup: head,
	}
	return &cupList{head, size, headCup, mapping}
}

func (cl *cupList) initialise(cups []int) {
	current := cl.head
	for _, c := range cups {
		cl.mapping[current.value] = current
		cup := newCup(c)
		current.next = cup

		prev := current
		current = current.next
		current.prev = prev
		cl.size++

		if current.value > cl.maxCup {
			cl.maxCup = current.value
		}
	}

	current.next = cl.head
	cl.head.prev = current
	cl.mapping[current.value] = current
}

func (cl *cupList) print(idx int) string {
	var sb strings.Builder
	curr := cl.head
	for i := 0; i < cl.size; i++ {
		if i == (idx % cl.size) {
			sb.WriteString(fmt.Sprintf("(%d) ", curr.value))
		} else {
			sb.WriteString(fmt.Sprintf("%d ", curr.value))
		}
		curr = curr.next
	}
	return sb.String()
}

func (cl *cupList) find(value int) *cup {
	curr := cl.head
	for i := 0; i < cl.size; i++ {
		if curr.value == value {
			return curr
		} else {
			curr = curr.next
		}
	}
	return nil
}

func (cl *cupList) delete(cupToDelete *cup) int {
	if cupToDelete == cl.head {
		cl.head = cl.head.next
	}
	prev := cupToDelete.prev
	next := cupToDelete.next
	prev.next = next
	next.prev = prev
	cl.size--
	delete(cl.mapping, cupToDelete.value)
	return cupToDelete.value
}

func (cl *cupList) max() int {
	max := 0
	curr := cl.head
	for i := 0; i < cl.size; i++ {
		if curr.value > max {
			max = curr.value
		}
		curr = curr.next
	}
	return max
}

func (cl *cupList) add(destCup *cup, cupToAdd int) *cup {
	temp := destCup.next

	c := newCup(cupToAdd)
	destCup.next = c
	c.prev = destCup
	temp.prev = c
	c.next = temp

	cl.size++
	cl.mapping[c.value] = c
	return c
}
