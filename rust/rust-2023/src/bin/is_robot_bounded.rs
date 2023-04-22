struct Solution {}

// 1041. 困于环中的机器人
impl Solution {
    pub fn is_robot_bounded(instructions: String) -> bool {
        let steps = [[0, 1], [1, 0], [0, -1], [-1, 0]];
        let mut direction = 0;
        let (mut x, mut y) = (0, 0);
        for c in instructions.chars() {
            match c {
                'G' => {
                    x += steps[direction][0];
                    y += steps[direction][1];
                }
                'L' => {
                    direction += 3;
                    direction %= 4;
                }
                _ => {
                    direction += 1;
                    direction %= 4;
                }
            }
        }
        direction != 0 || (x == 0 && y == 0)
    }
}

fn main() {}
