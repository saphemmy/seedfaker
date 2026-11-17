// Package ng provides the Nigeria locale for seedfaker.
package ng

import "github.com/saphemmy/seedfaker"

// Locale is Nigeria: Yoruba, Igbo and Hausa names, +234 phones, 6-digit NIPOST codes.
var Locale = seedfaker.Locale{
	Code:    "en_NG",
	Country: "Nigeria",
	FirstNames: []string{
		"Oluwafemi", "Chidi", "Ngozi", "Emeka", "Aisha", "Tunde", "Folake", "Obinna", "Amaka",
		"Ibrahim", "Fatima", "Yusuf", "Chioma", "Oluwaseun", "Adaeze", "Chinedu",
		"Halima", "Babajide", "Nneka", "Suleiman", "Temitope", "Ifeoma", "Musa",
		"Bukola", "Chukwuma", "Zainab", "Ayodeji", "Uchenna", "Funmilayo", "Kelechi", "Sadiq",
	},
	LastNames: []string{
		"Okafor", "Adeyemi", "Okonkwo", "Abubakar", "Eze", "Balogun", "Nwosu", "Bello",
		"Afolabi", "Danjuma", "Adeyinka", "Okeke", "Oladipo", "Chukwu", "Mohammed",
		"Olawale", "Nwachukwu", "Adebayo", "Obi", "Yakubu",
	},
	Titles:       []string{"Mr", "Mrs", "Miss", "Dr", "Prof", "Chief", "Alhaji", "Hajiya"},
	PhoneFormats: []string{"+234 70# ### ####", "+234 80# ### ####", "+234 81# ### ####", "+234 90# ### ####", "+234 91# ### ####"},
	Cities: []string{
		"Lagos", "Abuja", "Kano", "Ibadan", "Port Harcourt", "Benin City", "Kaduna",
		"Enugu", "Abeokuta", "Onitsha", "Jos", "Ilorin", "Owerri", "Uyo", "Warri",
	},
	States: []string{
		"Lagos", "Abuja", "Kano", "Oyo", "Rivers", "Edo", "Kaduna", "Enugu", "Ogun",
		"Anambra", "Plateau", "Kwara", "Imo", "Akwa Ibom", "Delta",
	},
	Streets: []string{
		"Allen Avenue", "Awolowo Road", "Herbert Macaulay Way", "Adeniran Ogunsanya Street",
		"Ahmadu Bello Way", "Ikorodu Road", "Aminu Kano Crescent", "Nnamdi Azikiwe Street",
		"Marina Street", "Broad Street",
	},
	PostcodeFormats: []string{"######"}, // NIPOST: 6 digits
}

func init() { seedfaker.Register(Locale) }
