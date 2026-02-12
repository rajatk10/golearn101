package main

import "fmt"

func areaCircle(a float32) float64 {
	return 3.14 * float64(a) * float64(a)
}
func celsiusToFahrenheit(c float64) float64 {
	return ((9.0/5.0)*c + 32)
	//It will do integer division is 9/5
}

func main() {
	fmt.Printf("Area of circle with radius %f is : %.5f \n", 3.69, areaCircle(3.69))
	fmt.Printf("Convert celsius %f to Fahrenheit : %.5f", 37.0, celsiusToFahrenheit(37))
}
