// Package gh provides the Ghana locale for seedfaker.
package gh

import "github.com/saphemmy/seedfaker"

// Locale is Ghana: Akan, Ewe and Ga names, +233 phones, GhanaPostGPS digital codes.
var Locale = seedfaker.Locale{
	Code:    "en_GH",
	Country: "Ghana",
	FirstNames: []string{
		"Kwame", "Ama", "Kofi", "Akosua", "Yaw", "Abena", "Kwaku", "Esi", "Kojo",
		"Adwoa", "Kwabena", "Akua", "Yawo", "Afua", "Kwasi", "Afia", "Kodjo",
		"Akofa", "Mawuli", "Dela", "Nii", "Naa", "Fiifi", "Araba", "Ekow",
	},
	LastNames: []string{
		"Mensah", "Osei", "Boateng", "Owusu", "Asante", "Agyeman", "Addo", "Appiah",
		"Darko", "Antwi", "Ofori", "Adjei", "Amoah", "Danso", "Acheampong",
		"Nkrumah", "Quartey", "Tetteh", "Lamptey", "Annan",
	},
	Titles:       []string{"Mr", "Mrs", "Miss", "Dr", "Prof", "Nana"},
	PhoneFormats: []string{"+233 24# ### ###", "+233 20# ### ###", "+233 54# ### ###", "+233 55# ### ###"},
	Cities: []string{
		"Accra", "Kumasi", "Tamale", "Takoradi", "Cape Coast", "Tema", "Sekondi",
		"Koforidua", "Sunyani", "Ho", "Techiman", "Obuasi", "Tarkwa", "Wa", "Bolgatanga",
	},
	States: []string{
		"Greater Accra", "Ashanti", "Northern", "Western", "Central", "Eastern",
		"Volta", "Bono", "Upper East", "Upper West",
	},
	Streets: []string{
		"Oxford Street", "Ring Road", "Liberation Road", "Independence Avenue",
		"High Street", "Spintex Road", "Cantonments Road", "Castle Road",
	},
	PostcodeFormats: []string{"??-###-####"}, // GhanaPostGPS: e.g. GA-543-0112
}

func init() { seedfaker.Register(Locale) }
