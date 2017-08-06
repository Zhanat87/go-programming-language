package lenconv

// MToF converts a Meter length to Feet.
func MToF(m Meter) Feet { return Feet(m * feetInMeter) }

// FToM converts a Feet length to Meter.
func FToM(f Feet) Meter { return Meter(f / feetInMeter) }
