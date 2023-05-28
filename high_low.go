package temperature

import "time"

// HighLow contains a high and low temperature
type HighLow struct {
	High *Temperature
	Low  *Temperature
}

func NewHighLow() *HighLow {
	return &HighLow{
		High: MinTemperature(),
		Low:  MaxTemperature(),
	}
}

// TimeMap is a map of temperatures over time
type TimeMap map[time.Time]*Temperature

// HighLow Creates a HighLow from a TimeMap
func (tm TimeMap) HighLow() *HighLow {
	hl := NewHighLow()
	for _, t := range tm {
		if t.GreaterThan(hl.High) {
			hl.High = t
		}
		if t.LessThan(hl.Low) {
			hl.Low = t
		}
	}
	return hl
}
