struct Solution {}
impl Solution {
    pub fn max_absolute_sum(nums: Vec<i32>) -> i32 {
        let mut minx = 0;
        let mut maxx = 0;
        let mut sum0 = 0;
        let mut sum1 = 0;
        for x in nums {
            sum0 += x;
            maxx = maxx.max(sum0);
            sum0 = sum0.max(0);

            sum1 += x;
            minx = minx.min(sum1);
            sum1 = sum1.min(0);

        }
        
        maxx.abs().max(minx.abs())
    }
}

fn main() {
    let nums = vec![2,-5,1,-4,3,-2];
    let ans = Solution::max_absolute_sum(nums);
    println!("{ans}");
}
