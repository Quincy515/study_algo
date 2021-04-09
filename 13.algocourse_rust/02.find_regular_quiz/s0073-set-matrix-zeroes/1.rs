impl Solution {
    pub fn set_zeroes(matrix: &mut Vec<Vec<i32>>) {
        let (mut m, mut n) = (matrix.len(), matrix[0].len());
        let mut row = Vec::with_capacity(m);
        row.resize(m, false);
        let mut col = Vec::with_capacity(n);
        col.resize(n, false);
        // let (mut row, mut col) = (vec![false; m], vec![false; n]);
        for i in 0..m {
            for j in 0..n {
                if matrix[i][j] == 0 {
                    row[i] = true;
                    col[j] = true;
                }
            }
        }
        for i in 0..m {
            for j in 0..n {
                if row[i] || col[j] {
                    matrix[i][j] = 0
                }
            }
        }
    }
}
