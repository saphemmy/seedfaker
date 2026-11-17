// Package us provides the United States locale for seedfaker.
package us

import "github.com/saphemmy/seedfaker"

// Locale is the United States: +1 phones, 5-digit ZIP and ZIP+4 codes.
var Locale = seedfaker.Locale{
	Code:            "en_US",
	Country:         "United States",
	FirstNames:      []string{"James", "Mary", "John", "Patricia", "Robert", "Jennifer", "Michael", "Linda", "David", "Elizabeth", "Richard", "Barbara", "Joseph", "Susan", "Thomas"},
	LastNames:       []string{"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez", "Wilson", "Anderson", "Taylor", "Thomas", "Moore"},
	Titles:          []string{"Mr", "Mrs", "Ms", "Miss", "Dr", "Prof"},
	PhoneFormats:    []string{"+1 (###) ###-####"},
	Cities:          []string{"New York", "Los Angeles", "Chicago", "Houston", "Phoenix", "Philadelphia", "San Antonio", "San Diego", "Dallas", "Austin"},
	States:          []string{"California", "Texas", "Florida", "New York", "Pennsylvania", "Illinois", "Ohio", "Georgia", "North Carolina", "Michigan"},
	Streets:         []string{"Main Street", "Oak Avenue", "Maple Drive", "Washington Boulevard", "Park Avenue", "Elm Street", "Cedar Lane", "Lake Street"},
	PostcodeFormats: []string{"#####", "#####-####"},
}

func init() { seedfaker.Register(Locale) }
