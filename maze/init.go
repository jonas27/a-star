package maze

import "fmt"

type Objects map[int]interface{}

type Maze struct {
	Field   map[int]map[int]int
	Objects map[*map[int]int]int
	Width   int
	Height  int
}

// the maze is a maps of maps.
func Init(baseInt, rows int, columns int, objects map[*map[int]int]int) {
	//startX, startY, endX, endY,
	k := make([]*map[int]int, rows, columns)
	for i := 0; i < rows; i++ {
		r := make(map[int]int)
		for j := 0; j < columns; j++ {
			r[j] = baseInt
		}
		k[i] = &r
	}
	for _, o := range k {
		fmt.Println(o)
	}
}

func AddObjects() {

}
