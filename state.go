package main

import (
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type State [][]int

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
