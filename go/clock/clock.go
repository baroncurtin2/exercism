package clock

import "fmt"

const (
	dayMinutes  int = 60 * 24
	hourMinutes int = 60
)

// Clock Define the Clock type here.
type Clock struct {
	minutes int
}

func normalize(hours int, minutes int) Clock {
	totalMinutes := hours*hourMinutes + minutes

	if totalMinutes < 0 {
		return Clock{dayMinutes + (totalMinutes % dayMinutes)}
	} else if totalMinutes >= dayMinutes {
		return Clock{totalMinutes % dayMinutes}
	}

	return Clock{totalMinutes}
}

func New(h, m int) Clock {
	return normalize(h, m)
}

func (c Clock) Add(m int) Clock {
	return normalize(0, c.minutes+m)
}

func (c Clock) Subtract(m int) Clock {
	return normalize(0, c.minutes-m)
}

func (c Clock) String() string {
	h := c.minutes / hourMinutes
	m := c.minutes % hourMinutes

	return fmt.Sprintf("%02d:%02d", h, m)
}
