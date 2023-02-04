// Throne Inheritance
// leetcode: https://leetcode-cn.com/problems/throne-inheritance/
use std::collections::{HashMap, HashSet};

struct ThroneInheritance {
    king: String,
    persons: HashMap<String, Vec<String>>,
    death: HashSet<String>
}

impl ThroneInheritance {
    fn new(kingName: String) -> Self {
        let mut ti = ThroneInheritance {
            persons: HashMap::new(),
            death: HashSet::new(),
            king: kingName.clone()
        };
        ti.persons.insert(kingName, Vec::new());
        ti
    }
    
    fn birth(&mut self, parent_name: String, child_name: String) {
        self.persons.insert(child_name.clone(), Vec::new());
        self.persons.get_mut(&parent_name).unwrap().push(child_name);
    }
    
    fn death(&mut self, name: String) {
        self.death.insert(name);
    }
    
    fn get_inheritance_order(&self) -> Vec<String> {
        let mut v = Vec::new();
        self.preorder(&self.king, &mut v);
        return v;
    }

    fn preorder(&self, name: &String, v: &mut Vec<String>) {
        if !self.death.contains(name) {
            v.push(name.clone());
        }
        if self.persons.contains_key(name) {
            for child in &self.persons[name] {
                self.preorder(child, v);
            }
        }
    }
}

fn main() {
    let mut obj = ThroneInheritance::new(String::from("xdd"));
    obj.birth(String::from("xdd"), String::from("123123123"));
}