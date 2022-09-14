use std::collections::{HashSet, HashMap};

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &[&'a str]) -> HashSet<&'a str> {
    if word.len() == 0 || possible_anagrams.len() == 0 {
        return HashSet::new();
    }
    let word_lower = word.to_lowercase();
    let word_chars_count: HashMap<char, u8> = count_chars(word);

    possible_anagrams
        .iter()
        .filter(|candidate| {
            let word_can = candidate.to_lowercase();
            word_can != word_lower && count_chars(&word_can) == word_chars_count
        })
        .copied()
        .collect::<HashSet<&str>>()
}

fn count_chars(word: &str) -> HashMap<char, u8> {
    word
        .to_lowercase()
        .chars()
        .fold(HashMap::new(), |mut map, c| {
            *map.entry(c).or_insert(0) += 1;
            map
        })
}
