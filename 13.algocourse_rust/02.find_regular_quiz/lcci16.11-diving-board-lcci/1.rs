impl Solution {
    pub fn diving_board(shorter: i32, longer: i32, k: i32) -> Vec<i32> {
        let mut res = Vec::new();
        if k == 0 {
            return res;
        }
        if longer == shorter {
            res.push(shorter * k);
            return res;
        }
        for i in 0..=k {
            let temp = shorter * (k - i) + longer * i;
            res.push(temp);
        }
        return res;
    }
}
