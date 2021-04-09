impl Solution {
    pub fn length_of_last_word(s: String) -> i32 {
        let seq: Vec<char> = s.chars().rev().collect();
        let (mut result, mut find) = (0i32, false);
        for ch in seq {
            if ch == ' ' && find {
                break;
            }
            if ch != ' ' {
                find = true;
                result += 1;
            }
        }
        result
    }
}
