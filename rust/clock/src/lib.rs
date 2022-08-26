use std::fmt::{Display, Formatter};

const MINS_IN_HOUR: i32 = 60;
const MINS_IN_DAY: i32 = 24 * MINS_IN_HOUR;

#[derive(Debug, Eq, PartialEq)]
pub struct Clock {
    minutes: i32,
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let total_minutes = hours * MINS_IN_HOUR + minutes;

        Self {
            minutes: normalize_minutes(&total_minutes)
        }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        let total_minutes = self.minutes + minutes;

        Self {
            minutes: normalize_minutes(&total_minutes)
        }
    }
}


impl Display for Clock {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(f, "{:02}:{:02}", self.minutes / 60, self.minutes % 60)
    }
}

fn normalize_minutes(minutes: &i32) -> i32 {
    if minutes < &0 {
        return &MINS_IN_DAY + (minutes % MINS_IN_DAY);
    } else if minutes >= &MINS_IN_DAY {
        return minutes % &MINS_IN_DAY;
    } else {
        *minutes
    }
}
