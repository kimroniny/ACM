struct Solution{}

// impl Solution2 {
//     pub fn vowel_strings(words: Vec<String>, left: i32, right: i32) -> i32 {
//         let mut ans = 0;
//         for i in left..right+1 {
//             let word = words.get(i as usize).unwrap();
//             ans += ("aeiou".chars().any(|c| word.starts_with(c)) && "aeiou".chars().any(|c| word.ends_with(c))) as i32
//         }
//         ans
//     }
// }

impl Solution {
    pub fn vowel_strings(words: Vec<String>, left: i32, right: i32) -> i32 {
        words[left as usize..(right+1) as usize]
        .iter()
        .filter(
            |word| 
            "aeiou".chars().any(|c| word.starts_with(c)) && 
            "aeiou".chars().any(|c| word.ends_with(c)))
        .count() 
        as i32
    }
}

fn main() {
    let words = vec![
        String::from("are"),
        String::from("amy"),
        String::from("u"),
    ];
    let left = 0;
    let right = 2;
    let solution = Solution::vowel_strings(words, left, right);
    println!("ans={solution}");
}
