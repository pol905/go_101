package reflections

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	t.Run("without maps", func(t *testing.T) {
		cases := []struct {
			Name          string
			Input         interface{}
			ExpectedCalls []string
		}{
			{
				Name: "Struct with one string field",
				Input: struct {
					Name string
				}{
					Name: "Siddhanth",
				},
				ExpectedCalls: []string{"Siddhanth"},
			},
			{
				Name: "Struct with two string fields",
				Input: struct {
					Name string
					City string
				}{Name: "Siddhanth", City: "Bengaluru"},
				ExpectedCalls: []string{"Siddhanth", "Bengaluru"},
			},
			{
				"Struct with non string field",
				struct {
					Name string
					Age  int
				}{Name: "Siddhanth", Age: 23},
				[]string{"Siddhanth"},
			},
			{
				"Nested fields",
				Person{
					Name: "Siddhanth",
					Profile: Profile{
						Age:  23,
						City: "Bengaluru",
					},
				},
				[]string{"Siddhanth", "Bengaluru"},
			},
			{
				"Pointers to things",
				&Person{
					Name: "Siddhanth",
					Profile: Profile{
						Age:  23,
						City: "Bengaluru",
					},
				},
				[]string{"Siddhanth", "Bengaluru"},
			},
			{
				"Slices",
				[]Profile{
					{Age: 23, City: "Bengaluru"},
					{Age: 22, City: "Chennai"},
				},
				[]string{"Bengaluru", "Chennai"},
			},
			{
				"Arrays",
				[2]Profile{
					{Age: 23, City: "Bengaluru"},
					{Age: 22, City: "Chennai"},
				},
				[]string{"Bengaluru", "Chennai"},
			},
		}

		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var got []string
				walk(test.Input, func(input string) {
					got = append(got, input)
				})

				if !reflect.DeepEqual(got, test.ExpectedCalls) {
					t.Errorf("got %v, want %v", got, test.ExpectedCalls)
				}
			})
		}
	})

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string

		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{Age: 23, City: "Bengaluru"}
			aChannel <- Profile{Age: 22, City: "Chennai"}
			close(aChannel)
		}()

		var got []string

		want := []string{"Bengaluru", "Chennai"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{Age: 23, City: "Bengaluru"},
				Profile{Age: 22, City: "Chennai"}
		}

		var got []string
		want := []string{"Bengaluru", "Chennai"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false

	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
