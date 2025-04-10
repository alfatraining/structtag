package structtag_test

import (
	"fmt"
	"reflect"

	"github.com/alfatraining/structtag"
)

func Example() {
	type Example struct {
		Field string `json:"foo,omitempty" xml:"bar"`
	}

	// Get the struct tag from the field
	tag := reflect.TypeOf(Example{}).Field(0).Tag

	// Parse the tag using structtag
	tags, err := structtag.Parse(string(tag))
	if err != nil {
		panic(err)
	}

	// Iterate over all tags
	for _, t := range tags.Tags() {
		fmt.Printf("Key: %s, Value: %v\n", t.Key, t.Value)
	}
	// Output:
	// Key: json, Value: foo,omitempty
	// Key: xml, Value: bar
}
