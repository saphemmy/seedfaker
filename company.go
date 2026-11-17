package seedfaker

var (
	companySuffixes = []string{"Inc", "LLC", "Ltd", "Group", "PLC", "Co", "Partners", "Holdings", "Systems", "Solutions"}
	jobTitles       = []string{
		"Software Engineer", "Product Manager", "Data Analyst", "Accountant", "Designer",
		"Sales Executive", "Operations Manager", "Marketing Lead", "Consultant",
		"Project Manager", "Business Analyst", "HR Manager", "Support Specialist", "Architect",
	}
	catchAdjective = []string{"Adaptive", "Integrated", "Streamlined", "Robust", "Scalable", "Proactive", "Customer-focused", "Open-source"}
	catchNoun1     = []string{"analyzing", "assurance", "benchmark", "capability", "framework", "infrastructure", "orchestration", "paradigm"}
	catchNoun2     = []string{"solutions", "systems", "networks", "initiatives", "platforms", "synergies", "architectures", "methodologies"}
	bsVerb         = []string{"implement", "leverage", "optimize", "streamline", "orchestrate", "reinvent", "aggregate", "transform"}
	bsAdjective    = []string{"world-class", "cross-platform", "scalable", "end-to-end", "next-generation", "frictionless", "holistic", "dynamic"}
	bsNoun         = []string{"experiences", "platforms", "channels", "ecosystems", "pipelines", "deliverables", "bandwidth", "paradigms"}
)

// Company returns a company name, e.g. "Okafor Systems".
func (f *Faker) Company() string { return f.LastName() + " " + f.pick(companySuffixes) }

// CompanySuffix returns a legal suffix such as "LLC".
func (f *Faker) CompanySuffix() string { return f.pick(companySuffixes) }

// JobTitle returns a job title.
func (f *Faker) JobTitle() string { return f.pick(jobTitles) }

// CatchPhrase returns a marketing catch phrase.
func (f *Faker) CatchPhrase() string {
	return f.pick(catchAdjective) + " " + f.pick(catchNoun1) + " " + f.pick(catchNoun2)
}

// Bs returns a line of corporate buzzword filler.
func (f *Faker) Bs() string {
	return f.pick(bsVerb) + " " + f.pick(bsAdjective) + " " + f.pick(bsNoun)
}
