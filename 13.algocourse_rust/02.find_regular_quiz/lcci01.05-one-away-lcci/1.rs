impl Solution {
    pub fn one_edit_away(first: String, second: String) -> bool {
        let word1: Vec<char> = first.chars().collect();
        let word2 = second.chars().collect::<Vec<char>>();
        let (l1, l2) = (word1.len(), word2.len());
        // 因为只有一次编辑的机会，如果两个字符串长度相差大于1，直接返回 false
        if (l1 as i64 - l2 as i64).abs() > 1 {
            return false;
        }
        let (mut i, mut j) = (0, 0);
        while i < l1 && j < l2 {
            // 匹配成功就继续
            if word1[i] == word2[j] {
                i += 1;
                j += 1;
                continue;
            }
            // 匹配失败会有3中情况
            // 如果这两个字符串长度相等，从他们的下一个开始比较
            // 如果 first 比 second 长，就把 first 截取一位和 second 比较
            // 如果 second 比 first 长，就把 scecond 截取一位和 first 比较
            if l1 == l2 {
                return &word1[i + 1..] == &word2[i + 1..];
            }
            if l1 > l2 {
                return &word1[i + 1..] == &word2[i..];
            }
            return &word1[i..] == &word2[i + 1..];
        }
        true
    }
}
