// This stub file contains items that aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

pub struct Player {
    pub health: u32,
    pub mana: Option<u32>,
    pub level: u32,
}

impl Player {
    pub fn new(level: u32) -> Self {
        Self {
            health: 100,
            mana: if level >= 10 { Some(100) } else { None },
            level,
        }
    }

    pub fn revive(&self) -> Option<Player> {
        if self.health == 0 {
            Some(Player::new(self.level))
        } else {
            None
        }
    }

    pub fn cast_spell(&mut self, mana_cost: u32) -> u32 {
        match self.mana {
            Some(mana) => {
                if mana < mana_cost {
                    0
                } else {
                    self.mana = Some(mana - mana_cost);
                    self.spell_damage(&mana_cost)
                }
            }
            None => {
                self.health = if self.health < mana_cost { 0 } else { self.health - mana_cost };
                0
            }
        }
    }

    fn spell_damage(&self, mana_cost: &u32) -> u32 {
        mana_cost * 2
    }
}
