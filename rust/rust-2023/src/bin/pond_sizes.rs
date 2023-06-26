struct Solution {}

impl Solution {
    pub fn pond_sizes(mut land: Vec<Vec<i32>>) -> Vec<i32> {
        fn dfs(land: &mut Vec<Vec<i32>>, x: usize, y: usize) -> i32 {
            if land[x][y] != 0 {
                return 0;
            }
            let mut part_size = 1;
            land[x][y] = -1;
            for path in [
                [0, 1],
                [0, -1],
                [1, 0],
                [-1, 0],
                [1, 1],
                [1, -1],
                [-1, 1],
                [-1, -1],
            ]
            .into_iter()
            {
                let (n_x, n_y) = (x as i32 + path[0], y as i32 + path[1]);
                if n_x < 0 || n_x >= land.len() as i32 || n_y < 0 || n_y >= land[0].len() as i32 {
                    continue;
                }
                part_size += dfs(land, n_x as usize, n_y as usize);
            }
            part_size
        }
        let mut ans = vec![];
        for x in 0..land.len() {
            for y in 0..land[0].len() {
                let size = dfs(&mut land, x, y);
                if size > 0 {
                    ans.push(size);
                }
            }
        }
        ans.sort();
        ans
    }
}

fn main() {}
