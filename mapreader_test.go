package mapcsv

import (
	"io"
	"strings"
	"testing"
)

func TestSetFields(t *testing.T) {
	var fields = []string{"field1", "field2", "field3"}
	var csv = "aaa,bbb,ccc"

	reader := NewMapReader(strings.NewReader(csv), fields, ',')
	rec, err := reader.Read()
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
		Csv string
		Sep rune
	}{
		{Csv: "field1,field2,field3\naaa,bbb,ccc\n", Sep: ','},
		{Csv: "field1|field2|field3\naaa|bbb|ccc\n", Sep: '|'},
		{Csv: "field1;field2;field3\naaa;bbb;ccc\n", Sep: ';'},
		{Csv: "field1\tfield2\tfield3\naaa\tbbb\tccc\n", Sep: '\t'},
	}
	for _, test := range tests {
		reader := NewMapReader(strings.NewReader(test.Csv), nil, test.Sep)
		test_fields := reader.Fields()
		for i, field := range standard_fields {
			if test_fields[i] != field {
				t.Error("Reader header incorrect:", test_fields)
			}
		}
		rec, err := reader.Read()
		if err != nil {
			t.Error("Error reading csv", test.Csv)
		}
		if rec["field1"] != "aaa" || rec["field2"] != "bbb" || rec["field3"] != "ccc" {
			t.Error("Error unmarshalling csv", test.Csv)
		}
	}
}

func TestContentless(t *testing.T) {
	var csv = "field1,field2,field3"

	reader := NewMapReader(strings.NewReader(csv), nil, ',')
	rec, err := reader.Read()
	if rec != nil || err != io.EOF {
		t.Error("Bizarre reading of content-less csv")
	}
}
