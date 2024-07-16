package reflection

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
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with single string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with >1 string field",
			struct {
				Name string
				City string
			}{"Chris", "Dublin"},
			[]string{"Chris", "Dublin"},
		},
		{
			"struct with int and string fields",
			struct {
				Name string
				City string
				Age  int
			}{"Chris", "Dublin", 134},
			[]string{"Chris", "Dublin"},
		},
		{
			"struct with lots of nested fields",
			Person{
				"Chris",
				Profile{134, "Dublin"},
			}, []string{"Chris", "Dublin"},
		},
		{
			"struct passed as a pointer",
			&Person{
				"Chris",
				Profile{134, "Dublin"},
			}, []string{"Chris", "Dublin"},
		},
		{
			"struct containing slices of values",
			[]Profile{{1, "Dublin"}, {2, "London"}, {3, "New York"}},
			[]string{"Dublin", "London", "New York"},
		},
		{
			"arrays",
			[2]Profile{
				{20, "Berlin"},
				{30, "Tokyo"},
			},
			[]string{"Berlin", "Tokyo"},
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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Katowice"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

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
		t.Errorf("expected %v to contains %q but it didn't", haystack, needle)
	}
}
