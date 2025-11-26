struct Solution {}

impl Solution {
    pub fn max_sum_div_three(nums: Vec<i32>) -> i32 {
        let mut total_sum = 0;
        let mut v = vec![vec![]; 3];
        for num in nums {
            v[(num % 3) as usize].push(num);
            total_sum += num;
        }

        let l1 = v[1].len();
        let l2 = v[2].len();
        v[1].sort_by(|a, b| b.cmp(a));
        v[2].sort_by(|a, b| b.cmp(a));

        let mut remove = i32::MAX;
        if total_sum % 3 == 0 {
            remove = 0;
        } else if total_sum % 3 == 1 {
            if l1 >= 1 {
                remove = v[1][l1-1];
            }
            if l2 >= 2 {
                remove = remove.min(v[2][l2-1] + v[2][l2-2]);
            }
        } else if total_sum % 3 == 2 {
            if l2 >= 1 {
                remove = v[2][l2-1];
            }
            if l1 >= 2 {
                remove = remove.min(v[1][l1-1] + v[1][l1-2]);
            }
        }
        total_sum - remove
    }

    pub fn max_sum_div_three_dp(nums: Vec<i32>) -> i32 {
        let mut dp = vec![vec![0; 3]; nums.len()+1];
        dp[0] = vec![0, i32::MIN, i32::MIN];
        for i in 0..nums.len() {
            let num = nums[i];
            for j in 0..3usize {
                dp[i+1][(j+num as usize)%3] = dp[i][(j+num as usize)%3].max(dp[i][j] + num)
            }
        }
        dp[nums.len()][0]
    }
}

fn main() {
    let ans = Solution::max_sum_div_three_dp(vec![3, 6, 5, 1, 8]);
    println!("{ans}");
}