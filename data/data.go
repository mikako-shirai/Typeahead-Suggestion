package data

import (
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readDataFromFile() []string {
	data, err := ioutil.ReadFile("./words/english-words.txt")
	check(err)

	dataStr := string(data)
	dictWords := strings.Fields(dataStr)

	return dictWords
}
