// Throne Inheritance
// leetcode: https://leetcode-cn.com/problems/throne-inheritance/
use std::collections::HashMap;

struct ThroneInheritance {
    king: String,
    edges: HashMap<String, Vec<String>>,
    death: HashMap<String, bool>,
}

impl ThroneInheritance {

    fn new(kingName: String) -> Self {
        let edges: HashMap<String, Vec<String>> = HashMap::new();
        let death: HashMap<String, bool> = HashMap::new();
        ThroneInheritance{
            edges,
            death,
            king: kingName
        }
    }
    
    fn birth(&mut self, parent_name: String, child_name: String) {
        self.edges.get(&parent_name).unwrap().push(child_name);
        println!("{:?}", self.edges.get(&parent_name).unwrap());
    }
    
    fn death(&mut self, name: String) {
        self.death.insert(name, true);
    }
    
}

fn main() {
    let mut obj = ThroneInheritance::new(String::from("xdd"));
    obj.birth(String::from("xdd"), String::from("123123123"));
}