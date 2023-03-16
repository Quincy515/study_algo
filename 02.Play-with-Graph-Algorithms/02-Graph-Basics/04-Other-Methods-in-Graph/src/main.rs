use std::{fmt::Display, fs};

#[derive(Debug, Default)]
struct AdjMatrix {
    v: usize,
    e: usize,
    adj: Vec<Vec<i32>>,
}

impl AdjMatrix {
    pub fn new(filename: &str) -> Self {
        let contents = fs::read_to_string(filename).unwrap();
        let mut lines = contents.lines();
        let first_line: Vec<usize> = lines
            .next()
            .unwrap()
            .split_whitespace()
            .map(|x| x.parse().unwrap())
            .collect();
        let v = first_line[0];
        let adj = vec![vec![0; v]; v];

        let e = first_line[1];
        let mut ret = AdjMatrix { v, e, adj };
        for _ in 0..e {
            let nums: Vec<usize> = lines
                .next()
                .unwrap()
                .split_whitespace()
                .map(|x| x.parse().unwrap())
                .collect();
            let (a, b) = (nums[0], nums[1]);
            ret.validate_vertex(a);
            ret.validate_vertex(b);
            if a == b {
                panic!("Self Loop is Detected!");
            }
            if ret.adj[a][b] == 1 {
                panic!("Parallel Edges are Detected!");
            }
            ret.adj[a][b] = 1;
            ret.adj[b][a] = 1;
        }
        ret
    }

    fn validate_vertex(&self, v: usize) {
        if v >= self.v {
            panic!("vertex {v} is invalid");
        }
    }

    pub fn v(&self) -> i32 {
        self.v as i32
    }

    pub fn e(&self) -> i32 {
        self.e as i32
    }

    pub fn has_edge(&self, v1: usize, v2: usize) -> bool {
        self.validate_vertex(v1);
        self.validate_vertex(v2);
        self.adj[v1][v2] == 1
    }

    pub fn adj_edge(&self, v: usize) -> Vec<usize> {
        self.validate_vertex(v);
        let mut res = vec![];
        for i in 0..self.v {
            if self.adj[v][i] == 1 {
                res.push(i);
            }
        }
        res

        //        self.adj[v]
        //            .iter()
        //            .enumerate()
        //            .filter_map(|(i, &w)| if w == 1 { Some(i) } else { None })
        //            .collect()
    }

    pub fn degree(&self, v: usize) -> usize {
        self.adj_edge(v).len()
    }
}

impl Display for AdjMatrix {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "V = {}, E = {}\n", self.v, self.e)?;
        for i in 0..self.v {
            for j in 0..self.v {
                write!(f, "{} ", self.adj[i][j])?;
            }
            writeln!(f)?;
        }
        Ok(())
    }
}

fn main() {
    let adj = AdjMatrix::new("../g.txt");
    println!("{adj}");
    println!("{:?}", adj.adj_edge(3));
}
