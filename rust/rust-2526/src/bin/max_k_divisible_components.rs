struct Solution {}

impl Solution {
    pub fn max_k_divisible_components(n: i32, edges: Vec<Vec<i32>>, values: Vec<i32>, k: i32) -> i32 {
        let mut del_edge_num = 0;
        let mut graph: Vec<Vec<i32>> = vec![vec![]; n as usize];
        for edge in edges {
            graph[edge[0] as usize].push(edge[1]);
            graph[edge[1] as usize].push(edge[0]);
        }
        fn dfs(graph: &Vec<Vec<i32>>, values: &Vec<i32>, k: i32, node: i32, parent: i32, del_edge_num: &mut i32) -> i64 {
            let mut sum = values[node as usize] as i64;
            for &edge in &graph[node as usize] {
                if edge != parent {
                    sum += dfs(graph, values, k, edge, node, del_edge_num);
                }
            }

            if sum % (k as i64) == 0 && parent != -1 {
                *del_edge_num += 1;
            }
            sum
        }
        dfs(&graph, &values, k, 0, -1, &mut del_edge_num);
        del_edge_num + 1
    }
}

fn main() {

}