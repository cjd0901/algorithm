struct Solution {}

impl Solution {
    pub fn final_value_after_operations(operations: Vec<String>) -> i32 {
        let mut op_count = vec![0; 2];
        for op in operations {
            match op.as_str() {
                "X++" | "++X" => op_count[1] += 1,
                "X--" | "--X" => op_count[0] += 1,
                _ => ()
            }
        }
        op_count[1] - op_count[0]
    }
}

fn main() {}