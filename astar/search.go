package astar

var (
	//visistedNodes *[]Object
	endR int
	endC int
)

func search(startRow int, startColumn int, endRow int, endColumn int, m *Maze) {
	endR = endRow
	endC = endColumn
	start := Object{
		Row:      startRow,
		Column:   startColumn,
		Score:    abs(startRow - endR + startColumn - endC),
		Distance: abs(startRow - endR + startColumn - endC),
	}
	start.scanhNeighbors(m)
}

func (o *Object) scanhNeighbors(m *Maze) {
	up := m.Field(o.Row+1, o.Column)
	if up != nil && up.isWalkable() && up.Distance+up.Cost < up.Score {
		up.Score = up.Distance + up.Cost
		o.Neighbors = append(o.Neighbors, up)
	}
	right := m.Field(o.Row, o.Column+1)
	if right != nil && right.isWalkable() && right.Distance+right.Cost < right.Score {
		right.Score = right.Distance + right.Cost
		o.Neighbors = append(o.Neighbors, right)
	}
	down := m.Field(o.Row-1, o.Column)
	if down != nil && down.isWalkable() && down.Distance+down.Cost < down.Score {
		down.Score = down.Distance + down.Cost
		o.Neighbors = append(o.Neighbors, down)
	}
	left := m.Field(o.Row, o.Column-1)
	if left != nil && left.isWalkable() && left.Distance+left.Cost < left.Score {
		left.Score = left.Distance + left.Cost
		o.Neighbors = append(o.Neighbors, left)
	}
	o.sortNeighbors()
}

func (o *Object) sortNeighbors() {
	//sort.Float64s()
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
