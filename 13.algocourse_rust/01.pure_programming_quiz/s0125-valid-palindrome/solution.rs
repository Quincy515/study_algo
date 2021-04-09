impl Solution {
    pub fn is_palindrome(s: String) -> bool {
        let sm = s
            .chars()
            .filter(|c| c.is_ascii_alphanumeric())
            .map(|c| c.to_ascii_lowercase());

        let t = sm.clone().rev();

        sm.eq(t)
    }
}
