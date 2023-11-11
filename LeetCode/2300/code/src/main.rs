struct Solution{}
impl Solution {
    pub fn successful_pairs2(mut spells: Vec<i32>, mut potions: Vec<i32>, success: i64) -> Vec<i32> {
        potions.sort_unstable();
        let mut ans = vec![0;spells.len()];
        for (idx, x) in spells.iter().enumerate() {
            let mut l = 0 as usize;
            let mut r = potions.len();
            while l != r {
                let mid = (l+r)>>1;
                let v = potions[mid];
                let temp = v as i64 * (*x as i64);
                // let temp = (v * x) as i64;
                if temp < success {
                    l = mid + 1;
                }else if temp >= success {
                    r = mid;
                }
            }
            ans[idx]=(potions.len()-l) as i32;
        }
        ans.to_owned()
    }
    pub fn successful_pairs(mut spells: Vec<i32>, mut potions: Vec<i32>, success: i64) -> Vec<i32> {
        potions.sort_unstable();
        let mut ans = vec![0;spells.len()];
        spells.iter().enumerate().for_each(|(idx, x)| {
            let mut l = 0 as usize;
            let mut r = potions.len();
            while l != r {
                let mid = (l+r)>>1;
                let v = potions[mid];
                let temp = v as i64 * (*x as i64);
                // let temp = (v * x) as i64;
                if temp < success {
                    l = mid + 1;
                }else if temp >= success {
                    r = mid;
                }
            }
            ans[idx]=(potions.len()-l) as i32;
        });
        ans
    }
}
fn main() {
    let spells = vec![100000];
    let potions = vec![100000];
    let success = 10000000000;
    let ans = Solution::successful_pairs(spells, potions, success);
    println!("ans={:?}", ans);
}
