package main

import (
    "fmt"
    "sort"
    "strings"
)

func groupAnagrams(words []string) [][]string {
    grouped := make(map[string][]string)
    var result [][]string

    for _, word := range words {
        chars := strings.Split(word, "")
        sort.Strings(chars)
        sortedWord := strings.Join(chars, "")

        grouped[sortedWord] = append(grouped[sortedWord], word)
    }

    for _, anagrams := range grouped {
        result = append(result, anagrams)
    }

    return result
}

func main() {
    inputArray := []string{"cook", "save", "taste", "aves", "vase", "state", "map"}
    groupedAnagrams := groupAnagrams(inputArray)

    for _, group := range groupedAnagrams {
        fmt.Println(group)
    }
}