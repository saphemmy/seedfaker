package tests

import (
	"errors"
	"testing"

	"github.com/saphemmy/seedfaker"
)

type person struct {
	Name  string
	Email string
}

func personFactory() seedfaker.Factory[person] {
	return seedfaker.Factory[person]{
		Build: func(f *seedfaker.Faker) person {
			return person{Name: f.Name(), Email: f.Email()}
		},
	}
}

func TestFactoryMake(t *testing.T) {
	f := seedfaker.New(1)
	ps := personFactory().Make(f, 5)
	if len(ps) != 5 {
		t.Fatalf("Make(5) = %d", len(ps))
	}
	for _, p := range ps {
		if p.Name == "" || p.Email == "" {
			t.Fatalf("empty instance: %+v", p)
		}
	}
	if empty := personFactory().Make(seedfaker.New(1), 0); len(empty) != 0 {
		t.Fatalf("Make(0) = %d", len(empty))
	}
}

func TestFactoryIsDeterministic(t *testing.T) {
	a := personFactory().Make(seedfaker.New(2), 20)
	b := personFactory().Make(seedfaker.New(2), 20)
	for i := range a {
		if a[i] != b[i] {
			t.Fatalf("instance %d diverged: %+v != %+v", i, a[i], b[i])
		}
	}
}

func TestFactoryCreatePersistsAll(t *testing.T) {
	f := seedfaker.New(3)
	var stored []person
	got, err := personFactory().Create(f, 4, func(p person) error {
		stored = append(stored, p)
		return nil
	})
	if err != nil || len(got) != 4 || len(stored) != 4 {
		t.Fatalf("Create = %d stored=%d err=%v", len(got), len(stored), err)
	}
}

func TestFactoryCreateStopsAtFirstError(t *testing.T) {
	f := seedfaker.New(4)
	boom := errors.New("insert failed")
	calls := 0
	got, err := personFactory().Create(f, 10, func(person) error {
		calls++
		if calls == 3 {
			return boom
		}
		return nil
	})
	if !errors.Is(err, boom) {
		t.Fatalf("expected the persist error, got %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("Create should return the 2 instances stored before the failure, got %d", len(got))
	}
	if calls != 3 {
		t.Fatalf("persist called %d times, want 3 (stopped at the failure)", calls)
	}
}
