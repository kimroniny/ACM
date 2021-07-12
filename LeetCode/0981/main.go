package main

import "fmt"

const MAXELE int = 120001

type StampValue struct {
	tsp   int
	value string
}

type TimeMap struct {
	info map[string][]*StampValue
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{info: make(map[string][]*StampValue)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if len(this.info[key]) == 0 {
		this.info[key] = make([]*StampValue, 0, MAXELE)
	}
	this.info[key] = append(this.info[key], &StampValue{value: value, tsp: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if !this.ExistKey(key) {
		return ""
	}
	l, r := 0, len(this.info[key])-1
	for l <= r {
		mid := (l + r) >> 1
		midTsp := this.info[key][mid].tsp
		if midTsp < timestamp {
			l = mid + 1
		} else if midTsp == timestamp {
			return this.info[key][mid].value
		} else {
			r = mid - 1
		}
	}
	if r == -1 {
		return ""
	}
	return this.info[key][r].value
}

func (this *TimeMap) ExistKey(key string) bool {
	_, ok := this.info[key]
	return ok
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

func main() {
	s := Constructor()
	// cmds := []string{"TimeMap", "set", "get", "get", "set", "get", "get"}
	// infos := [][]interface{}{{}, {"foo", "bar", 1}, {"foo", "", 1}, {"foo", "", 3}, {"foo", "bar2", 4}, {"foo", "", 4}, {"foo", "", 5}}
	cmds := []string{"TimeMap", "set", "set", "get", "get", "get", "get", "get"}
	infos := [][]interface{}{{}, {"love", "high", 10}, {"love", "low", 20}, {"love", "", 5}, {"love", "", 10}, {"love", "", 15}, {"love", "", 20}, {"love", "", 25}}
	for k, cmd := range cmds {
		info := infos[k]
		if cmd == "set" {
			s.Set(info[0].(string), info[1].(string), info[2].(int))
			fmt.Println("null")
		} else if cmd == "get" {
			fmt.Println(s.Get(info[0].(string), info[2].(int)))
		} else {
			fmt.Println("null")
		}
	}
}
