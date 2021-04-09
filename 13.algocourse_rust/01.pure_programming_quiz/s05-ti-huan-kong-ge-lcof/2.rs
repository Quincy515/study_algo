impl Solution {
    pub fn replace_space(s: String) -> String {
        let mut result = String::new();
        for c in s.chars() {
            match c {
                ' ' => result.push_str("%20"),
                _ => result.push(c),
            }
        }
        result
    }
}
