package main

import "fmt"

//declaring a new named type using the format: type name underlying-type
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	fmt.Println(Celsius.String(BoilingC))
}

func CToF(c Celsius) Fahrenheit {
	//Fahrenheit() is not a function its a conversion.  You can only convert named types only if the underline type is the same
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

//Creating a type 'method',  do assign behaviors specific to the named type
//method declaration goes as follows: func (var Type) methodName(params) (return values) {body}
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
