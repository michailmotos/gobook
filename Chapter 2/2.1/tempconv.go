// Package tempcov performs Celsius & Fahrenheit conversions.

package tempconv

import "fmt"

//Celsius ...
type Celsius float64

//Fahrenheit ...
type Fahrenheit float64

//Kelvin ...
type Kelvin float64

//Block comment
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%gC", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%gF", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }
