package main

import (
	"sort"
)

type Solver struct {
	problem    *Sudoku
	numGuesses int
}

func NewSolver(problem *Sudoku) *Solver {
	s := new(Solver)
	s.problem = problem
	s.numGuesses = 0
	return s
}

func (s *Solver) Solve() {
	s.recursiveBacktracking(s.problem)
}

func (s *Solver) recursiveBacktracking(problem *Sudoku) *Board {
	if problem.Complete() {
		return &problem.Board
	}
	tile, options := s.pickMinTile(problem.Board)
	for _, option := range s.orderedOptions(tile, options) {
		s.numGuesses++
		if s.consistent(tile, option, problem.Board) {
			problem.Board[tile.x][tile.y] = option
			result := s.recursiveBacktracking(problem)
			if result != nil {
				return result
			}
		}
		problem.Board[tile.x][tile.y] = 0
	}
	return nil
}

// Tile is a tuple for x-y locations in the sudoku grid
type Tile struct {
	x, y int
}

func (s *Solver) consistent(tile Tile, option int, board Board) bool {
	iunit := unit()
	for t := 0; t < 9; t++ {
		if board[tile.x][t] == option {
			return false
		}
		if board[t][tile.y] == option {
			return false
		}
	}
	for _, neighbor := range iunit[tile] {
		if board[neighbor.x][neighbor.y] == option {
			return false
		}
	}
	return true
}

func unit() map[Tile][]Tile {
	unit := make(map[int][]Tile)
	iunit := make(map[Tile][]Tile)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			var u int
			if (i >= 0 && i < 3) && (j >= 0 && j < 3) {
				u = 0
			} else if (i >= 0 && i < 3) && (j >= 3 && j < 6) {
				u = 1
			} else if (i >= 0 && i < 3) && (j >= 6 && j < 9) {
				u = 2
			} else if (i >= 3 && i < 6) && (j >= 0 && j < 3) {
				u = 3
			} else if (i >= 3 && i < 6) && (j >= 3 && j < 6) {
				u = 4
			} else if (i >= 3 && i < 6) && (j >= 6 && j < 9) {
				u = 5
			} else if (i >= 6 && i < 9) && (j >= 0 && j < 3) {
				u = 6
			} else if (i >= 6 && i < 9) && (j >= 3 && j < 6) {
				u = 7
			} else if (i >= 6 && i < 9) && (j >= 6 && j < 9) {
				u = 8
			}
			unit[u] = append(unit[u], Tile{i, j})
		}
	}
	for _, v := range unit {
		for _, i := range v {
			iunit[i] = v
		}
	}
	return iunit
}

func (s *Solver) pickMinTile(board Board) (Tile, map[Tile][]int) {
	options := make(map[Tile][]int)
	s.constrain(board, options)
	var tile Tile
	// pick a random key in options
	for k, _ := range options {
		tile = k
		break
	}
	for k, v := range options {
		if len(v) < len(options[tile]) {
			tile = k
		}
	}
	return tile, options
}

func (s *Solver) orderedOptions(tile Tile, options map[Tile][]int) ([]int) {
	sort.Ints(options[tile])
	return options[tile]
}

func (s *Solver) constrain(board Board, options map[Tile][]int) {
	// determine variables
	units := unit()
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				options[Tile{i, j}] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
			}
		}
	}
	for tile, _ := range options {
		// column
		for t := 0; t < 9; t++ {
			if board[tile.x][t] != 0 {
				options[tile] = remove(options[tile], board[tile.x][t])
			}
		}
		// row
		for t := 0; t < 9; t++ {
			if board[t][tile.y] != 0 {
				options[tile] = remove(options[tile], board[t][tile.y])
			}
		}
		// unit
		for _, t := range units[tile] {
			if board[t.x][t.y] != 0 {
				options[tile] = remove(options[tile], board[t.x][t.y])
			}
		}
		sort.Ints(options[tile])
	}
}

func remove(ints []int, i int) []int {
	var ret []int
	for _, v := range ints {
		if v != i {
			ret = append(ret, v)
		}
	}
	return ret
}
