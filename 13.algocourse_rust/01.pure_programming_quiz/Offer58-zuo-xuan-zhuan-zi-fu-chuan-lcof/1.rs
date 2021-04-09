impl Solution {
    pub fn reverse_left_words(s: String, n: i32) -> String {
        String::from(&s[n as usize..]) + &s[..n as usize]
    }
}
