package seedfaker

// EAN13 returns a valid 13-digit EAN barcode.
func (f *Faker) EAN13() string { return f.ean(13) }

// EAN8 returns a valid 8-digit EAN barcode.
func (f *Faker) EAN8() string { return f.ean(8) }

// ISBN13 returns a valid 13-digit ISBN (978 prefix).
func (f *Faker) ISBN13() string {
	d := make([]int, 13)
	d[0], d[1], d[2] = 9, 7, 8
	for i := 3; i < 12; i++ {
		d[i] = f.IntN(10)
	}
	d[12] = eanCheckDigit(d[:12])
	return digitsToString(d)
}

// ISBN10 returns a valid 10-character ISBN; the check character may be 'X'.
func (f *Faker) ISBN10() string {
	d := make([]int, 9)
	for i := range d {
		d[i] = f.IntN(10)
	}
	sum := 0
	for i, x := range d {
		sum += (10 - i) * x
	}
	check := (11 - sum%11) % 11
	b := make([]byte, 10)
	for i, x := range d {
		b[i] = byte('0' + x)
	}
	if check == 10 {
		b[9] = 'X'
	} else {
		b[9] = byte('0' + check)
	}
	return string(b)
}

func (f *Faker) ean(n int) string {
	d := make([]int, n)
	for i := 0; i < n-1; i++ {
		d[i] = f.IntN(10)
	}
	d[n-1] = eanCheckDigit(d[:n-1])
	return digitsToString(d)
}

// eanCheckDigit computes the EAN/ISBN-13 check digit: weights 3 and 1 from the right.
func eanCheckDigit(body []int) int {
	sum, weight := 0, 3
	for i := len(body) - 1; i >= 0; i-- {
		sum += body[i] * weight
		weight = 4 - weight // 3 -> 1 -> 3
	}
	return (10 - sum%10) % 10
}

func digitsToString(d []int) string {
	b := make([]byte, len(d))
	for i, x := range d {
		b[i] = byte('0' + x)
	}
	return string(b)
}
