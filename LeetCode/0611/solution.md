这种在有序数组中, 确定了一个数字, 找另外一个数字的题目, 一般都可以用二分法.

这个题目除了二分, 还有另外一种做法: 在确定了一个数字之后, 另外两个数字在朴素情况下, 是可以用过两个循环来遍历查找, 但是由于是有序数组, 而且三个数字之间有大小约束关系, 所以某些数字确定不可行之后, 其后的数字都可以确定为不可行, 就不需要遍历了, 利用这个特点, 就可以使用双指针来遍历.

```golang
func triangleNumber(nums []int) int {
	l := len(nums)
	ans := 0
	sort.Ints(nums)
	for i := 0; i < l; i++ {
		j, k := i+1, i+2
		for k < l {
			if nums[i]+nums[j] > nums[k] {
				ans += k - j
				k += 1
			} else {
				j += 1
				if j == k {
					k += 1
				}
			}
		}
	}
	return ans
}
```