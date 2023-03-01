struct Solution {}

//1144. 递减元素使数组呈锯齿状
impl Solution {
    pub fn moves_to_make_zigzag(nums: Vec<i32>) -> i32 {
        let calc = |start: usize| -> i32 {
            let mut sum = 0;
            for i in (start..nums.len()).step_by(2) {
                let mut ops = 0;
                if i > 0 && nums[i] >= nums[i - 1] {
                    ops = ops.max(nums[i] - nums[i - 1] + 1);
                }
                if i < nums.len() - 1 && nums[i] >= nums[i + 1] {
                    ops = ops.max(nums[i] - nums[i + 1] + 1);
                }
                sum += ops;
            }
            sum
        };
        calc(0).min(calc(1))
    }
}

fn main() {}
