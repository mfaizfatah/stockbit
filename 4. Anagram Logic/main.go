package main

import (
	"fmt"
	"sort"
	"strings"
)

var (
	words = []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
)

func checkAnagram() {
	var (
		w   []string
		key string
	)

	groupWord := make(map[string][]string)

	for _, word := range words {
		w = strings.Split(word, "")
		sort.Strings(w)
		key = strings.Join(w, "")
		groupWord[key] = append(groupWord[key], word)
	}

	for _, gword := range groupWord {
		fmt.Print(" [ ")
		for _, gw := range gword {
			fmt.Print(gw, " ")
		}
		fmt.Println("]")
	}
}

func main() {
	checkAnagram()
}
