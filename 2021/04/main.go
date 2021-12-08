package main

import (
	"fmt"
	"strings"

	"github.com/clockworksoul/adventofcode"
)

type Row struct {
	Numbers []int
	Draws   []int
}

func (r *Row) Draw(n int) bool {
	if r.Has(n) != -1 {
		r.Draws = append(r.Draws, n)
		return true
	}
	return false
}

func (r *Row) Has(n int) int {
	for i, v := range r.Numbers {
		if v == n {
			return i
		}
	}
	return -1
}

func (r *Row) Winner() bool {
	var count int

	for i := 0; i < len(r.Draws) && count < len(r.Numbers); i++ {
		if r.Has(r.Draws[i]) != -1 {
			count++
		}
	}

	return count >= len(r.Numbers)
}

func (r *Row) String() string {
	if len(r.Numbers) == 0 {
		return ""
	}

	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("%2d", r.Numbers[0]))

	for i := 1; i < len(r.Numbers); i++ {
		b.WriteString(fmt.Sprintf(" %2d", r.Numbers[i]))
	}

	return b.String()
}

type Board struct {
	Rows    []*Row
	Columns []*Row
}

func NewBoard(size int) *Board {
	b := &Board{
		Rows:    make([]*Row, 0),
		Columns: make([]*Row, size),
	}

	for i := range b.Columns {
		b.Columns[i] = &Row{}
	}

	return b
}

func (b *Board) AddRow(r *Row) {
	b.Rows = append(b.Rows, r)

	for i, c := range b.Columns {
		c.Numbers = append(c.Numbers, r.Numbers[i])
	}
}

func (b *Board) Draw(n int) bool {
	var drawn bool

	for _, r := range b.Rows {
		if r.Draw(n) {
			drawn = true
			break
		}
	}

	if !drawn {
		return false
	}

	for _, c := range b.Columns {
		if c.Draw(n) {
			drawn = true
			break
		}
	}

	return drawn
}

func (b *Board) Has(n int) int {
	for i, v := range b.Rows {
		if v.Has(n) != -1 {
			return i
		}
	}
	return -1
}

func (b *Board) Score(n int) int {
	m := make(map[int]interface{})
	for _, r := range b.Rows {
		for _, n := range r.Numbers {
			m[n] = true
		}
	}

	for _, r := range b.Rows {
		for _, n := range r.Draws {
			delete(m, n)
		}
	}

	var sum int

	for k := range m {
		sum += k
	}

	return sum * n
}

func (b *Board) String() string {
	if len(b.Rows) == 0 {
		return ""
	}

	bb := strings.Builder{}

	for _, r := range b.Rows {
		bb.WriteString(r.String())
		bb.WriteString("\n")
	}

	return bb.String()
}

func (b *Board) Winner() bool {
	// Check all the rows
	for _, r := range b.Rows {
		if r.Winner() {
			return true
		}
	}

	// Check all the columns
	for _, c := range b.Columns {
		if c.Winner() {
			return true
		}
	}

	return false
}

func main() {
	var draws []int
	var boards []*Board
	var board *Board

	adventofcode.IngestFile("input.txt", func(s string) {
		s = strings.TrimSpace(s)

		switch {
		case draws == nil:
			draws = adventofcode.MustSplitAndParseInts(s, ",")
		case s == "":
			if board != nil {
				boards = append(boards, board)
			}
			board = NewBoard(5)
		default:
			r := adventofcode.MustSplitAndParseInts(s, " +")
			row := &Row{Numbers: r}
			board.AddRow(row)
		}
	})
	if board.Rows != nil {
		boards = append(boards, board)
	}

	part1(boards, draws)
	part2(boards, draws)
}

func part1(boards []*Board, draws []int) {
	for _, d := range draws {
		for _, b := range boards {
			b.Draw(d)

			if b.Winner() {
				fmt.Println("Part 1:", b.Score(d))
				return
			}
		}
	}
}

func part2(boards []*Board, draws []int) {
	var d int
	var winners = make(map[int]*Board)

	for i := range boards {
		winners[i] = boards[i]
	}

	for _, d = range draws {
		for i, b := range boards {
			// Already a winner? Ignore.
			if winners[i] == nil {
				continue
			}

			b.Draw(d)

			if b.Winner() {
				if len(winners) == 1 {
					for _, b := range winners {
						fmt.Println("Part 2:", b.Score(d))
					}
					return
				} else {
					delete(winners, i)
				}
			}
		}
	}
}
