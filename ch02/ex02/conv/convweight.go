package conv

import "fmt"

type Kilogram float64
type Pound float64

func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }
func (p Pound) String() string    { return fmt.Sprintf("%glb", p) }

// KToP はキログラムの重さをポンドへ変換します。
func KToP(k Kilogram) Pound { return Pound(k * 2.20462) }

// PToK はポンドの重さをキログラムへ変換します。
func PToK(p Pound) Kilogram { return Kilogram(p / 2.20462) }
