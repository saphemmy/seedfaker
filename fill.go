package seedfaker

import (
	"fmt"
	"reflect"
	"time"
)

var generators = map[string]func(*Faker) any{}

// RegisterGenerator makes gen resolvable by Fill under a `faker:"name"` tag. It
// panics on a duplicate name, catching a collision at startup. Register custom
// generators from an init function.
func RegisterGenerator(name string, gen func(*Faker) any) {
	if _, dup := generators[name]; dup {
		panic("seedfaker: generator " + name + " registered twice")
	}
	generators[name] = gen
}

func init() {
	for name, gen := range map[string]func(*Faker) any{
		"name":      func(f *Faker) any { return f.Name() },
		"firstName": func(f *Faker) any { return f.FirstName() },
		"lastName":  func(f *Faker) any { return f.LastName() },
		"email":     func(f *Faker) any { return f.Email() },
		"username":  func(f *Faker) any { return f.Username() },
		"phone":     func(f *Faker) any { return f.Phone() },
		"uuid":      func(f *Faker) any { return f.UUID() },
		"company":   func(f *Faker) any { return f.Company() },
		"city":      func(f *Faker) any { return f.City() },
		"country":   func(f *Faker) any { return f.Country() },
		"word":      func(f *Faker) any { return f.Word() },
		"sentence":  func(f *Faker) any { return f.Sentence(6) },
	} {
		generators[name] = gen
	}
}

// Fill populates the exported fields of the struct pointed to by v. A `faker:"name"`
// tag selects a registered generator; faker:"-" and unexported fields are skipped;
// an untagged field is filled by its kind (string, bool, the integer and float
// kinds, time.Time, nested structs and non-nil-able pointers). It returns an error
// for a nil or non-struct-pointer argument, an unknown tag, or a generator whose
// type does not fit the field. Fields fill in declared order, so Fill is
// deterministic.
func (f *Faker) Fill(v any) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.IsNil() || rv.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("seedfaker: Fill requires a non-nil pointer to a struct, got %T", v)
	}
	return f.fillStruct(rv.Elem())
}

func (f *Faker) fillStruct(sv reflect.Value) error {
	st := sv.Type()
	for i := 0; i < sv.NumField(); i++ {
		field := st.Field(i)
		if !field.IsExported() {
			continue
		}
		tag := field.Tag.Get("faker")
		if tag == "-" {
			continue
		}
		if err := f.fillField(sv.Field(i), tag); err != nil {
			return fmt.Errorf("seedfaker: field %s: %w", field.Name, err)
		}
	}
	return nil
}

var timeType = reflect.TypeOf(time.Time{})

func (f *Faker) fillField(fv reflect.Value, tag string) error {
	if tag != "" {
		gen, ok := generators[tag]
		if !ok {
			return fmt.Errorf("unknown faker tag %q", tag)
		}
		rv := reflect.ValueOf(gen(f))
		if !rv.Type().AssignableTo(fv.Type()) {
			return fmt.Errorf("generator %q yields %s, not %s", tag, rv.Type(), fv.Type())
		}
		fv.Set(rv)
		return nil
	}
	switch fv.Kind() {
	case reflect.String:
		fv.SetString(f.Word())
	case reflect.Bool:
		fv.SetBool(f.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fv.SetInt(int64(f.IntBetween(0, 1<<31-1)))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fv.SetUint(uint64(f.IntBetween(0, 1<<31-1)))
	case reflect.Float32, reflect.Float64:
		fv.SetFloat(f.Float64())
	case reflect.Struct:
		if fv.Type() == timeType {
			fv.Set(reflect.ValueOf(f.Date()))
			return nil
		}
		return f.fillStruct(fv)
	case reflect.Pointer:
		if fv.IsNil() {
			fv.Set(reflect.New(fv.Type().Elem()))
		}
		return f.fillField(fv.Elem(), "")
	}
	return nil
}
