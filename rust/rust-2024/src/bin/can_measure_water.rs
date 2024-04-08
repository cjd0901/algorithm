use std::collections::{HashSet, VecDeque};

struct Solution {}

impl Solution {
    pub fn can_measure_water(jug1_capacity: i32, jug2_capacity: i32, target_capacity: i32) -> bool {
        let mut stack: VecDeque<(i32, i32)> = VecDeque::new();
        let mut hit: HashSet<(i32, i32)> = HashSet::new();
        stack.push_back((0, 0));
        while !stack.is_empty() {
            let (jug1_remain, jug2_remain) = stack.pop_back().unwrap();
            if hit.contains(&(jug1_remain, jug2_remain)) {
                continue;
            }

            hit.insert((jug1_remain, jug2_remain));
            if jug1_remain == target_capacity
                || jug2_remain == target_capacity
                || jug1_remain + jug2_remain == target_capacity
            {
                return true;
            }

            stack.push_back((jug1_capacity, jug2_remain));
            stack.push_back((jug1_remain, jug2_capacity));
            stack.push_back((jug1_remain, 0));
            stack.push_back((0, jug2_remain));
            stack.push_back((
                jug1_remain - jug1_remain.min(jug2_capacity - jug2_remain),
                jug2_remain + jug1_remain.min(jug2_capacity - jug2_remain),
            ));
            stack.push_back((
                jug1_remain + jug2_remain.min(jug1_capacity - jug1_remain),
                jug2_remain - jug2_remain.min(jug1_capacity - jug1_remain),
            ));
        }
        false
    }
}

fn main() {}