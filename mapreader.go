package mapcsv

import (
	"encoding/csv"
	"io"
	"log"
)

type mapReader struct {
	csv.Reader
	fields []string
}

func NewMapReader(fd io.Reader, fields []string, sep rune) *mapReader {
	reader := csv.NewReader(fd)
	if sep == 0 {
		sep = ','
	}
	reader.Comma = sep

	// If no fields passed to constructor, try to create them from first line of file
	var err error
	if fields == nil {
		fields, err = reader.Read()
		if err != nil {
			log.Fatalln("Error getting fieldnames from csv file.")
		}
	}

	return &mapReader{
		*reader,
		fields,
	}
}

func (r *mapReader) Fields() []string {
	return r.fields
}

func (r *mapReader) AsMap() (map[string]string, error) {
	record, err := r.Read()
	if err != nil {
		return nil, err
	}
	m := make(map[string]string, len(r.fields))
	for i, f := range r.fields {
		m[f] = record[i]
	}

	return m, nil
}

func (r *mapReader) AsMaps() (maps []map[string]string, err error) {
	for {
		record, err := r.Read()
		if err == io.EOF {
			return maps, nil
		}
		if err != nil {
			return nil, err
		}

		m := make(map[string]string, len(r.fields))
		for i, f := range r.fields {
			m[f] = record[i]
		}
		maps = append(maps, m)
	}
}
