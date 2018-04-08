package main

import (
	"io/ioutil"
	"path"
	"time"
	"fmt"
)

func main() {
	puzzles, _ := ioutil.ReadDir("puzzles")
	start := time.Now()
	for _, puzzle := range puzzles {
		p := LoadBoard(path.Join("puzzles", puzzle.Name()))
		solver := NewSolver(p)
		solver.Solve()
		p.WriteBoard(path.Join("solutions", puzzle.Name()))

	}
	end := time.Now()
	fmt.Printf("Solved in %.08fs.\n", end.Sub(start).Seconds())
}
