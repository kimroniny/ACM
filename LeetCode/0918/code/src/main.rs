use core::num;

struct Solution {}
impl Solution {
    pub fn max_subarray_sum_circular(nums: Vec<i32>) -> i32 {
        use std::cmp::max;
        let n = nums.len() as i32;
        // let nums:Vec<i32> = (0..n*2-1).map(|i| nums[(i%n) as usize]).collect();
        let mut tail: i32 = -1;
        let mut ans = 0;
        let mut sum = 0;
        for i in 0..(n*2-1) {
            let idx = i%n;
            if idx > tail && i >= n {continue;}
            sum += nums[idx as usize];
            if sum <= 0 { tail = idx as i32; sum = 0;}
            else  {
                ans = max(ans, sum);
            }
        }
        ans
    }
}
fn main() {
    let nums = vec![-10,1,-13,4];
    let ans = Solution::max_subarray_sum_circular(nums);
    println!("{ans}");
}
