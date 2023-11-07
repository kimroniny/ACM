use std::collections::HashMap;
use std::cmp::max;

struct Solution{}

impl Solution {
    pub fn max_product(words: Vec<String>) -> i32 {
        let mut records = HashMap::new();
        for word in words.iter() {
            // let mut x = 0;
            // for w in word.as_bytes() {
            //     x |= 1<<(w-('a' as u8));
            // }

            let x = word.bytes().fold(0, |acc, c| acc | (1<<(c-b'a')));

            records
                .entry(x)
                .and_modify(|c| { *c = word.len().max(*c)})
                .or_insert(word.len());

            // if let Some(a) = records.get(&x) {
            //     if *a < word.len() {
            //         records.insert(x, word.len());    
            //     }
            // }else {
            //     records.insert(x, word.len());
            // }
        }
        let mut ans = 0;
        for key in records.keys() {
            for key2  in records.keys() {
                if key & key2 == 0 {
                    ans = max(ans, records.get(key).unwrap() *  records.get(key2).unwrap());
                }
            }
        }
        ans as i32
    }
}

fn main() {
    let words = Vec::from(["abcw","baz","foo","bar","xtfn","abcdef"].map(|x| x.to_string()));
    let ans = Solution::max_product(words);
    println!("ans={ans}");
}
