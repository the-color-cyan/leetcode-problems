use std::collections::HashMap;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut map: HashMap<i32, i32> = HashMap::new();

        for (i, n) in nums.iter().enumerate() {
            let diff = target - n;
            let diff_index = map.get(&diff);

            match diff_index {
                Some(j) => return vec![*j, i as i32],
                None => map.insert(*n, i as i32),
            };
        }

        vec![]
    }
}
