impl Solution {
    pub fn is_palindrome(s: String) -> bool {
        if s.is_empty() {
            return true; // 边界情况，字符串为空或者长度为0，确定是回文字符串
        }
        let seq = s.chars().collect::<Vec<_>>();
        let (mut i, mut j) = (0_usize, seq.len() - 1);
        while i < j {
            while i < j && !seq[i].is_ascii_alphanumeric() {
                // 判断是否是字母或数字
                i += 1;
            }
            while i < j && !seq[j].is_ascii_alphanumeric() {
                j -= 1;
            }
            if i < j && seq[i].to_ascii_lowercase() != seq[j].to_ascii_lowercase() {
                return false; // 转成小写后是否相等
            }
            // 由于 j 是 usize，所以 j 有可能变为 0，然后 j-=1 后溢出
            if i == j {
                return true;
            }
            i += 1;
            j -= 1;
        }
        return true;
    }
}
