package astar

import "fmt"

type Object struct {
	Row       int
	Column    int
	Value     int
	Cost      int
	Score     int
	Distance  int
	Option    string
	Option2   string
	Neighbors []*Object
}

type Maze struct {
	Plane     []*map[int]*Object
	BaseCost  int
	BaseValue int
	rows      int
	columns   int
}

// the maze is an array of maps.
// No controls on inputs (e.g. minus integers will result in runtime exception)
func Init(baseValue, baseCost, rows int, columns int) *Maze {
	//startX, startY, endX, endY,
	plane := make([]*map[int]*Object, rows, columns)
	for i := 0; i < rows; i++ {
		r := make(map[int]*Object)
		for j := 0; j < columns; j++ {
			r[j] = &Object{
				Row:    i,
				Column: j,
				Value:  baseValue,
				Cost:   baseCost,
			}
		}
		plane[i] = &r
	}
	return &Maze{
		Plane:     plane,
		BaseCost:  baseCost,
		BaseValue: baseValue,
		rows:      rows,
		columns:   columns,
	}
}

func (m *Maze) SetObject(objects *[]*Object) {
	for _, o := range *objects {
		r := *m.Plane[o.Row]
		r[o.Column] = o
	}
}

// maps would be nice here as deletion of independent items would be easily possible
func (m *Maze) DelRows(r int) {

}

// Adds rows with baseInt. Can later be modified
func (m *Maze) AddRows(r int) {
	for i := m.rows - 1; i < r; i++ {
		row := make(map[int]*Object)
		for j := 0; j < m.columns; j++ {
			row[j] = &Object{
				Row:    i,
				Column: j,
				Value:  m.BaseValue,
				Cost:   m.BaseCost,
			}
		}
		m.Plane[i] = &row
	}
}

func (m *Maze) Field(row int, column int) *Object {
	if row < 0 || row > m.rows-1 || column < 0 || column > m.columns-1 {
		return nil
	}
	r := *m.Plane[row]
	return r[column]
}

func (m *Maze) PrintValues() {
	for i := m.rows - 1; i > -1; i-- {
		r := *m.Plane[i]
		for j := 0; j < m.columns; j++ {
			fmt.Print(r[j].Value)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

// implement this with typewriter for column spacing
func (m *Maze) PrintCosts() {
	for i := 0; i < m.rows; i++ {
		r := *m.Plane[i]
		for j := 0; j < m.columns; j++ {
			fmt.Print(r[j].Cost)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func (o *Object) isWalkable() bool {
	if o.Cost > 0 {
		return true
	}
	return false
}
