package main

import (
	"fmt"
	"math"
	"sort"
)

type Info struct {
	idx, x int
}

type InfoList []*Info

func (a InfoList) Len() int           { return len(a) }
func (a InfoList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a InfoList) Less(i, j int) bool { return a[i].x > a[j].x }

const MOD int = 1000000007

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	minx, sum := 0, 0
	infos1 := make(InfoList, 0, len(nums1))
	infos2 := make(InfoList, 0, len(nums2))
	for k, v := range nums1 {
		x := int(math.Abs(float64(v - nums2[k])))
		sum = (sum + x) % MOD
		infos1 = append(infos1, &Info{idx: k, x: v})
		infos2 = append(infos2, &Info{idx: k, x: nums2[k]})
	}
	sort.Sort(infos1)
	sort.Sort(infos2)
	p := 0
	for i := 0; i < len(infos2); i++ {
		orix := int(math.Abs(float64(nums1[infos2[i].idx] - nums2[infos2[i].idx])))
		for p < len(infos1) {
			x := int(math.Abs(float64(infos1[p].x-infos2[i].x))) - orix
			if p > 0 {
				lastx := int(math.Abs(float64(infos1[p-1].x-infos2[i].x))) - orix
				if lastx < x {
					x = lastx
				}
			}
			if minx > x {
				minx = x
			}
			if infos2[i].x >= infos1[p].x {
				break
			} else {
				p += 1
			}
		}
		// 最后的边界
		if (p == len(infos1)) {
			x := int(math.Abs(float64(infos1[p-1].x-infos2[i].x))) - orix
			if minx > x {
				minx = x
			}
		}
	}
	x := ((sum + minx) % MOD + MOD) % MOD
	return x
}

func main() {
	fmt.Println(
		minAbsoluteSumDiff(
			[]int{38, 48, 73, 55, 25, 47, 45, 62, 15, 34, 51, 20, 76, 78, 38, 91, 69, 69, 73, 38, 74, 75, 86, 63, 73, 12, 100, 59, 29, 28, 94, 43, 100, 2, 53, 31, 73, 82, 70, 94, 2, 38, 50, 67, 8, 40, 88, 87, 62, 90, 86, 33, 86, 26, 84, 52, 63, 80, 56, 56, 56, 47, 12, 50, 12, 59, 52, 7, 40, 16, 53, 61, 76, 22, 87, 75, 14, 63, 96, 56, 65, 16, 70, 83, 51, 44, 13, 14, 80, 28, 82, 2, 5, 57, 77, 64, 58, 85, 33, 24},
			[]int{90, 62, 8, 56, 33, 22, 9, 58, 29, 88, 10, 66, 48, 79, 44, 50, 71, 2, 3, 100, 88, 16, 24, 28, 50, 41, 65, 59, 83, 79, 80, 91, 1, 62, 13, 37, 86, 53, 43, 49, 17, 82, 27, 17, 10, 89, 40, 82, 41, 2, 48, 98, 16, 43, 62, 33, 72, 35, 10, 24, 80, 29, 49, 5, 14, 38, 30, 48, 93, 86, 62, 23, 17, 39, 40, 96, 10, 75, 6, 38, 1, 5, 54, 91, 29, 36, 62, 73, 51, 92, 89, 88, 74, 91, 87, 34, 49, 56, 33, 67},
		),
	)
}
