package astar

import (
	"log"
	"testing"
)

func TestInit(*testing.T) {
	maze := Init(0, 1, 7, 7)
	for _, o := range maze.Plane {
		log.Println(o)
	}
}

func TestMaze_SetObjects(t *testing.T) {
	maze := Init(0, 1, 7, 7)
	maze.SetObject(&[]*Object{
		// walls
		{0, 1, 1, 1},
		{1, 1, 1, 1000},
		{1, 2, 1, 1000},
		{1, 4, 1, 1000},
		{1, 6, 1, 1000},
		{1, 7, 1, 1000},
		{1, 8, 1, 1000},
		{2, 4, 1, 1000},
		{2, 8, 1, 1000},
		{3, 2, 1, 1000},
		{3, 3, 1, 1000},
		{3, 4, 1, 1000},
		{3, 6, 1, 1000},
		{3, 7, 1, 1000},
		{3, 8, 1, 1000},
		{5, 1, 1, 1000},
		{5, 2, 1, 1000},
		{5, 4, 1, 1000},
		{5, 5, 1, 1000},
		{5, 56, 1, 1000},
		{5, 58, 1, 1000},
		{6, 2, 1, 1000},
		// start
		{2, 2, 2, 1},
		// end
		{4, 4, 3, 1},
	})
	maze.PrintValues()
	//maze.PrintCosts()
}
