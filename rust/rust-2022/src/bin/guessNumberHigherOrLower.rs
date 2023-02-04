// Guess Number Higher or Lower
// leetcode: https://leetcode-cn.com/problems/guess-number-higher-or-lower/

impl Solution {
    unsafe fn guess_number_higher(n: i32) -> i32 {
        let (mut lo, mut hi) = (1, n);
        while lo < hi {
            let mid = lo + (hi - lo) / 2;
            unsafe {
                let g = guess(mid);
                if g == 0 {
                    return mid
                }else if g == -1 {
                    hi = mid;
                }else {
                    lo = mid + 1;
                }
            }
        }
        return lo
    }
}