package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMakeGrid(t *testing.T) {
	grid := MakeGrid()
	assert.Equal(t, grid[3][2].x, 3)
	assert.Equal(t, grid[3][2].y, 2)
}


func TestGetNeighboursAlive(t *testing.T) {
	grid := MakeGrid()
	grid[0][0].thisGen = true
	grid[0][1].thisGen = false
	grid[1][0].thisGen = false
	grid[1][1].thisGen = true

	assert.Equal(t, grid[0][0].GetNeighboursAlive(*grid), 1)
}


