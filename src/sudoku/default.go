package sudoku

import "fmt"

// DefaultPuzzle is a default Sudoku puzzle, holding 9x9 cells
type DefaultPuzzle struct {
	Grid [][]int
	RowP int
	ColP int
	Size int
}

func NewDefault(rows [][]int) Sudoku {
	d := &DefaultPuzzle{
		Size: 9,
		Grid: make([][]int, 9),
	}
	for i := 0; i < d.Size; i++ {
		d.Grid[i] = rows[i]
	}

	return d
}

func (d DefaultPuzzle) InCol(col, i int) bool {
	for row := 0; row < d.Size; row++ {
		if d.Grid[row][col] == i {
			return true
		}
	}
	return false
}

func (d DefaultPuzzle) InRow(row, i int) bool {
	for col := 0; col < d.Size; col++ {
		if d.Grid[row][col] == i {
			return true
		}
	}
	return false
}

func (d DefaultPuzzle) InBox(rowStart, colStart, i int) bool {
	if rowStart%3 != 0 {
		panic(`Invalid starting row`)
	} else if colStart%3 != 0 {
		panic(`Invalid starting col`)
	}

	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if d.Grid[row+rowStart][col+colStart] == i {
				return true
			}
		}
	}
	return false
}

func (d DefaultPuzzle) HasEmptyCell(rowStart, colStart int) bool {
	r, c := d.FindEmptyCell(rowStart, colStart)
	return r == -1 && c == -1
}

func (d DefaultPuzzle) FindEmptyCell(rowStart, colStart int) (r int, c int) {
	for row := rowStart; row < d.Size; row++ {
		for col := colStart; col < d.Size; col++ {
			if d.Grid[row][col] == EmptyCell {
				return row, col
			}
		}
	}
	return -1, -1
}

func (d DefaultPuzzle) FindEmptyCellRev(rowEnd, colEnd int) (r int, c int) {
	for row := d.Size; row > rowEnd; row-- {
		for col := d.Size; col > colEnd; col-- {
			if d.Grid[row][col] == EmptyCell {
				return row, col
			}
		}
	}
	return -1, -1
}

func (d DefaultPuzzle) Get(row, col int) int {
	return d.Grid[row][col]
}

func (d *DefaultPuzzle) Set(row, col, i int) {
	d.Grid[row][col] = i
}

func (d *DefaultPuzzle) Unset(row, col int) {
	d.Grid[row][col] = EmptyCell
}

func (d DefaultPuzzle) Print() {
	fmt.Println(`-------------`)
	for row := 0; row < d.Size; row++ {
		if row > 0 && row%3 == 0 {
			fmt.Print("----|---|----\n")
		}
		for col := 0; col < d.Size; col++ {

			if col%3 == 0 {
				fmt.Print(`|`)
			}
			fmt.Print(d.Grid[row][col])
			if col == 8 {
				fmt.Print(`|`)
			}
		}
		fmt.Print("\n")
	}
	fmt.Println(`-------------`)
}
