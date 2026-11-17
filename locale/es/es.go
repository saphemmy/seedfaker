// Package es provides the Spain locale for seedfaker.
package es

import "github.com/saphemmy/seedfaker"

// Locale is Spain: +34 phones, 5-digit postal codes.
var Locale = seedfaker.Locale{
	Code:            "es_ES",
	Country:         "Spain",
	FirstNames:      []string{"Lucia", "Hugo", "Martina", "Mateo", "Sofia", "Daniel", "Paula", "Pablo", "Valeria", "Alvaro", "Carmen", "Diego", "Sara", "Adrian", "Marta"},
	LastNames:       []string{"Garcia", "Rodriguez", "Gonzalez", "Fernandez", "Lopez", "Martinez", "Sanchez", "Perez", "Gomez", "Ruiz", "Diaz", "Moreno", "Alvarez", "Romero"},
	Titles:          []string{"Sr", "Sra", "Srta", "Dr", "Dra"},
	PhoneFormats:    []string{"+34 6## ## ## ##", "+34 7## ## ## ##"},
	Cities:          []string{"Madrid", "Barcelona", "Valencia", "Sevilla", "Zaragoza", "Malaga", "Murcia", "Palma", "Bilbao", "Alicante"},
	States:          []string{"Madrid", "Cataluna", "Andalucia", "Valencia", "Galicia", "Pais Vasco", "Castilla y Leon", "Canarias"},
	Streets:         []string{"Calle Mayor", "Gran Via", "Paseo de Gracia", "Calle de Alcala", "Calle Serrano", "Rambla de Catalunya"},
	PostcodeFormats: []string{"#####"},
}

func init() { seedfaker.Register(Locale) }
