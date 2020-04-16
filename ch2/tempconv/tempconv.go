package tempconv

/*

This package defines two types, Celsius and Fahrenheit,
for the two units of temperature.
Even though both have the same underlying type, float64,
they are not the same type, so they cannot be compared or combined in arithmetic
expressions. Distinguishing the types makes it possible to avoid errors like
inadvertently combining temperatures in the two different scales; an explicit
type conversion like Celsius(t) or Fahrenheit(t) is required to convert from a
float64. Celsius(t) and Fahrenheit(t) are conversions,
not function calls. They don’t change the value or representation in any way,
but they make the change of meaning explicit. On the other hand,
the function CToF and FToC convert between the two scales; they do return
different values.
 */
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
