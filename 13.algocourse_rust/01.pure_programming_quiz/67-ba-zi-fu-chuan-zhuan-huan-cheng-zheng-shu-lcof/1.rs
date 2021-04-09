impl Solution {
    pub fn str_to_int(str: String) -> i32 {
        let mut s = str.trim();
        if s.is_empty() {
            return 0;
        }
        let mut negative = false;
        let mut res = 0_i64;
        for (i, ch) in str.chars().enumerate() {
            if ch == '+' && i == 0 {
                continue;
            }
            if ch == '-' && i == 0 {
                negative = true;
                continue;
            }
            if !ch.is_digit(10) {
                break;
            }
            res = 10i64 * res + ch.to_digit(10).unwrap() as i64;
            if negative {
                if -res < i32::min_value() as i64 {
                    return i32::min_value();
                }
            } else {
                if res > i32::max_value() as i64 {
                    return i32::max_value();
                }
            }
        }
        if negative {
            -res as i32
        } else {
            res as i32
        }
    }
}
