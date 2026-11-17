// Package in provides the India locale for seedfaker.
package in

import "github.com/saphemmy/seedfaker"

// Locale is India: +91 phones, 6-digit PIN codes.
var Locale = seedfaker.Locale{
	Code:            "en_IN",
	Country:         "India",
	FirstNames:      []string{"Aarav", "Priya", "Arjun", "Ananya", "Vihaan", "Diya", "Aditya", "Saanvi", "Reyansh", "Ishaan", "Myra", "Kabir", "Aisha", "Rohan", "Kavya"},
	LastNames:       []string{"Sharma", "Patel", "Singh", "Kumar", "Gupta", "Reddy", "Nair", "Iyer", "Mehta", "Rao", "Das", "Chopra", "Bose", "Malhotra"},
	Titles:          []string{"Mr", "Mrs", "Ms", "Dr", "Shri", "Smt"},
	PhoneFormats:    []string{"+91 ##### #####"},
	Cities:          []string{"Mumbai", "Delhi", "Bengaluru", "Hyderabad", "Chennai", "Kolkata", "Pune", "Ahmedabad", "Jaipur", "Surat"},
	States:          []string{"Maharashtra", "Delhi", "Karnataka", "Telangana", "Tamil Nadu", "West Bengal", "Gujarat", "Rajasthan", "Uttar Pradesh", "Kerala"},
	Streets:         []string{"MG Road", "Brigade Road", "Linking Road", "Park Street", "Anna Salai", "Marine Drive"},
	PostcodeFormats: []string{"######"},
}

func init() { seedfaker.Register(Locale) }
