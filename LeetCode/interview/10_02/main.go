package main

import (
	"fmt"
	"sort"
)

type Runes []rune

func (a Runes) Len() int           { return len(a) }
func (a Runes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Runes) Less(i, j int) bool { return a[i] < a[j] }

// var maps map[string][]string

// func groupAnagrams(strs []string) [][]string {
// 	maps = make(map[string][]string)
// 	l := len(strs)
// 	for _, v := range strs {
// 		runes := Runes(v)
// 		sort.Sort(runes)
// 		runestr := string(runes)
// 		if _, ok := maps[runestr]; !ok {
// 			maps[runestr] = make([]string, 0, l)
// 		}
// 		maps[runestr] = append(maps[runestr], v)
// 	}
// 	result := make([][]string, 0, len(maps))
// 	for _, v := range maps {
// 		result = append(result, v)
// 	}
// 	return result
// }

// var maps map[string]int
// func groupAnagrams(strs []string) [][]string {
// 	l := len(strs)
// 	cnt := 0
// 	cnts := make([]string, l)
// 	maps = make(map[string]int)
// 	for k, v := range strs {
// 		runes := Runes(v)
// 		sort.Sort(runes)
// 		runestr := string(runes)
// 		cnts[k] = runestr
// 		if _,ok := maps[runestr]; !ok {
// 			maps[runestr] = cnt
// 			cnt += 1
// 		}
// 	}
// 	result := make([][]string, len(maps))
// 	for k,v := range strs {
// 		result[maps[cnts[k]]] = append(result[maps[cnts[k]]], v)
// 	}
// 	return result
// }

var maps map[string]int

func groupAnagrams(strs []string) [][]string {
	cnt := 0
	maps = make(map[string]int)
	result := make([][]string, 0, len(maps))
	for _, v := range strs {
		runes := Runes(v)
		sort.Sort(runes)
		runestr := string(runes)
		if _, ok := maps[runestr]; !ok {
			maps[runestr] = cnt
			result = append(result, make([]string, 0))
			result[cnt] = append(result[cnt], v)
			cnt += 1
		} else {
			result[maps[runestr]] = append(result[maps[runestr]], v)
		}

	}
	return result
}
func main() {
	str := "123456"

	fmt.Println(str)
}
