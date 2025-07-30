package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	type Profile struct {
		Age  uint
		City string
	}
	type Person struct {
		Name string
		Data Profile
	}

	cases := []struct {
		Name          string
		Input         any
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name     string
				Location string
			}{
				"Chris",
				"London",
			},
			[]string{"Chris", "London"},
		},
		{
			"struct with fields of different types",
			struct {
				Name string
				Age  uint
			}{
				"Chris",
				33,
			},
			[]string{"Chris"},
		},
		{
			"nested fields",
			Person{
				"Anna-bella",
				Profile{
					125,
					"BKC, Mumbai",
				},
			},
			[]string{"Anna-bella", "BKC, Mumbai"},
		},
		{
			"pointer to a structure",
			&Person{
				"Pluto",
				Profile{
					22,
					"Zirakpur",
				},
			},
			[]string{"Pluto", "Zirakpur"},
		},
		{
			"slices as input",
			[]Profile{
				{22, "Colombia"},
				{25, "Georgia"},
			},
			[]string{"Colombia", "Georgia"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("map as input", func(t *testing.T) {
		input := map[string]string{
			"Cow": "Moo",
			"Cat": "Meow",
		}

		var got []string
		walk(input, func(input string) {
			got = append(got, input)
		})

		for _, val := range input {
			assertContains(t, got, val)
		}
	})

	t.Run("channel as input", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Adam"}
			aChannel <- Profile{32, "Eve"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Adam", "Eve"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("func as input", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{22, "Pluto"}, Profile{22, "Pegion"}
		}

		var got []string
		want := []string{"Pluto", "Pegion"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func assertContains(t *testing.T, values []string, target string) {
	t.Helper()
	contains := false
	for _, val := range values {
		if val == target {
			contains = true
			break
		}
	}

	if !contains {
		t.Errorf("%v does not contain %s", values, target)
	}
}
