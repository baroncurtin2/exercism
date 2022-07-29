package expenses

import (
	"errors"
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	records := []Record{}

	for _, x := range in {
		if predicate(x) {
			records = append(records, x)
		}
	}

	return records
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	records := Filter(in, ByDaysPeriod(p))

	amount := 0.0

	for _, r := range records {
		amount += r.Amount
	}

	return amount
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	filteredCatRecords := Filter(in, ByCategory(c))

	if len(filteredCatRecords) == 0 {
		errMsg := fmt.Sprintf("unknown category %s", c)
		return 0.0, errors.New(errMsg)
	}

	filteredPeriodRecords := Filter(filteredCatRecords, ByDaysPeriod(p))

	if len(filteredPeriodRecords) == 0 {
		return 0.0, nil
	}

	amount := 0.0
	for _, r := range filteredPeriodRecords {
		amount += r.Amount
	}

	return amount, nil
}
