package main

import (
	"fmt"
	"math/rand"
	"time"
)

const BoardSize int = 70

type Cell struct {
	x, y             int
	thisGen, nextGen bool
}

type Grid [BoardSize][BoardSize]*Cell

func RandomBool() bool {
	return rand.Float64() < 0.3
}

func (c *Cell) SetNextGen(grid Grid) {
	/*
	 Any live cell with fewer than two live neighbours dies, as if caused by under-population.
	 Any live cell with two or three live neighbours lives on to the next generation.
	 Any live cell with more than three live neighbours dies, as if by overcrowding.
	 Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
	*/
	neighbours_alive := c.GetNeighboursAlive(grid)
	c.nextGen = (!c.thisGen && neighbours_alive == 3) || (c.thisGen && 1 < neighbours_alive && neighbours_alive < 4)
}

func (c Cell) GetNeighboursAlive(grid Grid) int {
	neighbours_alive := 0
	for _x := -1; _x <= 1; _x++ {
		for _y := -1; _y <= 1; _y++ {
			x := c.x + _x
			y := c.y + _y
			inbound := x > 0 && y > 0 && x < BoardSize && y < BoardSize
			if inbound && grid[x][y].thisGen {
				neighbours_alive++
			}
		}
	}
	return neighbours_alive
}

func (cell Cell) PrintState() {
	if cell.thisGen {
		fmt.Printf("o")
	} else {
		fmt.Printf(".")
	}
}

func (grid Grid) Play() {
	for _, row := range grid {
		for _, cell := range row {
			cell.PrintState()
			cell.SetNextGen(grid)
		}
		fmt.Printf("\n")
	}
	for _, row := range grid {
		for _, cell := range row {
			cell.thisGen = cell.nextGen
		}
	}
}

func MakeGrid() *Grid {
	grid := new(Grid)
	rand.Seed(time.Now().UTC().UnixNano())
	for x, row := range grid {
		for y, _ := range row {
			c := &Cell{x, y, RandomBool(), RandomBool()}
			grid[x][y] = c
		}
	}
	return grid
}

func main() {
	grid := MakeGrid()
	for i := 0; i < 300; i++ {
		fmt.Print("\033[H\033[2J")
		grid.Play()
		time.Sleep(time.Second / 2)
	}
}
