package mapcsv

import (
	"bytes"
	"testing"
)

func TestBasic(t *testing.T) {
	var tests = []struct {
		Test   map[string]string
		Fields []string
		Sep    rune
		Result string
	}{
		{
			Test: map[string]string{
				"field1": "aaa",
				"field2": "bbb",
				"field3": "ccc",
			},
			Fields: []string{"field1", "field2", "field3"},
			Sep:    ',',
			Result: "field1,field2,field3\naaa,bbb,ccc\n",
		},
	}

	for _, test := range tests {
		var buf bytes.Buffer
		w := NewMapWriter(&buf, test.Fields, test.Sep)

		err := w.WriteHeader()
		if err != nil {
			t.Error("Error writing header", test.Fields)
		}

		err = w.Write(test.Test)
		if err != nil {
			t.Error("Error writing row", test.Test)
		}

		w.Flush()
		s := buf.String()
		if s != test.Result {
			t.Error(s, "!=", test.Result)
		}
	}
}
