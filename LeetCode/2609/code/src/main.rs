struct Solution{}

impl Solution {
    pub fn find_the_longest_balanced_substring(s: String) -> i32 {
        let mut temp = 0;
        let mut temp1 = 0;
        let mut flag = false;
        let mut ans = 0;
        for i in s.as_bytes() {
            if *i == b'0' {
                if temp1 != 0 {
                    ans = ans.max(temp1);
                    temp1 = 0;
                    temp = 0;
                }
                temp += 1;
            }else {
                if temp != 0 {
                    temp -= 1;
                    temp1 += 1;
                }else {
                    ans = ans.max(temp1);
                }
            }
        }
        ans = ans.max(temp1);
        ans*2
    }
}

fn main() {
    let s = String::from("001011");
    let ans = Solution::find_the_longest_balanced_substring(s);
    println!("ans={ans}");
}
