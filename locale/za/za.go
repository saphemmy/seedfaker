// Package za provides the South Africa locale for seedfaker.
package za

import "github.com/saphemmy/seedfaker"

// Locale is South Africa: +27 phones, 4-digit postal codes.
var Locale = seedfaker.Locale{
	Code:    "en_ZA",
	Country: "South Africa",
	FirstNames: []string{
		"Thabo", "Lerato", "Sipho", "Nomsa", "Mandla", "Zanele", "Themba", "Nandi",
		"Bongani", "Palesa", "Lwazi", "Ayanda", "Kagiso", "Thandiwe", "Siphokazi",
	},
	LastNames:       []string{"Nkosi", "Dlamini", "Khumalo", "Mthembu", "Ndlovu", "Zulu", "Mokoena", "Botha", "Naidoo", "Van der Merwe"},
	Titles:          []string{"Mr", "Mrs", "Miss", "Dr"},
	PhoneFormats:    []string{"+27 ## ### ####"},
	Cities:          []string{"Johannesburg", "Cape Town", "Durban", "Pretoria", "Gqeberha", "Bloemfontein", "East London", "Pietermaritzburg", "Polokwane", "Nelspruit"},
	States:          []string{"Gauteng", "Western Cape", "KwaZulu-Natal", "Eastern Cape", "Free State", "Limpopo", "Mpumalanga", "North West", "Northern Cape"},
	Streets:         []string{"Church Street", "Long Street", "Commissioner Street", "West Street", "Jan Smuts Avenue", "Rivonia Road"},
	PostcodeFormats: []string{"####"},
}

func init() { seedfaker.Register(Locale) }
