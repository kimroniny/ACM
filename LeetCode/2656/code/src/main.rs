struct Solution {}

impl Solution {
    pub fn maximize_sum(nums: Vec<i32>, k: i32) -> i32 {
        (nums.iter().max().unwrap()*2+k-1)*k/2
    }
}

fn main() {
    let nums = vec![1,2,3,4,5];
    let k = 3;
    let ans = Solution::maximize_sum(nums, k);
    println!("ans={ans}");
}
