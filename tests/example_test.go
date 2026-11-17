package tests

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/saphemmy/seedfaker"
	"github.com/saphemmy/seedfaker/locale/ng"
)

// The core guarantee: the same seed yields the same data on every run.
func Example() {
	f := seedfaker.New(42)
	first, second := seedfaker.New(42).Name(), seedfaker.New(42).Name()
	fmt.Println(first == second)
	fmt.Println(len(f.UUID()))
	// Output:
	// true
	// 36
}

func Example_locale() {
	f := seedfaker.New(1, seedfaker.WithLocale(ng.Locale))
	fmt.Println(strings.HasPrefix(f.Phone(), "+234"))
	// Output: true
}

func ExampleFactory() {
	type User struct{ Name, Email string }
	fac := seedfaker.Factory[User]{Build: func(f *seedfaker.Faker) User {
		return User{Name: f.Name(), Email: f.Email()}
	}}
	a := fac.Make(seedfaker.New(7), 3)
	b := fac.Make(seedfaker.New(7), 3)
	fmt.Println(reflect.DeepEqual(a, b))
	// Output: true
}
