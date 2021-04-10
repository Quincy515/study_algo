pub struct Solution {}

// 因为只有 4 种颜色字母，所以用一个数组存储每个字母出现的次数
// 统计每个字母出现次数较小的次数，得到猜中和伪猜中的次数
// 再次遍历两个数组，比较相同位置相同颜色的次数，得到猜中的次数
// 用总次数减去猜中次数，就是伪猜中次数
impl Solution {
  pub fn master_mind(solution: String, guess: String) -> Vec<i32> {
    let s = solution.chars().collect::<Vec<char>>();
    let g: Vec<char> = guess.chars().collect();
    let mut a = vec![0; 4];
    let mut b = vec![0; 4];
    for i in 0..4 {
      match s[i] {
        'R' => a[0] += 1,
        'Y' => a[1] += 1,
        'G' => a[2] += 1,
        'B' => a[3] += 1,
        _ => (),
      }
      match g[i] {
        'R' => b[0] += 1,
        'Y' => b[1] += 1,
        'G' => b[2] += 1,
        'B' => b[3] += 1,
        _ => (),
      }
    }
    let mut total = 0;
    let mut ans = 0;
    for i in 0..4 {
      total += std::cmp::min(a[i], b[i]);
      ans += (s[i] == g[i]) as i32;
    }
    vec![ans, total - ans]
  }
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test_16_15() {
    assert_eq!(vec![1, 1], master_mind("RGBY".to_string(), "GGRR".to_string()));
  }
}