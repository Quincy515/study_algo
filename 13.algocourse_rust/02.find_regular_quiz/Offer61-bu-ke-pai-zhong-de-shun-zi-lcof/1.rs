impl Solution {
    pub fn is_straight(nums: Vec<i32>) -> bool {
        // 1. 定义一个临时数组
        let mut tmp = Vec::with_capacity(nums.len());
        // 2. 统计 0 的个数
        let mut cnt = 0;
        // 3. 填充这个临时数组
        for n in nums {
            if n == 0 {
                cnt += 1;
                continue;
            }
            tmp.push(n);
        }
        // 4. 判断临时数组是否是连续的
        tmp.sort();
        for i in 1..tmp.len() {
            if tmp[i] - tmp[i - 1] > 1 && cnt - (tmp[i] - tmp[i - 1]) + 1 >= 0 {
                continue;
            } else if tmp[i] - tmp[i - 1] == 1 {
                continue;
            } else {
                return false;
            }
        }
        true
    }
}
