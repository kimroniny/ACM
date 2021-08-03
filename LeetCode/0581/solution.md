两种方法:

### 排序
时间复杂度 `O(NlogN)` .
这种方法很容易理解, 比对排序前后的数组, 找到元素不相同的首尾位置即可 .

### 定界
时间复杂度 `O(N)` .
分成三个区间, 左区间是递增的, 右区间是递增的, 中间区间并不一定是目标区间, 目标区间要比中间区间更大, 因为左区间递增的元素有可能大于等于中间区间的元素, 同理, 右区间递增的元素有可能小于等于中间区间的元素. 因此需要中间区间的最小最大元素需要各自向左右区间拓展, 找到他们的边界.

但是在寻找边界的过程中, 有可能会遇到比中间区间最小元素更小的, 比中间区间最大元素更大的, 所以需要更新最大最小值. 

```golang
// 自己第一次想的
func findUnsortedSubarray(nums []int) int {
	flagl, flagr := -1,-1
	minx, maxx := math.MaxInt32, math.MinInt32
	for i := 1; i < len(nums); i++ {
		if flagl == -1 && nums[i] < nums[i-1] {
			flagl = i
		}
		if flagl != -1 {
			minx = min(minx, nums[i])
		}
	}
	if flagl == -1 {return 0}
	for i := len(nums)-2; i >= 0; i-- {
		if flagr == -1 && nums[i] > nums[i+1] {
			flagr = i
		}
		if flagr != -1 {
			maxx = max(maxx, nums[i])
		}
	}
	flagl, flagr = flagl-1, flagr+1
	for flagl >= 0 && minx < nums[flagl] {
		flagl -= 1
	}
	for flagr < len(nums) && maxx > nums[flagr] {
		flagr += 1
	}
	// fmt.Println(flagl, flagr, flagr-flagr-1)
	return flagr-flagl-1
}
```
另外一种定界的思想, 因为最小值只会在中间区间以及右区间, 然后用该最小值向左区间扫描边界, 最大值只会在中间区间以及左区间, 然后用该最大值向右区间扫描边界, 因此可以直接使用一个循环遍历.
```golang
// leetcode 官方题解
func findUnsortedSubarray(nums []int) int {
    n := len(nums)
    minn, maxn := math.MaxInt64, math.MinInt64
    left, right := -1, -1
    for i, num := range nums {
        if maxn > num {
            right = i
        } else {
            maxn = num
        }
        if minn < nums[n-i-1] {
            left = n - i - 1
        } else {
            minn = nums[n-i-1]
        }
    }
    if right == -1 {
        return 0
    }
    return right - left + 1
}
```