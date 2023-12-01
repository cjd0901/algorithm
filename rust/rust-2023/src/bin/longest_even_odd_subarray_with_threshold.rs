struct Solution {}

impl Solution {
    pub fn longest_alternating_subarray(nums: Vec<i32>, threshold: i32) -> i32 {
        let len = nums.len();
        let (mut start, mut end): (usize, usize) = (0, 0);
        let mut ans = 0;
        while end < len && start < len {
            if nums[start] % 2 > 0 || nums[start] > threshold {
                start += 1;
                continue;
            }

            end = start;
            while end < len {
                if nums[end] > threshold {
                    start = end + 1;
                    end = start;
                    break;
                }

                if end > start && end < len && nums[end] % 2 == nums[end - 1] % 2 {
                    start = end;
                    break;
                }

                ans = ans.max(end as i32 - start as i32 + 1);
                end += 1;
            }
        }
        ans
    }
}

fn main() {
    let ans = Solution::longest_alternating_subarray(vec![2, 2, 5, 1, 6, 7, 8], 17);
    println!("{ans}");
}
