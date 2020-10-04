Sudoku solvers
---

Exploring different Sudoku puzzle solving algorithms, using Go. 

# Techniques

## Backtracking / Bruteforce

> A brute force algorithm visits the empty cells in some order, filling in digits sequentially, or backtracking when the number is found to be not valid.
> <br>[Wikipedia](https://en.wikipedia.org/wiki/Sudoku_solving_algorithms#Backtracking)

Backtracking is one of the first solutions you'd think of, but it's highly inefficient. 
As long as the puzzle is valid, a solution will be found. Even complex puzzles are solved within seconds using modern day technology.
 The algorithm is fairly simple as you just look up collisions for a single cell at a time, iterating over the possibilities
 until progressing.
  

# Build / Use

## Requirements
* Docker

Execute  `build.sh` to build the binaries for **Linux** & **Mac**. They will be placed in the project's `./bin` directory.

## Usage

Write your Sudoku input in a JSON file following the [default example](examples/default.json). 
Empty cells must be filled by a zero (0). Input sanitation hasn't been implemented! 

`./bin/bruteforce -input examples/default.json` 