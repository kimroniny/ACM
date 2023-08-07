struct Solution {}
impl Solution {
    pub fn min_falling_path_sum(matrix: Vec<Vec<i32>>) -> i32 {
        use std::cmp::min;
        let m = matrix[0].len();
        let mut dp = vec![0;m];
        let mut dp0 = vec![0;m];
        for i in 0..m {
            dp0[i] = matrix[0][i];
        }
        for i in 1..matrix.len() {
            for j in 0..matrix[i].len() {
                let x = matrix[i][j];
                dp[j] = dp0[j];
                if j > 0 {
                    dp[j] = min(dp[j], dp0[j-1]);
                }
                if j < matrix[i].len()-1 {
                    dp[j] = min(dp[j], dp0[j+1])
                }
                dp[j] += x;
            }
            for j in 0..matrix[i].len() {
                dp0[j] = dp[j]; dp[j] = 0;
            }
        }
        *dp0.iter().min().unwrap()
    }
}

fn main() {
    let matrix = vec![
        vec![2,1,3],
        vec![6,5,4],
        vec![7,8,9],
    ];
    let ans = Solution::min_falling_path_sum(matrix);
    println!("{}", ans);
}