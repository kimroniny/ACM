计算每个元素和前面每个元素的差值，并使用前面元素对应该差值的答案数，更新该元素对应该差值的答案数，即在原基础上做加 1 操作即可。

统计答案的时候，要注意是没有加 1 的，因为一个等差数列的第二个元素会在第一个元素的答案基础上执行加1操作，但是两个元素是不能构成等差数列的，所以每个元素和之前元素的所有差值对应的答案都多加了个 1，故而统计答案的时候必须减去 1 ，所以不需要加 1 。

```golang
func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	dp := make([]map[int]int, n)
	ans := 0
	for i := 1; i < n; i++ {
		dp[i] = make(map[int]int)
		for j := 0; j < i; j++ {
			d := nums[i] - nums[j]
			dp[i][d] += dp[j][d] + 1
			ans += dp[j][d]
		}
	}
	return ans
}
```