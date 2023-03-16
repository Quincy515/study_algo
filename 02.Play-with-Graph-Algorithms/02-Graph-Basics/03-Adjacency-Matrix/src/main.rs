use std::{fmt::Display, fs};

#[derive(Debug)]
struct AdjMatrix {
    v: i32,
    e: i32,
    adj: Vec<Vec<i32>>,
}

impl AdjMatrix {
    pub fn new(filename: &str) -> Self {
        let contents = fs::read_to_string(filename).unwrap();
        let mut lines = contents.lines();
        let first_line: Vec<i32> = lines
            .next()
            .unwrap()
            .split_whitespace()
            .map(|x| x.parse().unwrap())
            .collect();
        let v = first_line[0];
        let mut adj = vec![vec![0; v as usize]; v as usize];

        let e = first_line[1];
        for _ in 0..e as usize {
            let nums: Vec<usize> = lines
                .next()
                .unwrap()
                .split_whitespace()
                .map(|x| x.parse().unwrap())
                .collect();
            let (a, b) = (nums[0], nums[1]);
            adj[a][b] = 1;
            adj[b][a] = 1;
        }
        AdjMatrix { v, e, adj }
    }
}

impl Display for AdjMatrix {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "V = {}, E = {}\n", self.v, self.e)?;
        for i in 0..self.v {
            for j in 0..self.v {
                write!(f, "{} ", self.adj[i as usize][j as usize])?;
            }
            writeln!(f)?;
        }
        Ok(())
    }
}

fn main() {
    let adj = AdjMatrix::new("g.txt");
    println!("{adj}");
}
