// Package gb provides the United Kingdom locale for seedfaker.
package gb

import "github.com/saphemmy/seedfaker"

// Locale is the United Kingdom: +44 phones and the six real postcode outward-code
// shapes (A9, A99, A9A, AA9, AA99, AA9A), each with the "9AA" inward code.
var Locale = seedfaker.Locale{
	Code:         "en_GB",
	Country:      "United Kingdom",
	FirstNames:   []string{"Oliver", "Olivia", "Harry", "Amelia", "Jack", "Isla", "George", "Ava", "Noah", "Emily", "Charlie", "Sophia", "Leo", "Grace", "Oscar"},
	LastNames:    []string{"Smith", "Jones", "Taylor", "Brown", "Williams", "Wilson", "Evans", "Thomas", "Roberts", "Walker", "Wright", "Thompson", "White", "Hughes"},
	Titles:       []string{"Mr", "Mrs", "Ms", "Miss", "Dr", "Prof"},
	PhoneFormats: []string{"+44 7### ######", "+44 20 #### ####"},
	Cities:       []string{"London", "Birmingham", "Manchester", "Leeds", "Liverpool", "Sheffield", "Bristol", "Glasgow", "Edinburgh", "Cardiff"},
	States:       []string{"Greater London", "West Midlands", "Greater Manchester", "West Yorkshire", "Merseyside", "Kent", "Essex", "Hampshire"},
	Streets:      []string{"High Street", "Station Road", "Church Lane", "Victoria Road", "Mill Lane", "Park Road", "Queens Road", "Kings Road"},
	PostcodeFormats: []string{
		"?# #??", "?## #??", "?#? #??", "??# #??", "??## #??", "??#? #??",
	},
}

func init() { seedfaker.Register(Locale) }
