// Package de provides the Germany locale for seedfaker.
package de

import "github.com/saphemmy/seedfaker"

// Locale is Germany: +49 phones, 5-digit postal codes.
var Locale = seedfaker.Locale{
	Code:            "de_DE",
	Country:         "Germany",
	FirstNames:      []string{"Ben", "Emma", "Paul", "Mia", "Leon", "Hannah", "Finn", "Emilia", "Elias", "Lina", "Noah", "Marie", "Luis", "Lena", "Felix"},
	LastNames:       []string{"Muller", "Schmidt", "Schneider", "Fischer", "Weber", "Meyer", "Wagner", "Becker", "Schulz", "Hoffmann", "Koch", "Bauer", "Richter", "Klein"},
	Titles:          []string{"Herr", "Frau", "Dr", "Prof"},
	PhoneFormats:    []string{"+49 1## #######", "+49 30 ########"},
	Cities:          []string{"Berlin", "Hamburg", "Munich", "Cologne", "Frankfurt", "Stuttgart", "Dusseldorf", "Leipzig", "Dortmund", "Essen"},
	States:          []string{"Bavaria", "Berlin", "Hamburg", "Hesse", "Saxony", "North Rhine-Westphalia", "Baden-Wurttemberg", "Lower Saxony"},
	Streets:         []string{"Hauptstrasse", "Bahnhofstrasse", "Schulstrasse", "Gartenstrasse", "Dorfstrasse", "Bergstrasse"},
	PostcodeFormats: []string{"#####"},
}

func init() { seedfaker.Register(Locale) }
