impl Solution {
    pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
        if nums.is_empty() {
            return 0;
        }
        let mut slow = 0; // slow 慢指针
        for i in 1..nums.len() {
            // i 是快指针
            if nums[i] != nums[slow] {
                // 不重复的值赋值给 nums[slow+1]
                slow += 1;
                nums[slow] = nums[i]
            } // 重复的值将跳过
        }
        k as i32 + 1 // 返回数组的长度
    }
}
