package main

import (
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type State [9][9]int

func LoadStates(filename string) []*State {
	js, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	var states []*State
	err = json.Unmarshal(js, &states)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	return states
}

func (s *State) Fitness() int {
	collisions := 0
	// number of 0s
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s[i][j] == 0 {
				collisions++
			}
		}
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// ignore 0s
			if s[i][j] != 0 {
				// check cols
				for k := 0; k < 9; k++ {
					if s[i][j] == s[i][k] && j != k {
						collisions += 2000
					}
				}
				// check rows
				for k := 0; k < 9; k++ {
					if s[i][j] == s[k][j] && i != k {
						collisions += 2000
					}
				}
				row := (i / 3) * 3
				col := (j / 3) * 3
				// check squares
				for k := row; k < row+3; k++ {
					for l := col; l < col+3; l++ {
						if s[k][l] == s[i][j] && k != i && l != j {
							collisions += 1
						}
					}
				}
			}
		}
	}

	return collisions / 2
}