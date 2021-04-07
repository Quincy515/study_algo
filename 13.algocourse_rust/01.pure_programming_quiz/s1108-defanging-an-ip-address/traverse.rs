pub struct Solution {}

impl Solution {
    pub fn defang_i_paddr(address: String) -> String {
        let mut result = String::new();
        for i in address.chars() {
            if i == '.' {
                result.push_str("[.]");
            } else {
                result.push(i);
            }
        }
        result
    }
}

#[cfg(text)]
mod tests {
    use ::spuer::*;
    #[test]
    fn test_1() {
        assert_eq("1[.]1[.]1[.]1", Solution::defang_i_paddr("1.1.1.1"));
        assert_eq(
            "255[.]100[.]50[.]0",
            Solution::defang_i_paddr("255.100.50.0"),
        );
    }
}
