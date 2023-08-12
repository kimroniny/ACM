struct Solution {}
impl Solution {
    pub fn diagonal_sum(mat: Vec<Vec<i32>>) -> i32 {
        let n = mat.len() as i32;
        (0..n*n)
        .into_iter()
        .filter(|x| (x/n == x%n) || (x/n == n-1-x%n))
        .map(|x| mat[(x/n) as usize][(x%n) as usize])
        .sum()
    }
}

fn main() {
    let mat = vec![
        vec![1,2,3],
        vec![4,5,6],
        vec![7,8,9],
    ];
    let ans = Solution::diagonal_sum(mat);
    println!("{ans}");
}