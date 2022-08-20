// This stub file contains items that aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

use std::collections::HashMap;

pub fn can_construct_note(magazine: &[&str], note: &[&str]) -> bool {
    let mag_map = generate_hash_map(magazine);
    let note_map = generate_hash_map(note);

    for (word, count) in &note_map {
        match mag_map.get(word) {
            Some(mag_count) => {
                if count > mag_count {
                    return false;
                }
            }
            None => { return false; }
        }
    }

    true
}

pub fn generate_hash_map<'a>(text: &'a [&str]) -> HashMap<&'a str, u32> {
    text.iter().fold(HashMap::new(), |mut acc, word| {
        *acc.entry(word).or_insert(0) += 1;
        acc
    })
}
