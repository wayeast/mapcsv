package mapcsv

import (
	"io"
	"strings"
	"testing"
)

func testFieldsEq(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if strings.Compare(a[i], b[i]) != 0 {
			return false
		}
	}

	return true
}

func TestSetFields(t *testing.T) {
	var fields = []string{"field1", "field2", "field3"}
	var csv = "aaa,bbb,ccc"

	reader := NewMapReader(strings.NewReader(csv), fields, 0)
	if !testFieldsEq(reader.Fields(), fields) {
		t.Errorf("Error setting fields: '%v' should be '%v'", reader.Fields(), fields)
	}
	rec, err := reader.AsMap()
	if err != nil {
		t.Error("Error reading csv", csv)
	}
	if rec["field1"] != "aaa" || rec["field2"] != "bbb" || rec["field3"] != "ccc" {
		t.Error("Error unmarshalling csv", csv)
	}
}

func TestSeparators(t *testing.T) {
	var standard_fields = []string{"field1", "field2", "field3"}
	var tests = []struct {
		Name string
		Csv  string
		Sep  rune
	}{
		{Name: "Basic", Csv: "field1,field2,field3\naaa,bbb,ccc\n", Sep: ','},
		{Name: "Default", Csv: "field1,field2,field3\naaa,bbb,ccc\n", Sep: 0},
		{Name: "Pipe", Csv: "field1|field2|field3\naaa|bbb|ccc\n", Sep: '|'},
		{Name: "Semicolon", Csv: "field1;field2;field3\naaa;bbb;ccc\n", Sep: ';'},
		{Name: "Tab", Csv: "field1\tfield2\tfield3\naaa\tbbb\tccc\n", Sep: '\t'},
	}

	for _, test := range tests {
		reader := NewMapReader(strings.NewReader(test.Csv), nil, test.Sep)
		test_fields := reader.Fields()
		if !testFieldsEq(test_fields, standard_fields) {
			t.Errorf("%s: Reader header incorrect -- '%v'", test.Name, test_fields)
		}
		rec, err := reader.AsMap()
		if err != nil {
			t.Errorf("%s: Error reading csv -- '%s'", test.Name, test.Csv)
		}
		if rec["field1"] != "aaa" || rec["field2"] != "bbb" || rec["field3"] != "ccc" {
			t.Errorf("%s: Error unmarshalling csv -- '%s'", test.Name, test.Csv)
		}
	}
}

func TestContentless(t *testing.T) {
	var csv = "field1,field2,field3"

	reader := NewMapReader(strings.NewReader(csv), nil, 0)
	rec, err := reader.AsMap()
	if rec != nil || err != io.EOF {
		t.Error("Bizarre reading of content-less csv")
	}
}
