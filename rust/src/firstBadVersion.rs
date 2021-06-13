// First Bad Version
// leetcode: https://leetcode-cn.com/problems/first-bad-version/

impl Solution {
    pub fn first_bad_version(&self, n: i32) -> i32 {
        let (mut lo, mut hi) = (1, n);
        while lo < hi {
            let mid = lo + (hi - lo) / 2;
            if !self.isBadVersion(mid) {
                lo = mid+1;
            }else {
                hi = mid
            }
        }
        lo
    }
}
