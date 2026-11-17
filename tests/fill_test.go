package tests

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/saphemmy/seedfaker"
)

type address struct {
	City string `faker:"city"`
	Zip  string
}

type user struct {
	ID        string `faker:"uuid"`
	Email     string `faker:"email"`
	Name      string `faker:"name"`
	Age       int
	Score     float64
	Active    bool
	CreatedAt time.Time
	Home      address  // nested struct
	Billing   *address // pointer struct, allocated by Fill
	Ignored   string   `faker:"-"`
	secret    string   // unexported, must be skipped
}

func TestFillPopulatesEveryKind(t *testing.T) {
	f := seedfaker.New(1)
	var u user
	if err := f.Fill(&u); err != nil {
		t.Fatalf("Fill: %v", err)
	}
	if len(u.ID) != 36 {
		t.Fatalf("ID not a uuid: %q", u.ID)
	}
	if !strings.Contains(u.Email, "@") {
		t.Fatalf("Email not filled: %q", u.Email)
	}
	if u.Name == "" || u.Home.City == "" {
		t.Fatalf("name/nested city empty: %+v", u)
	}
	if u.Billing == nil || u.Billing.City == "" {
		t.Fatalf("pointer struct not allocated/filled: %+v", u.Billing)
	}
	if u.CreatedAt.IsZero() {
		t.Fatal("time.Time field not filled")
	}
	if u.Ignored != "" {
		t.Fatalf("faker:\"-\" field was filled: %q", u.Ignored)
	}
	if u.secret != "" {
		t.Fatalf("unexported field was filled: %q", u.secret)
	}
}

func TestFillIsDeterministic(t *testing.T) {
	var a, b user
	if err := seedfaker.New(9).Fill(&a); err != nil {
		t.Fatal(err)
	}
	if err := seedfaker.New(9).Fill(&b); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(a, b) {
		t.Fatalf("Fill not deterministic:\n%+v\n%+v", a, b)
	}
}

func TestFillRejectsBadArguments(t *testing.T) {
	f := seedfaker.New(1)
	if err := f.Fill(nil); err == nil {
		t.Fatal("Fill(nil) should error")
	}
	if err := f.Fill(user{}); err == nil {
		t.Fatal("Fill of a non-pointer should error")
	}
	var p *user
	if err := f.Fill(p); err == nil {
		t.Fatal("Fill of a nil typed pointer should error")
	}
	n := 5
	if err := f.Fill(&n); err == nil {
		t.Fatal("Fill of a pointer to non-struct should error")
	}
}

func TestFillRejectsUnknownTag(t *testing.T) {
	type bad struct {
		X string `faker:"no-such-generator"`
	}
	if err := seedfaker.New(1).Fill(&bad{}); err == nil {
		t.Fatal("unknown faker tag should error")
	}
}

func TestFillRejectsTypeMismatch(t *testing.T) {
	type bad struct {
		Age int `faker:"email"` // email yields string, not int
	}
	err := seedfaker.New(1).Fill(&bad{})
	if err == nil || !strings.Contains(err.Error(), "Age") {
		t.Fatalf("type mismatch should error naming the field, got %v", err)
	}
}

func TestRegisterGeneratorRejectsDuplicate(t *testing.T) {
	assertPanics(t, "duplicate email generator", func() {
		seedfaker.RegisterGenerator("email", func(*seedfaker.Faker) any { return "x" })
	})
}
