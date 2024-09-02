struct NumArray {
    values: Vec<i32>
}


/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl NumArray {

    fn new(nums: Vec<i32>) -> Self {
        NumArray { values: nums }
    }
    
    fn update(&self, index: i32, val: i32) {
        let mut x = (1 << ((index+1).ilog2() + 1))-1;
        while x < self.values.len() as u32 {
            self.values[x as usize] += index;
            x = ((x+1)<<1)-1;
        }
    }
    
    fn sum_range(&self, left: i32, right: i32) -> i32 {
        let mut x = left;
        let mut ans = self.values[right as usize];
        while x >= 0 {
            ans -= self.values[x as usize];
            // x = (x+1).ilog2();
        }
        ans
    }
}

/**
 * Your NumArray object will be instantiated and called as such:
 * let obj = NumArray::new(nums);
 * obj.update(index, val);
 * let ret_2: i32 = obj.sum_range(left, right);
 */

/**
 * Your RangeModule object will be instantiated and called as such:
 * let obj = RangeModule::new();
 * obj.add_range(left, right);
 * let ret_2: bool = obj.query_range(left, right);
 * obj.remove_range(left, right);
 */
fn main() {
    println!("Hello, world!");
}
