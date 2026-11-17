package seedfaker

import "strings"

var loremWords = []string{
	"lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit",
	"sed", "do", "eiusmod", "tempor", "incididunt", "ut", "labore", "et", "dolore",
	"magna", "aliqua", "enim", "minim", "veniam", "quis", "nostrud", "exercitation",
	"ullamco", "laboris", "nisi", "aliquip", "commodo", "consequat", "duis", "aute",
}

// Word returns one lorem word.
func (f *Faker) Word() string { return f.pick(loremWords) }

// Words returns n lorem words joined by spaces, or "" if n <= 0.
func (f *Faker) Words(n int) string {
	if n <= 0 {
		return ""
	}
	p := make([]string, n)
	for i := range p {
		p[i] = f.Word()
	}
	return strings.Join(p, " ")
}

// Sentence returns a capitalised sentence of n words ending in a period; n <= 0 gives 6.
func (f *Faker) Sentence(n int) string {
	if n <= 0 {
		n = 6
	}
	s := f.Words(n)
	return strings.ToUpper(s[:1]) + s[1:] + "."
}

// Sentences returns n sentences joined by spaces; n <= 0 gives 3.
func (f *Faker) Sentences(n int) string {
	if n <= 0 {
		n = 3
	}
	p := make([]string, n)
	for i := range p {
		p[i] = f.Sentence(f.IntBetween(4, 10))
	}
	return strings.Join(p, " ")
}

// Paragraph returns a paragraph of n sentences; n <= 0 gives 3.
func (f *Faker) Paragraph(n int) string {
	if n <= 0 {
		n = 3
	}
	return f.Sentences(n)
}

// Paragraphs returns n paragraphs joined by blank lines; n <= 0 gives 3.
func (f *Faker) Paragraphs(n int) string {
	if n <= 0 {
		n = 3
	}
	p := make([]string, n)
	for i := range p {
		p[i] = f.Paragraph(f.IntBetween(2, 5))
	}
	return strings.Join(p, "\n\n")
}

// Text returns lorem text roughly maxChars long (at least one word); it panics if
// maxChars < 1.
func (f *Faker) Text(maxChars int) string {
	if maxChars < 1 {
		panic("seedfaker: Text requires maxChars >= 1")
	}
	var b strings.Builder
	for {
		w := f.Word()
		if b.Len() > 0 && b.Len()+1+len(w) > maxChars {
			break
		}
		if b.Len() > 0 {
			b.WriteByte(' ')
		}
		b.WriteString(w)
		if b.Len() >= maxChars {
			break
		}
	}
	return b.String()
}
