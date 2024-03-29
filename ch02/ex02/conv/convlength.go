package conv

import "fmt"

type Meter float64
type Feet float64

func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }

// MToFt はメートルの長さをフィートへ変換します。
func MToFt(m Meter) Feet { return Feet(m / 0.3048) }

// FtToM はフィートの長さをメートルへ変換します。
func FtToM(f Feet) Meter { return Meter(f * 0.3048) }
