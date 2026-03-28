package valueobjects

import "fmt"

type Percentage struct {
	value float64
}

func NewPercentage(value float64) (Percentage, error) {
	if value < 0 || value > 100 {
		return Percentage{}, fmt.Errorf("percentage must be 0-100")
	}
	return Percentage{value: value}, nil
}

func (p Percentage) Value() float64 {
	return p.value
}
