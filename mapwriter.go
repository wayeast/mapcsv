package mapcsv

import (
	"encoding/csv"
	"fmt"
	"io"
)

type mapWriter struct {
	w *csv.Writer
	h []string
}

type MapWriter interface {
	Flush()
	Write(record map[string]string) error
	WriteAll(records []map[string]string) error
	WriteHeader() error
}

func NewMapWriter(w io.Writer, fields []string, sep rune) *mapWriter {
	writer := csv.NewWriter(w)
	writer.Comma = sep

	return &mapWriter{
		w: writer,
		h: fields,
	}
}

func (w *mapWriter) Flush() {
	w.w.Flush()
}

func (w *mapWriter) Write(record map[string]string) error {
	var row []string
	for _, field := range w.h {
		v, ok := record[field]
		if ok {
			row = append(row, v)
		} else {
			return fmt.Errorf("Field %s not found in record %v", field, record)
		}
	}
	err := w.w.Write(row)
	return err
}

func (w *mapWriter) WriteAll(records []map[string]string) error {
	var rs [][]string
	for _, rec := range records {
		var row []string
		for _, field := range w.h {
			v, ok := rec[field]
			if ok {
				row = append(row, v)
			} else {
				return fmt.Errorf("Field %s not found in record %v", field, rec)
			}
		}
		rs = append(rs, row)
	}
	err := w.w.WriteAll(rs)
	return err
}

func (w *mapWriter) WriteHeader() error {
	err := w.w.Write(w.h)
	return err
}
