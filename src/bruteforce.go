package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"vpmv/sudoku/stopwatch"
	"vpmv/sudoku/sudoku"
)

var puzzleType string
var puzzleInput string

var puzzle sudoku.Sudoku

// check if number is not present in row/col/box
func isValid(row, col, i int) bool {
	return !puzzle.InRow(row, i) && !puzzle.InCol(col, i) && !puzzle.InBox(row-row%3, col-col%3, i)
}

// solve a default Sudoku puzzle
func solveDefault() bool {
	var (
		row int
		col int
	)

	row, col = puzzle.FindEmptyCell(row, col)
	// all cells have been filled
	if row == -1 && col == -1 {
		return true
	}

	// try number in current cell
	for num := 1; num <= 9; num++ {
		if isValid(row, col, num) {
			// try solving the next cell / rest of the puzzle using current number
			puzzle.Set(row, col, num)
			if solveDefault() {
				return true
			}
			// invalid number
			puzzle.Unset(row, col)
		}
	}
	return false
}

// entrypoint
func main() {
	// switch puzzleType {
	// default:
	// }

	stopwatch.Start()
	if solveDefault() {
		stopwatch.Stop()
		puzzle.Print()

		stopwatch.Print()
	} else {
		fmt.Println(`No solution found!`)
	}
}

// initialize cmd
func init() {
	flag.StringVar(&puzzleInput, `input`, `sudoku.json`, `Sudoku input file`)
	flag.StringVar(&puzzleType, `type`, `default`, `Sudoku type`)
	flag.Parse()

	b, err := ioutil.ReadFile(puzzleInput)
	if os.IsNotExist(err) {
		fmt.Println(`Couldn't find input file!`)
		os.Exit(1)
	} else if err != nil {
		fmt.Println(`Error reading input file: `, err)
		os.Exit(1)
	}

	// initialize puzzle according to type/input
	switch puzzleType {
	default:
		fmt.Println(`Unsupported puzzle type!`)
		os.Exit(1)
	case `default`:
		var grid [][]int
		if err := json.Unmarshal(b, &grid); err != nil {
			fmt.Println(`Error decoding input: `, err)
			os.Exit(1)
		}
		puzzle = sudoku.NewDefault(grid)
	}
}
