package structtag

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	test := []struct {
		name    string
		tag     string
		exp     []*Tag
		invalid bool
	}{
		{
			name: "empty tag",
			tag:  "",
		},
		{
			name:    "tag with one key (invalid)",
			tag:     "json",
			invalid: true,
		},
		{
			name: "tag with one key (valid)",
			tag:  `json:""`,
			exp: []*Tag{
				{Key: "json", Value: ""},
			},
		},
		{
			name: "tag with one key and dash value",
			tag:  `json:"-"`,
			exp: []*Tag{
				{Key: "json", Value: "-"},
			},
		},
		{
			name: "tag with key and value",
			tag:  `json:"foo"`,
			exp: []*Tag{
				{Key: "json", Value: "foo"},
			},
		},
		{
			name: "tag with key, value, and modifier",
			tag:  `json:"foo,omitempty"`,
			exp: []*Tag{
				{Key: "json", Value: "foo,omitempty"},
			},
		},
		{
			name: "tag with multiple keys",
			tag:  `json:"" hcl:""`,
			exp: []*Tag{
				{Key: "json", Value: ""},
				{Key: "hcl", Value: ""},
			},
		},
		{
			name: "tag with multiple keys and values",
			tag:  `json:"foo" hcl:"bar"`,
			exp: []*Tag{
				{Key: "json", Value: "foo"},
				{Key: "hcl", Value: "bar"},
			},
		},
		{
			name: "tag with multiple keys and modifiers",
			tag:  `json:"foo,omitempty" structs:"bar,omitnested"`,
			exp: []*Tag{
				{Key: "json", Value: "foo,omitempty"},
				{Key: "structs", Value: "bar,omitnested"},
			},
		},
		{
			name: "tag with quoted value",
			tag:  `json:"foo,bar:\"baz\""`,
			exp: []*Tag{
				{Key: "json", Value: "foo,bar:\"baz\""},
			},
		},
		{
			name: "tag with trailing space",
			tag:  `json:"foo" `,
			exp: []*Tag{
				{Key: "json", Value: "foo"},
			},
		},
		{
			name: "tag that is a simple comma",
			tag:  `json:"," `,
			exp: []*Tag{
				{Key: "json", Value: ","},
			},
		},
		{
			name: "tag that contains an escaped comma",
			tag:  `json:"abc\\,def" `,
			exp: []*Tag{
				{Key: "json", Value: "abc\\,def"},
			},
		},
	}

	for _, ts := range test {
		t.Run(ts.name, func(t *testing.T) {
			tags, err := Parse(ts.tag)
			invalid := err != nil

			if invalid != ts.invalid {
				t.Errorf("invalid case\n\twant: %+v\n\tgot : %+v\n\terr : %s", ts.invalid, invalid, err)
			}

			if invalid {
				return
			}

			for i, tag := range tags.Tags() {
				if tag.String() != ts.exp[i].String() {
					t.Errorf("parse\n\twant: %#v\n\tgot : %#v", ts.exp[i], tag)
				}
			}

			trimmedInput := strings.TrimSpace(ts.tag)
			if trimmedInput != tags.String() {
				t.Errorf("parse string\n\twant: %#v\n\tgot : %#v", trimmedInput, tags.String())
			}
		})
	}
}

func TestTags_String(t *testing.T) {
	tag := `json:"foo" structs:"bar,omitnested" hcl:"-"`

	tags, err := Parse(tag)
	if err != nil {
		t.Fatal(err)
	}

	if tags.String() != tag {
		t.Errorf("string\n\twant: %#v\n\tgot : %#v", tag, tags.String())
	}
}
