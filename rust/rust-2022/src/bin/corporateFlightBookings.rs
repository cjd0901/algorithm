// Corporate Flight Bookings
// leetcode: https://leetcode-cn.com/problems/corporate-flight-bookings/
struct Solution {}
impl Solution {
    pub fn corp_flight_bookings(bookings: Vec<Vec<i32>>, n: i32) -> Vec<i32> {
        let mut nums = vec![0; n as usize];
        for booking in bookings.iter() {
            nums[booking[0] as usize-1] += booking[2];
            if booking[1] < n {
                nums[booking[1] as usize] -= booking[2];
            }
        }
        for i in 1..n as usize {
            nums[i] += nums[i-1];
        }
        nums
    }
}

fn main() {
    let mut v = vec![];
    v.push(vec![1, 2, 10]);
    v.push(vec![2, 3, 20]);
    v.push(vec![2, 5, 25]);
    let ans = Solution::corp_flight_bookings(v, 5);
    println!("{:?}", ans);
}