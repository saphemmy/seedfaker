// Package ke provides the Kenya locale for seedfaker.
package ke

import "github.com/saphemmy/seedfaker"

// Locale is Kenya: +254 phones, 5-digit postal codes.
var Locale = seedfaker.Locale{
	Code:    "en_KE",
	Country: "Kenya",
	FirstNames: []string{
		"Wanjiru", "Kamau", "Achieng", "Otieno", "Njeri", "Mwangi", "Akinyi", "Kipchoge",
		"Wambui", "Omondi", "Nyambura", "Kibet", "Adhiambo", "Mutua", "Chebet",
	},
	LastNames:       []string{"Kamau", "Otieno", "Mwangi", "Omondi", "Kiprotich", "Wanjiku", "Njoroge", "Ochieng", "Maina", "Kiplagat"},
	Titles:          []string{"Mr", "Mrs", "Miss", "Dr"},
	PhoneFormats:    []string{"+254 7## ### ###"},
	Cities:          []string{"Nairobi", "Mombasa", "Kisumu", "Nakuru", "Eldoret", "Thika", "Malindi", "Kitale", "Garissa", "Nyeri"},
	States:          []string{"Nairobi", "Mombasa", "Kisumu", "Nakuru", "Uasin Gishu", "Kiambu", "Machakos", "Kilifi", "Nyeri", "Kakamega"},
	Streets:         []string{"Moi Avenue", "Kenyatta Avenue", "Tom Mboya Street", "Haile Selassie Avenue", "Ngong Road", "Uhuru Highway"},
	PostcodeFormats: []string{"#####"},
}

func init() { seedfaker.Register(Locale) }
