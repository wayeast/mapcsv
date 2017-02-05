package mapcsv

import (
	"encoding/csv"
	"io"
	"log"
)

type mapReader struct {
	r *csv.Reader
	h []string
}

type MapReader interface {
	Read() (map[string]string, error)
	ReadAll() ([]map[string]string, error)
}

func NewMapReader(fd io.Reader, fields []string, sep rune) *mapReader {
	reader := csv.NewReader(fd)
	reader.Comma = sep

	var err error
	if fields == nil {
		fields, err = reader.Read()
		if err != nil {
			log.Fatalln("Error getting fieldnames from csv file.")
		}
	}

	return &mapReader{
		r: reader,
		h: fields,
	}
}

func (r *mapReader) Read() (map[string]string, error) {
	record, err := r.r.Read()
	if err != nil {
		return nil, err
	}
	m := make(map[string]string)
	for i, h := range r.h {
		m[h] = record[i]
	}

	return m, nil
}

func (r *mapReader) ReadAll() (maps []map[string]string, err error) {
	for {
		record, err := r.r.Read()
		if err == io.EOF {
			return maps, nil
		}
		if err != nil {
			return nil, err
		}

		m := make(map[string]string)
		for i, h := range r.h {
			m[h] = record[i]
		}
		maps = append(maps, m)
	}
}
