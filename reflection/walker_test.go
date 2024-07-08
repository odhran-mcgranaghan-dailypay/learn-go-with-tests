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
}
