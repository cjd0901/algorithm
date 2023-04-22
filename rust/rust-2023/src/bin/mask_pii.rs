struct Solution {}

// 831. 隐藏个人信息
impl Solution {
    pub fn mask_pii(s: String) -> String {
        if let Some(index) = s.find("@") {
            let s = s.to_lowercase().chars().collect::<Vec<_>>();
            return format!("{}*****{}", s[0], s[index - 1..].iter().collect::<String>());
        }
        let s = s
            .chars()
            .into_iter()
            .filter(|x| x.is_digit(10))
            .collect::<Vec<_>>();
        format!(
            "{}***-***-{}",
            ["", "+*-", "+**-", "+***-"][s.len() - 10],
            s[s.len() - 4..].iter().collect::<String>()
        )
    }
}

fn main() {}
