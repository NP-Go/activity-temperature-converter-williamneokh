package main

import (
	"fmt"
)

func convertKelvin(temperatureInput float64) (float64, float64) {
	//Insert your code here

	resultFahrenheit := temperatureInput*9/5 - 459.67
	resultCelsius := temperatureInput - 273.150

	//Do not remove this line
	return resultFahrenheit, resultCelsius
}

func convertCelsius(temperatureInput float64) (float64, float64) {
	//Insert your code here
	resultFahrenheit := temperatureInput*9/5 + 32
	resultKelvin := temperatureInput + 273.15

	//Do not remove this line
	return resultFahrenheit, resultKelvin
}

func convertFahrenheit(temperatureInput float64) (float64, float64) {
	//Insert your code here
	resultKelvin := (temperatureInput + 459.67) * 5 / 9
	resultCelsius := (temperatureInput - 32) * 5 / 9
	//Do not remove this line
	return resultKelvin, resultCelsius
}

func main() {
	var temperatureChoice int
	var temperatureInput float64
	fmt.Println("Enter your temperature format (1 for Kelvin, 2 for Celsius, 3 for Fahrenheit: ")
	fmt.Scanln(&temperatureChoice)
	fmt.Println("Enter the temperature: ")
	fmt.Scanln(&temperatureInput)

	if temperatureChoice == 1 {
		//Insert Code here
		resultFahrenheit, resultCelsius := convertKelvin(temperatureInput)
		//DO not remove this
		fmt.Println("Fahrenheit: ", resultFahrenheit, " Celsius: ", resultCelsius)
	} else if temperatureChoice == 2 {
		//Insert Code here
		resultFahrenheit, resultKelvin := convertCelsius(temperatureInput)
		//DO not remove this
		fmt.Println("Fahrenheit: ", resultFahrenheit, " Kelvin: ", resultKelvin)
	} else if temperatureChoice == 3 {
		//Insert Code here
		resultKelvin, resultCelsius := convertFahrenheit(temperatureInput)
		//DO not remove this
		fmt.Println("Kelvin: ", resultKelvin, " Celsius: ", resultCelsius)
	} else {
		fmt.Println("No Conversion")
	}

}
