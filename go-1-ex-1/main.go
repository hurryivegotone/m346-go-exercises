package main

import "fmt"

var firstName, lastName string = "Gregor", "Samsa"
var dayOfBirth, monthOfBirth, yearOfBirth int16 = 7, 3, 1884
var numberOfSiblings int8 = 1
var heightInMeters float32 = 1.76
var zodiacSign rune = '\u2653'

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
