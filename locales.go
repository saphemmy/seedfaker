package seedfaker

// EN is the default, locale-neutral English profile, used when New is given no
// WithLocale. Country-specific locales live in the locale/ subpackages (locale/ng,
// locale/us, ...); import the ones you need.
var EN = Locale{
	Code:    "en",
	Country: "United States",
	FirstNames: []string{
		"James", "Mary", "John", "Patricia", "Robert", "Jennifer", "Michael", "Linda",
		"David", "Elizabeth", "William", "Susan", "Joseph", "Jessica", "Thomas", "Sarah",
	},
	LastNames: []string{
		"Smith", "Johnson", "Williams", "Brown", "Jones", "Miller", "Davis", "Wilson",
		"Taylor", "Anderson", "Thomas", "Moore", "Martin", "Lee", "Clark", "Lewis",
	},
	Titles:          []string{"Mr", "Mrs", "Ms", "Miss", "Dr", "Prof"},
	PhoneFormats:    []string{"+1 ### ### ####"},
	Cities:          []string{"Springfield", "Riverside", "Franklin", "Clinton", "Georgetown", "Salem", "Fairview", "Madison", "Arlington", "Ashland"},
	States:          []string{"California", "Texas", "Florida", "New York", "Illinois", "Ohio", "Georgia", "Washington"},
	Streets:         []string{"Main Street", "Oak Avenue", "Maple Drive", "Washington Boulevard", "Park Avenue", "Elm Street", "Cedar Lane", "Lake Street"},
	PostcodeFormats: []string{"#####"},
}

func init() { Register(EN) }
