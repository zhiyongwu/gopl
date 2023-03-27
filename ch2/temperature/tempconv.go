package temperature

import "fmt"

type Celsius float64

type Fahrenheit float64

type Kelvin float64

const (
	AbsoluteZero Celsius = -273.15
	FreezingC    Celsius = 0
	Boiling      Celsius = 100

	AbsoluteKelvin Kelvin = -273.15
)

func (c Celsius) String() string {
	fmt.Println("String() called")
	return fmt.Sprintf("%g.C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g.F", f)
}
