package grid

import (
	"../fileops"
	"bytes"
)

type Grid struct {
	tiles [][]bool
	Width, Height int
}

type Pos struct {
	Row, Column int
}

func toGrid(lines []string) *Grid {

	// Declare the grid variables and allocate the top level slice
	grid := &Grid{}
	grid.Height = len(lines)
	grid.Width = len(lines[0])
	grid.tiles = make([][]bool, grid.Height)

	for lineIdx, line := range lines {
		// Allocate space for each line
		grid.tiles[lineIdx] = make([]bool, grid.Width)
		for tileIdx, tile := range line {
			switch string(tile) {
			case ".":
				grid.tiles[lineIdx][tileIdx] = false
			case "O":
				grid.tiles[lineIdx][tileIdx] = true
			}
		}
	}
	return grid
}

func NewGrid(path string) (*Grid, error) {
	lines, err := fileops.ReadLines(path)
	return toGrid(lines), err
}

func (src *Grid) Clone() *Grid {
	tgt := *src
	return &tgt
}

func (src *Grid) NextState(row int, column int) bool {
	aliveNeighbours := 0
	var decision bool

	for i := row - 1; i <= row+1; i++ {
		for j := column - 1; j <= column+1; j++ {
			// Ignore the element itself
			if i == row && j == column {
				continue
			}
			// Ignore positions outside of the boundaries
			if i < 0 || j < 0 || i > src.Height-1 || j > src.Width-1 {
				continue
			}
			// Increase the counter if this tile is "alive"
			if src.tiles[i][j] == true {
				aliveNeighbours++
			}
		}
	}

	// Return the new state based on the number of alive neighbours
	if aliveNeighbours == 3 {
		decision = true
	} else if aliveNeighbours == 2 {
		decision = src.tiles[row][column]
	} else {
		decision = false
	}

	return decision
}

func (tgt *Grid) SetTile(row int, column int, val bool){
	tgt.tiles[row][column] = val
}

func (g Grid) String() string {

	var buf bytes.Buffer

	for _, line := range g.tiles {
		for _, tile := range line {
			if tile {
				buf.WriteByte('X')
			} else {
				buf.WriteByte(' ')
			}
		}
		buf.WriteByte('\n')
	}
	return buf.String()

}
