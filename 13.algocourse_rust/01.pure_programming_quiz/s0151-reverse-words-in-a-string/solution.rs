use std::iter::FromIterator;
pub struct Solution {}

impl Solution {
    pub fn reverse_words(s: String) -> String {
        if s == "" || s.len() == 0 {
            return s; // 字符串为空，或长度为0，就返回它自己
        }
        let mut char_vec: Vec<char> = s.chars().collect(); // 把字符串转为字符数组
        let mut p = 0; // 定义游标 p
        let mut q = 0; // 定义游标 q
        let mut end = s.len() - 1; // 定义游标 end 初始化指向字符串末尾
        while end >= 0 && char_vec[end] == ' ' {
            end -= 1; // 当 end 为空格时不断向前移动
        }
        while q <= end {
            let start = p; // 记录单词开始下标
            while q <= end && char_vec[q] == ' ' {
                q += 1; // 游标 q 跳过空格
            }
            while q <= end && char_vec[q] != ' ' {
                char_vec[p] = char_vec[q]; // q 不指向空格，就把字符拷贝到游标 p 的位置
                p += 1; // 同时把 p 向后移动一位
                q += 1; // 同时把 q 向后移动一位
            }
            // 拷贝完一个单词后去翻转这个单词
            char_vec[start..p].reverse();
            if q <= end {
                // 当 q 还没有移动到 end 时，说明还有单词要处理
                char_vec[p] = ' '; // 在 p 的位置插入一个空格
                p += 1; // 并向后移动一位
            }
        }
        // 循环结束后，把字符串中有效部分翻转
        char_vec[..p].reverse();
        char_vec.truncate(p);
        String::from_iter(char_vec)
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_151() {
        assert_eq!(
            "blue is sky the".to_owned(),
            Solution::reverse_words("the sky is blue".to_owned())
        );
    }
    #[test]
    fn reverse_words_test2() {
        let s = "  hello world  ".to_string();
        assert_eq!(Solution::reverse_words(s), "world hello".to_string());
    }

    #[test]
    fn reverse_words_test3() {
        let s = "  Bob    Loves  Alice   ".to_string();
        assert_eq!(Solution::reverse_words(s), "Alice Loves Bob".to_string());
    }

    #[test]
    fn reverse_words_test4() {
        let s = "Alice does not even like bob".to_string();
        assert_eq!(
            Solution::reverse_words(s),
            "bob like even not does Alice".to_string()
        );
    }
}
