package lenconv

import "fmt"

type Meter float64
type Feet float64

const feetInMeter = 3.28084

func (m Meter) String() string { return fmt.Sprintf("%g Meter", m) }
func (f Feet) String() string { return fmt.Sprintf("%g Feet", f) }
