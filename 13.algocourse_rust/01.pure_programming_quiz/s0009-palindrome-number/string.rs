impl Solution {
    pub fn is_palindrome(x: i32) -> bool {
        if x >= 0 && x < 10 {
            return true;
        }
        if x.to_string().chars().rev().collect::<String>() == x.to_string() {
            true
        } else {
            false
        }
    }
}
