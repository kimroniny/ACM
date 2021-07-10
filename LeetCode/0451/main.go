package main

import (
	"bytes"
	"fmt"
	"sort"
)

var maps map[rune]int

type Info struct {
	c     rune
	count int
}

type Infos []Info

func (a Infos) Len() int           { return len(a) }
func (a Infos) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Infos) Less(i, j int) bool { return a[i].count > a[j].count }

var infos Infos

func frequencySort(s string) string {
	maps = make(map[rune]int)
	for _, v := range s {
		if _, ok := maps[v]; !ok {
			maps[v] = 1
		} else {
			maps[v] += 1
		}
	}
	infos = make([]Info, 0, len(maps))
	for k, v := range maps {
		infos = append(infos, Info{c: k, count: v})
	}
	sort.Sort(infos)
	var buffer bytes.Buffer
	for _, v := range infos {
		for i := 0; i < v.count; i++ {
			buffer.WriteRune(v.c)
		}
	}
	return buffer.String()
}

func main() {
	fmt.Println(frequencySort(
		"aavbvv",
	))
}
