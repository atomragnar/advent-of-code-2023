package seven

import (
	"bufio"
	"container/heap"
	"fmt"
	util2 "github.com/atomragnar/advent-of-code-2023/day/util"
	"io"
	"log"
	"log/slog"
	"strconv"
	"strings"
)

func Solution(isPartTwo bool) {
	day := 7
	var fn util2.BufferProcessor
	if isPartTwo {
		fn = partTwo
	} else {
		fn = partOne
	}
	err := util2.ProcessInput(util2.DataPath(day), fn)
	if err != nil {
		return
	}
}

func HandValue(cards []Card, max int) int {
	l := len(cards)
	if l == 1 {
		return 12
	}
	j := 0

	for _, c := range cards {
		if c.isJ {
			j = c.amount
		}
	}

	if l == 2 {
		if max == 4 {
			if j > 0 {
				return 12
			}
			return 11
		} else {
			if j >= 2 {
				return 12
			} else if j > 0 {
				return 11
			} else {
				return 10
			}

		}

	}

	if l == 3 {

		if max == 3 {
			if j == 2 {
				return 12
			} else if j > 0 {
				return 11
			} else {
				return 9
			}

		} else {

			if j >= 2 {
				return 11
			} else if j == 1 {
				return 10
			} else {
				return 8
			}
		}

	}

	if l == 4 {
		if j > 0 {
			return 9
		}

		return 2
	}

	if l == 5 {
		if j > 0 {
			return 2
		}
		return 1
	}

	return 0
}

type Card struct {
	isJ    bool
	amount int
	value  int
}

func NewCard(value int) Card {
	return Card{
		amount: 1,
		value:  value,
		isJ:    false,
	}
}

func (c Card) increment() {
	c.amount++
}

type CardHand struct {
	c      string
	s      int
	v      []int
	max    int
	points int
	cards  []Card
}

var valueMap = map[uint8]int{
	65: 13,
	75: 12,
	81: 11,
	74: 0,
	//74: 10,
	84: 9,
	57: 8,
	56: 7,
	55: 6,
	54: 5,
	53: 4,
	52: 3,
	51: 2,
	50: 1,
}

func NewCardHand(s string) CardHand {
	m := make(map[uint8]Card)
	c := CardHand{
		v: make([]int, 0),
	}
	c.c = s
	for i := 0; i < len(s); i++ {
		value := valueMap[s[i]]

		c.v = append(c.v, value)

		if _, ok := m[s[i]]; !ok {
			card := NewCard(value)
			if value == 0 {
				card.isJ = true
			}
			m[s[i]] = card
		} else {
			temp := m[s[i]]
			temp.amount++
			m[s[i]] = temp
		}
	}

	c.max = -1
	c.cards = make([]Card, 0)

	for _, v := range m {
		c.cards = append(c.cards, v)
		if c.max == -1 {
			c.max = v.amount
		} else {
			if v.amount > c.max {
				c.max = v.amount
			}
		}
	}

	c.s = HandValue(c.cards, c.max)

	if c.s == 12 {
		fmt.Println(c.c)
	}
	return c
}

func QueueItem(s []string) util2.QueueItem {
	c := NewCardHand(s[0])
	points, err := strconv.Atoi(strings.TrimSpace(s[1]))
	if err != nil {
		log.Fatal("Error converting to num", err.Error())
	}
	c.points = points
	q := util2.QueueItem{Value: c}
	return q
}

func Prio(this, other util2.QueueItem) bool {

	if this.Value.(CardHand).s > other.Value.(CardHand).s {
		return false
	}

	if other.Value.(CardHand).s > this.Value.(CardHand).s {
		return true
	}

	ts := this.Value.(CardHand).v
	os := other.Value.(CardHand).v
	//fmt.Println(ts, os)

	for i := 0; i < len(ts); i++ {

		if ts[i] > os[i] {
			return false
		}

		if os[i] > ts[i] {
			return true
		}
	}
	return false
}

func partTwo(reader *bufio.Reader) error {
	return partOne(reader)
}

func partOne(reader *bufio.Reader) error {
	var result int
	result = 0
	pq := util2.NewQueue(Prio)
	for {

		chunk, err := reader.ReadBytes('\n')

		if err != nil && err != io.EOF {
			slog.Error("Error reading file", "error", err)
			return err
		}

		d := strings.Split(string(chunk), " ")
		heap.Push(pq, QueueItem(d))

		if err == io.EOF {
			break
		}

	}
	var item interface{}
	i := 0
	for len(pq.Q) != 0 {
		i++
		item = heap.Pop(pq)
		fmt.Printf("Index: %d - Priority: %d\n", item.(util2.QueueItem).Index, item.(util2.QueueItem).Priority)
		fmt.Printf("Cards: %s - Strenght: %d - Max: %v value: %v \n", item.(util2.QueueItem).Value.(CardHand).c, item.(util2.QueueItem).Value.(CardHand).s, item.(util2.QueueItem).Value.(CardHand).max, item.(util2.QueueItem).Value.(CardHand).v)
		fmt.Printf("i: %d\n", i)
		p := i * item.(util2.QueueItem).Value.(CardHand).points
		result += p
		if item == nil {
			break
		}
	}

	fmt.Printf("Results is: %d", result)
	return nil
}
