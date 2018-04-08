package main

import (
	"io/ioutil"
	"encoding/json"
)

// Board represents a sudoku board
type Board [9][9]int

// Sudoku represents a sudoku problem
type Sudoku struct {
	Board Board
}

// LoadBoard loads the board state from a file
func LoadBoard(filename string) *Sudoku {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	s := new(Sudoku)
	json.Unmarshal(data, &s.Board)
	return s
}

// WriteBoard saves the state of the sudoku board to file
func (s *Sudoku) WriteBoard(filename string) {
	data, err := json.Marshal(s.Board)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(filename, data, 0777)
}

// Complete checks to see if all values are non-zero
func (s *Sudoku) Complete() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s.Board[i][j] == 0 {
				return false
			}
		}
	}
	return true
}
