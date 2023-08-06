struct Solution {

}
impl Solution {
    pub fn reverse_string(s: &mut Vec<char>) {
        let length = s.len();
        for i in 0..length/2 {
            let t = s[i];
            s[i] = s[length-i-1];
            s[length-i-1] = t;
        }
    }
}

fn main() {
    let s = &mut vec!['1', '2', '3'];
    Solution::reverse_string(s);
    println!("{:?}", s);
}