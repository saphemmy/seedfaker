package seedfaker

import "fmt"

var colorNames = []string{
	"Red", "Green", "Blue", "Yellow", "Orange", "Purple", "Pink", "Brown",
	"Black", "White", "Gray", "Cyan", "Magenta", "Teal", "Navy", "Maroon",
}

// HexColor returns a CSS hex colour such as "#1a2b3c".
func (f *Faker) HexColor() string { return fmt.Sprintf("#%06x", f.IntN(1<<24)) }

// RGB returns red, green and blue components, each 0-255.
func (f *Faker) RGB() (r, g, b int) { return f.IntN(256), f.IntN(256), f.IntN(256) }

// RGBString returns an "rgb(r, g, b)" string.
func (f *Faker) RGBString() string {
	r, g, b := f.RGB()
	return fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)
}

// ColorName returns a colour name.
func (f *Faker) ColorName() string { return f.pick(colorNames) }
