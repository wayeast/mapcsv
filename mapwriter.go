package mapcsv

import (
	"encoding/csv"
	"fmt"
	"io"
)

type mapWriter struct {
	csv.Writer
	fields []string
}

func NewMapWriter(w io.Writer, fields []string, sep rune) *mapWriter {
	writer := csv.NewWriter(w)
	if fields == nil {
		panic("fields argument may not be  nil")
	}

	if sep == 0 {
		sep = ','
	}
	writer.Comma = sep

	return &mapWriter{
		*writer,
		fields,
	}
}

func (w *mapWriter) WriteHeader() error {
	err := w.Write(w.fields)
	return err
}

func (w *mapWriter) WriteRow(row map[string]string) error {
	record := make([]string, len(w.fields))
	for i, field := range w.fields {
		v, ok := row[field]
		if ok {
			record[i] = v
		} else {
			return fmt.Errorf("Field %s not found in row %v", field, row)
		}
	}

	return w.Write(record)
}

func (w *mapWriter) WriteRows(rows []map[string]string) error {
	rs := make([][]string, len(rows))
	for i, row := range rows {
		record := make([]string, len(row))
		for j, field := range w.fields {
			v, ok := row[field]
			if ok {
				record[j] = v
			} else {
				return fmt.Errorf("Field %s not found in row %v", field, row)
			}
		}
		rs[i] = record
	}
	return w.WriteAll(rs)
}
