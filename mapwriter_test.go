package mapcsv

import (
	"bytes"
	"testing"
)

func TestWriter(t *testing.T) {
	var tests = []struct {
		Name   string
		Test   map[string]string
		Fields []string
		Sep    rune
		Result string
	}{
		{
			Name: "Basic",
			Test: map[string]string{
				"field1": "aaa",
				"field2": "bbb",
				"field3": "ccc",
			},
			Fields: []string{"field1", "field2", "field3"},
			Sep:    ',',
			Result: "field1,field2,field3\naaa,bbb,ccc\n",
		},
		{
			Name: "Default separator",
			Test: map[string]string{
				"field1": "aaa",
				"field2": "bbb",
				"field3": "ccc",
			},
			Fields: []string{"field1", "field2", "field3"},
			Sep:    0,
			Result: "field1,field2,field3\naaa,bbb,ccc\n",
		},
	}

	for _, test := range tests {
		var buf bytes.Buffer
		w := NewMapWriter(&buf, test.Fields, test.Sep)

		err := w.WriteHeader()
		if err != nil {
			t.Errorf("%s: Error writing header -- '%v'", test.Name, test.Fields)
		}

		err = w.WriteRow(test.Test)
		if err != nil {
			t.Errorf("%s: Error writing row -- '%v'", test.Name, test.Test)
		}

		w.Flush()
		s := buf.String()
		if s != test.Result {
			t.Errorf("%s: '%s' should be '%s'", test.Name, s, test.Result)
		}
	}
}
