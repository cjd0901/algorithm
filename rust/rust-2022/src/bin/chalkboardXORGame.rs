impl Solution {
    pub fn xor_game(nums: Vec<i32>) -> bool {
        let n = nums.len();
        let mut sum = 0;
        for i in &nums {
            sum ^= i;
        }
        n %2 == 0 || sum == 0
    }
}