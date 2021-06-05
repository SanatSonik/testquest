package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"time"
)

type Word struct {
	wordBytes []byte
	count     int
}

func listWord(file []byte, listWordArr chan []byte) {
	var wordArr []byte
	for _, b := range file {
		if b >= 65 && b <= 90 {
			wordArr = append(wordArr, b+32)
		} else if b >= 97 && b <= 122 {
			wordArr = append(wordArr, b)
		} else if len(wordArr) > 0 {
			listWordArr <- wordArr
			wordArr = []byte{}
		}
	}
	close(listWordArr)
}

func uniqueChecker(resultArr []*Word, word []byte) (bool, int) {
	i := 0
	for _, wordStruct := range resultArr {
		if string(wordStruct.wordBytes) == string(word) {
			return true, i
		}
		i++

	}
	return false, 0
}

func listResult(resultArr []*Word, listWordArr chan []byte) {

	for word := range listWordArr {
		Checker, resultI := uniqueChecker(resultArr, word)

		if !Checker {
			resultArr = append(resultArr, &Word{word, 1})
		} else {
			if len(resultArr) > 0 {
				resultArr[resultI].count++
			}
		}
	}

}

func frequentWords(io.Writer) {
	start := time.Now()

	file, err := ioutil.ReadFile("mobydick.txt")
	if err != nil {
		panic(err)
	}

	var result []*Word
	list := make(chan []byte)

	go listWord(file, list)

	listResult(result, list)

	sort.Slice(result, func(i, j int) bool {
		return result[i].count > result[j].count
	})

	for i := 0; i < len(result) && i < 20; i++ {
		fmt.Println(string(result[i].wordBytes), ":", result[i].count)
	}

	fmt.Printf("Process took %s\n", time.Since(start))
}

func main() {
	start := time.Now()

	file, err := ioutil.ReadFile("mobydick.txt")
	if err != nil {
		panic(err)
	}

	var result []*Word
	list := make(chan []byte)

	go listWord(file, list)

	listResult(result, list)

	sort.Slice(result, func(i, j int) bool {
		return result[i].count > result[j].count
	})

	for i := 0; i < len(result) && i < 20; i++ {
		fmt.Println(string(result[i].wordBytes), ":", result[i].count)
	}

	fmt.Printf("Process took %s\n", time.Since(start))
}
