// Peak Index in a Mountain Array
// leetcode: https://leetcode-cn.com/problems/peak-index-in-a-mountain-array/

struct Solution {

}

impl Solution {
    pub fn peak_index_in_a_mountain_array(arr: Vec<i32>) -> i32 {
        let (mut lo, mut hi) = (0, arr.len()-2);
        let mut ans: i32 = 0;
        while lo <= hi {
            let mid = lo + (hi-lo) / 2;
            if arr[mid] > arr[mid+1] {
                ans = mid as i32;
                hi = mid - 1;
            }else {
                lo = mid + 1;
            }
        }
        return ans
    }
}

// [3,4,5,1]
fn main() {
    let mut arr = Vec::<i32>::new();
    arr.push(3);
    arr.push(4);
    arr.push(5);
    arr.push(1);
    let ans = Solution::peak_index_in_a_mountain_array(arr);
    println!("ans{}", ans)
}