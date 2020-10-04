package sudoku

const EmptyCell int = 0

type Sudoku interface {
	// input is present in column
	InCol(col, i int) bool
	// input is present in row
	InRow(row, i int) bool
	// input is present in box
	InBox(rowStart, colStart, i int) bool
	// puzzle has empty cell
	HasEmptyCell(rowStart, colStart int) bool
	// get first empty cell
	FindEmptyCell(rowStart, colStart int) (r int, c int)
	// get last empty cell
	FindEmptyCellRev(rowStart, colStart int) (r int, c int)
	// get cell value
	Get(row, col int) int
	// set cell value
	Set(row, col, i int)
	// unset cell value
	Unset(row, col int)
	// print grid
	Print()
}
