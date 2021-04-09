impl Solution {
    pub fn is_straight(nums: Vec<i32>) -> bool {
        let mut sets = std::collections::HashSet::new();
        for n in nums {
            if n != 0 && !sets.insert(n) {
                return false;
            }
            let min = *sets.iter().min().unwrap();
            let max = sets.into_iter().max().unwrap();
            max - min < 5
        }
    }
}
