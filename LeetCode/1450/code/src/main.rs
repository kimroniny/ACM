use std::ops::Index;

impl Solution {
    pub fn busy_student(start_time: Vec<i32>, end_time: Vec<i32>, query_time: i32) -> i32 {
        start_time.iter().zip(end_time.iter()).map(|(&a, &b)| (query_time >= a && query_time <= b) as i32).sum()
    }
}