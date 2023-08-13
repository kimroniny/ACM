struct Solution{}
impl Solution {
    pub fn merge(nums1: &mut Vec<i32>, m: i32, nums2: &mut Vec<i32>, n: i32) {
        let mut i = m-1; 
        let mut j = n-1;
        let mut p = m+n-1;
        while i >= 0 || j >=0 {
            if i < 0 {
                nums1[p as usize] = nums2[j as usize];
                j -= 1; p -= 1;
            } else 
            if j < 0 {
                nums1[p as usize] = nums1[i as usize];
                i -= 1; p -= 1;
            } else {
                if nums1[i as usize] < nums2[j as usize] {
                    nums1[p as usize] = nums2[j as usize];
                    j-=1; p-=1;
                }else {
                    nums1[p as usize] = nums1[i as usize];
                    i -= 1; p-= 1;
                }
            }
            
        }
        
    }
}
fn main() {
    
    let mut nums1 = vec![4,5,6,0,0,0];
    let l1 = nums1.len() as usize;
    let mut nums2 = vec![1,2,3];
    let l2 = nums2.len() as usize;
    Solution::merge(&mut nums1, (l1 -l2) as i32, &mut nums2, l2 as i32);
    println!("{:#?}", nums1);

}