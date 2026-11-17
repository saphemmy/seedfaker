// Package fr provides the France locale for seedfaker.
package fr

import "github.com/saphemmy/seedfaker"

// Locale is France: +33 phones, 5-digit postal codes.
var Locale = seedfaker.Locale{
	Code:            "fr_FR",
	Country:         "France",
	FirstNames:      []string{"Gabriel", "Emma", "Louis", "Jade", "Raphael", "Louise", "Jules", "Alice", "Adam", "Chloe", "Hugo", "Lea", "Arthur", "Manon", "Nathan"},
	LastNames:       []string{"Martin", "Bernard", "Dubois", "Thomas", "Robert", "Richard", "Petit", "Durand", "Leroy", "Moreau", "Simon", "Laurent", "Lefebvre", "Michel"},
	Titles:          []string{"M", "Mme", "Mlle", "Dr", "Pr"},
	PhoneFormats:    []string{"+33 6 ## ## ## ##", "+33 7 ## ## ## ##"},
	Cities:          []string{"Paris", "Marseille", "Lyon", "Toulouse", "Nice", "Nantes", "Strasbourg", "Montpellier", "Bordeaux", "Lille"},
	States:          []string{"Ile-de-France", "Provence-Alpes-Cote d'Azur", "Auvergne-Rhone-Alpes", "Occitanie", "Nouvelle-Aquitaine", "Grand Est", "Hauts-de-France"},
	Streets:         []string{"Rue de la Paix", "Avenue des Champs-Elysees", "Boulevard Saint-Germain", "Rue de Rivoli", "Rue du Faubourg", "Place de la Republique"},
	PostcodeFormats: []string{"#####"},
}

func init() { seedfaker.Register(Locale) }
