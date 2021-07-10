package main

import (
	"fmt"
	"sort"
	"strconv"
)

type M map[string]map[string]int

var info M
var tables []int
var meal map[string]bool
var meals []string

func displayTable(orders [][]string) [][]string {
	info = make(M)
	meal = make(map[string]bool)
	for _, v := range orders {
		if _, ok := info[v[1]]; !ok {
			info[v[1]] = make(map[string]int)
		}
		if _, ok := info[v[1]][v[2]]; ok {
			info[v[1]][v[2]] += 1
		} else {
			info[v[1]][v[2]] = 1
		}
		meal[v[2]] = true
	}
	meals = make([]string, 0, len(meal))
	for k, _ := range meal {
		meals = append(meals, k)
	}
	sort.Strings(meals)
	tables = make([]int, 0, len(info))
	for k, _ := range info {
		x, _:=strconv.Atoi(k)
		tables = append(tables, x)
	}
	sort.Ints(tables)
	result := make([][]string, len(tables)+1)
	result[0] = make([]string, 0, len(meals)+1)
	result[0] = append([]string{"Table"}, meals...)
	for k, v := range tables {
		vi := strconv.Itoa(v)
		result[k+1] = make([]string, len(meals)+1)
		result[k+1][0] = vi
		for x, y := range meals {
			result[k+1][x+1] = strconv.Itoa(info[vi][y])
		}
	}
	return result
}

func main() {
	fmt.Println(
		displayTable(
			[][]string{
				{"Laura", "2", "Bean Burrito"},
				{"Jhon", "2", "Beef Burrito"},
				{"Melissa", "2", "Soda"}},
		),
	)
}
