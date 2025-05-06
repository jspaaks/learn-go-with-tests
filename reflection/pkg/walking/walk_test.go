package walking

import (
	"reflect"
	"slices"
	"testing"
)

func TestWalk(t *testing.T) {
	type Address struct {
		HouseNumber int
		StreetName  string
	}
	type Person struct {
		Name    string
		Age     int
		Address Address
	}
	type Testcase struct {
		Name        string
		Input       any
		Expected    []string
		AssertEqual func(t *testing.T, want []string, got []string)
	}
	assertEqualOrdered := func(t *testing.T, want []string, got []string) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v but got %v instead", want, got)
		}
	}
	assertEqualUnordered := func(t *testing.T, want []string, got []string) {
		t.Helper()
		// test slice length
		if len(want) != len(got) {
			t.Errorf("slices have different length")
		}
		// verify there aren't any repeated values in slice 'want'
		wantedSet := map[string]struct{}{}
		for _, item := range want {
			wantedSet[item] = struct{}{}
		}
		if len(wantedSet) != len(want) {
			t.Errorf("testing equality of 'want' slice is unreliable if there are repeated items")
		}
		// test presence of each item from 'want' in 'got'
		for _, item := range want {
			if !slices.Contains(got, item) {
				t.Errorf("expected to find item %q but didn't", item)
			}
		}
	}
	var makeChannelAndSendSomeContent = func() chan Address {
		ch := make(chan Address)
		go func() {
			ch <- Address{HouseNumber: 34, StreetName: "5th"}
			ch <- Address{HouseNumber: 12, StreetName: "42nd"}
			close(ch)
		}()
		return ch
	}
	testcases := []Testcase{
		{
			"struct with one string field",
			struct {
				StreetName string
			}{"Pennsylvania Ave"},
			[]string{"Pennsylvania Ave"},
			assertEqualOrdered,
		},
		{
			"struct with two string fields",
			struct {
				StreetName string
				City       string
			}{"Pennsylvania Ave", "Washington"},
			[]string{"Pennsylvania Ave", "Washington"},
			assertEqualOrdered,
		},
		{
			"struct with one int and one string field",
			Address{HouseNumber: 123, StreetName: "Main St."},
			[]string{"Main St."},
			assertEqualOrdered,
		},
		{
			"struct with nested struct",
			Person{
				Name: "Chris",
				Age:  32,
				Address: Address{
					HouseNumber: 123,
					StreetName:  "Main St.",
				},
			},
			[]string{"Chris", "Main St."},
			assertEqualOrdered,
		},
		{
			"struct pointer",
			&Person{
				Name: "Chris",
				Age:  32,
				Address: Address{
					HouseNumber: 123,
					StreetName:  "Main St.",
				},
			},
			[]string{"Chris", "Main St."},
			assertEqualOrdered,
		},
		{
			"slices",
			[]Address{
				{
					HouseNumber: 1001,
					StreetName:  "5th",
				},
				{
					HouseNumber: 122,
					StreetName:  "42nd",
				},
			},
			[]string{"5th", "42nd"},
			assertEqualOrdered,
		},
		{
			"arrays",
			[2]Address{
				{
					HouseNumber: 1001,
					StreetName:  "5th",
				},
				{
					HouseNumber: 122,
					StreetName:  "42nd",
				},
			},
			[]string{"5th", "42nd"},
			assertEqualOrdered,
		},
		{
			"maps",
			map[string]string{
				"cow":   "moo",
				"sheep": "baa",
			},
			[]string{"baa", "moo"},
			assertEqualUnordered,
		},
		{
			"channel",
			makeChannelAndSendSomeContent(),
			[]string{"5th", "42nd"},
			assertEqualOrdered,
		},
		{
			"func",
			func() (Address, Address) {
				a := Address{StreetName: "3rd", HouseNumber: 3456}
				b := Address{StreetName: "5th", HouseNumber: 856}
				return a, b
			},
			[]string{"3rd", "5th"},
			assertEqualOrdered,
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			var got []string
			walk(testcase.Input, func(input string) {
				got = append(got, input)
			})
			testcase.AssertEqual(t, testcase.Expected, got)
		})
	}
}
